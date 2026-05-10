package tools

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// DocumentResult 文档解析结果
type DocumentResult struct {
	Title   string
	Content string
	Type    string
}

// FetchDocument 获取并解析在线文档（对标 Python tools/document_tools.py）
func FetchDocument(rawURL string) (*DocumentResult, error) {
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Get(rawURL)
	if err != nil {
		return nil, fmt.Errorf("获取文档失败: %w", err)
	}
	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")
	parsed, _ := url.Parse(rawURL)

	if strings.Contains(contentType, "text/html") {
		return parseHTML(resp.Body)
	}

	fileName := parsed.Path
	if idx := strings.LastIndex(fileName, "/"); idx >= 0 {
		fileName = fileName[idx+1:]
	}
	if fileName == "" {
		fileName = "document"
	}

	if strings.Contains(contentType, "application/pdf") {
		return &DocumentResult{
			Title:   fileName,
			Content: "[PDF文件 - 建议使用 PDF 解析器]",
			Type:    "pdf",
		}, nil
	}

	body, _ := io.ReadAll(resp.Body)
	return &DocumentResult{
		Title:   fileName,
		Content: string(body),
		Type:    "text",
	}, nil
}

// parseHTML 解析 HTML 文档
func parseHTML(bodyReader io.Reader) (*DocumentResult, error) {
	doc, err := goquery.NewDocumentFromReader(bodyReader)
	if err != nil {
		return nil, err
	}

	title := doc.Find("title").Text()
	if title == "" {
		title = "未命名文档"
	}

	var contentParts []string
	doc.Find("p, h1, h2, h3, h4, h5, h6").Each(func(_ int, s *goquery.Selection) {
		if text := strings.TrimSpace(s.Text()); text != "" {
			contentParts = append(contentParts, text)
		}
	})

	return &DocumentResult{
		Title:   title,
		Content: strings.Join(contentParts, "\n"),
		Type:    "html",
	}, nil
}
