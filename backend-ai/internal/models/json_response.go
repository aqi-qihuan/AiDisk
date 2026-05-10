package models

// JsonData 统一响应格式（对标 Python models/json_response.py）
type JsonData struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
	Type string `json:"type"`
}

func Success(data any) JsonData {
	return JsonData{Code: 0, Data: data, Msg: "", Type: "text"}
}

func Error(msg string, code int) JsonData {
	return JsonData{Code: code, Msg: msg, Data: "", Type: "text"}
}

func StreamData(data any, msg ...string) JsonData {
	m := ""
	if len(msg) > 0 {
		m = msg[0]
	}
	return JsonData{Code: 0, Data: data, Msg: m, Type: "stream"}
}
