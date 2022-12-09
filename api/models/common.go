package models

type JSONResult struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type JSONErrorResult struct {
	Error string `json:"error"`
}
type ResponseModel struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}
