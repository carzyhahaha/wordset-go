package common

const SUCCESS = 0
const FAIL = -1

type ApiCode struct {
}

type ApiResponse struct {
	Code int         `json:"name"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewApiResponse(code int, msg string, data interface{}) ApiResponse {
	return ApiResponse{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}
