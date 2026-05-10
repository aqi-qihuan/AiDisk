package tools

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var (
	httpClient     *http.Client
	httpClientOnce sync.Once
)

func getHTTPClient() *http.Client {
	httpClientOnce.Do(func() {
		httpClient = &http.Client{
			Timeout: 30 * time.Second,
		}
	})
	return httpClient
}

// WebSearch 搜索网络信息（对标 Python tools/chat_tools.py web_search）
func WebSearch(query string) (string, error) {
	searchURL := fmt.Sprintf("https://html.duckduckgo.com/html/?q=%s", url.QueryEscape(query))

	resp, err := getHTTPClient().Get(searchURL)
	if err != nil {
		return "", fmt.Errorf("搜索失败: %w", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", fmt.Errorf("解析搜索结果: %w", err)
	}

	var results []string
	doc.Find(".result").Each(func(i int, s *goquery.Selection) {
		if i >= 5 {
			return
		}
		title := ""
		s.Find(".result__a").Each(func(_ int, a *goquery.Selection) {
			title = a.Text()
		})
		snippet := ""
		s.Find(".result__snippet").Each(func(_ int, sn *goquery.Selection) {
			snippet = sn.Text()
		})
		if title != "" {
			results = append(results, fmt.Sprintf("来源: %s\n内容: %s", title, snippet))
		}
	})

	if len(results) == 0 {
		return "未找到相关搜索结果", nil
	}

	return strings.Join(results, "\n\n"), nil
}
