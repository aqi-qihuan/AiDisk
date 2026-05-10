package models

// DocumentRequest 文档处理请求（对标 Python models/doc_schemas.py）
type DocumentRequest struct {
	URL                    string `json:"url"`
	SummaryType            string `json:"summary_type"`
	Language               string `json:"language"`
	Length                 string `json:"length"`
	AdditionalInstructions string `json:"additional_instructions"`
}

// DocQueryRequest 文档查询请求
type DocQueryRequest struct {
	Query        string `json:"query"`
	DocumentURL  string `json:"document_url"`
	AccountID    string `json:"account_id"`
}
