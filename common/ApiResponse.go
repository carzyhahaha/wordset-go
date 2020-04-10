package common

const SUCCESS = 0
const FAIL = -1

type ApiCode struct {

}

type ApiResponse struct {
	Code int
	Msg string
	Data interface{}
}

func (d *ApiResponse) Success(data interface{}) ApiResponse{
	return ApiResponse{
		Code: SUCCESS,
		Msg:  "success",
		Data: data,
	}
}

func (d *ApiResponse) fail(msg string) ApiResponse {
	return ApiResponse{
		Code: FAIL,
		Msg:  msg,
		Data: nil,
	}
}






