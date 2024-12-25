package ctl

import "todo-list/pkg/errorStatus"

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

type DataList struct {
	Item  interface{} `json:"item"`
	Total int64       `json:"total"`
}

type TokenData struct {
	User        interface{} `json:"user"`
	AccessToken string      `json:"access_token"`
}

type TrackedErrorResponse struct {
	Response
	TraceId string `json:"trace_id"`
}

func RespSuccess(code ...int) *Response {
	status := errorStatus.Success
	if code != nil {
		status = code[0]
	}
	return &Response{
		Status: status,
		Msg:    "success",
	}
}

func RespSuccessWithData(data interface{}, code ...int) *Response {
	status := errorStatus.Success
	if code != nil {
		status = code[0]
	}
	return &Response{
		Status: status,
		Data:   data,
		Msg:    "success",
	}
}

func RespError(err error, data interface{}, code ...int) *TrackedErrorResponse {
	status := errorStatus.Error
	if code != nil {
		status = code[0]
	}
	return &TrackedErrorResponse{
		Response: Response{
			Status: status,
			Data:   data,
			Msg:    errorStatus.GetMsg(status),
			Error:  err.Error(),
		},
		TraceId: "",
	}
}
