package model

type JsonData struct {
	Code    int     `json:"code"`
	Data    any     `json:"data"`
	Msg     *string `json:"msg"`
	Success bool    `json:"success"`
}

func Success(data any) *JsonData {
	return &JsonData{Code: 0, Data: data, Msg: nil, Success: true}
}

func Error(msg string, code int) *JsonData {
	return &JsonData{Code: code, Data: nil, Msg: &msg, Success: false}
}

type AiResponseDto struct {
	Content      string `json:"content"`
	Model        string `json:"model"`
	Provider     string `json:"provider"`
	Success      bool   `json:"success"`
	ErrorMessage string `json:"errorMessage"`
	Metadata     map[string]interface{} `json:"metadata"`
}

type OllamaResponseDto struct {
	Model   string `json:"model"`
	Message struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"message"`
	Done bool `json:"done"`
}

// Error codes matching Java BizCodeEnum
const (
	CodeAccountExists    = 250001
	CodeAccountNotExists = 250002
	CodeWrongPassword    = 250003
	CodeNotLogin         = 20004
	CodeFileNotFound     = 220404
	CodeFileUploadError  = 220408
	CodeFileSizeExceeded = 220409
	CodeFileNameDuplicate = 220405
	CodeBatchDelIllegal = 220406
	CodeRecycleIllegal  = 280406
	CodeFileTypeError   = 220407
	CodeChunkTaskNotExists = 230408
	CodeChunkCountInsufficient = 230409
	CodeStorageNotEnough = 240403
	CodeTargetParentIllegal = 250403
	CodeCancelShareIllegal = 260403
	CodeShareCodeIllegal = 260404
	CodeShareNotExists  = 260405
	CodeShareCanceled   = 260406
	CodeShareExpired    = 260407
	CodeShareFileIllegal = 260408
	CodeBatchOpError    = 270101
	CodeParamError      = 220500
)
