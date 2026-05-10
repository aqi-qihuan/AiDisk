package services

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/aqi/AqiCloud-Ai-Agent-Go/internal/models"
	"github.com/ledongthuc/pdf"
)

// DocumentService 文档服务（对标 Python services/doc_service.py）
type DocumentService struct{}

// FetchDocument 获取并解析文档内容
func (ds *DocumentService) FetchDocument(rawURL string) (title string, content string, docType string, err error) {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return "", "", "", err
	}

	resp, err := http.Get(rawURL)
	if err != nil {
		return "", "", "", err
	}
	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")

	if strings.Contains(contentType, "text/html") {
		return ds.parseHTML(resp.Body)
	}
	if strings.Contains(contentType, "application/pdf") {
		body, _ := io.ReadAll(resp.Body)
		fileName := parsed.Path
		if idx := strings.LastIndex(fileName, "/"); idx >= 0 {
			fileName = fileName[idx+1:]
		}
		text, err := ds.extractPDFText(body)
		if err != nil {
			return fileName, "", "pdf", nil
		}
		return fileName, text, "pdf", nil
	}

	body, _ := io.ReadAll(resp.Body)
	fileName := parsed.Path
	if idx := strings.LastIndex(fileName, "/"); idx >= 0 {
		fileName = fileName[idx+1:]
	}
	return fileName, string(body), "text", nil
}

// extractPDFText 从 PDF 二进制数据中提取文本
func (ds *DocumentService) extractPDFText(data []byte) (string, error) {
	reader, err := pdf.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return "", fmt.Errorf("PDF 解析失败: %w", err)
	}

	var buf strings.Builder
	n := reader.NumPage()
	for i := 1; i <= n; i++ {
		page := reader.Page(i)
		if page.V.IsNull() {
			continue
		}
		text, err := page.GetPlainText(nil)
		if err == nil && text != "" {
			buf.WriteString(text)
			buf.WriteString("\n\n")
		}
	}

	result := strings.TrimSpace(buf.String())
	if result == "" {
		return "", fmt.Errorf("PDF 中未提取到文本内容（可能是扫描版图片PDF）")
	}
	return result, nil
}

// parseHTML 解析 HTML 文档
func (ds *DocumentService) parseHTML(body io.Reader) (title, content, docType string, err error) {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return "", "", "", err
	}

	title = doc.Find("title").Text()
	if title == "" {
		title = "未命名文档"
	}

	var parts []string
	doc.Find("p, h1, h2, h3, h4, h5, h6").Each(func(_ int, s *goquery.Selection) {
		if text := strings.TrimSpace(s.Text()); text != "" {
			parts = append(parts, text)
		}
	})

	return title, strings.Join(parts, "\n"), "html", nil
}

// ChunkContent 按段落分割长文档
func (ds *DocumentService) ChunkContent(content string, maxChunkSize int) []string {
	if maxChunkSize <= 0 {
		maxChunkSize = 1000000
	}

	if len(content) <= maxChunkSize {
		return []string{content}
	}

	var chunks []string
	var currentChunk string

	paragraphs := strings.Split(content, "\n\n")
	for _, p := range paragraphs {
		if len(currentChunk)+len(p) > maxChunkSize && currentChunk != "" {
			chunks = append(chunks, currentChunk)
			currentChunk = p
		} else {
			if currentChunk != "" {
				currentChunk += "\n\n"
			}
			currentChunk += p
		}
	}
	if currentChunk != "" {
		chunks = append(chunks, currentChunk)
	}
	return chunks
}

// BuildDocInput 构建 LLM 输入文本
func BuildDocInput(req models.DocumentRequest, title, content string) string {
	length := req.Length
	if length == "" {
		length = "无限制"
	}
	instructions := req.AdditionalInstructions
	if instructions == "" {
		instructions = "无"
	}

	return fmt.Sprintf(
		"文档标题: %s\n文档内容: %s\n总结类型: %s\n输出语言: %s\n最大长度: %s\n额外要求: %s",
		title, content, req.SummaryType, req.Language, length, instructions,
	)
}
