package controller

import "wordset/common"

const API_SUCCESS int = 0
const API_FAIL int = -10411

func ApiSuccess(data interface{}) common.ApiResponse {
	return common.ApiResponse{
		Code: API_SUCCESS,
		Msg:  "",
		Data: data,
	}
}

func ApiFail(msg string) common.ApiResponse {
	return common.ApiResponse{
		Code: API_FAIL,
		Msg:  msg,
		Data: nil,
	}
}
