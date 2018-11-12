package common

type Response struct {
	Code int32	`json:"code"`
	Data interface{}	`json:"data"`
	Message string	`json:"message"`
	Success bool	`json:"success"`
}

func NewSuccessResponse(data interface{}) *Response {
	return &Response{
		0,
		data,
		"",
		true,
	}
}

func NewErrorResponse(code int32, message string) *Response {
	return &Response{
		code,
		nil,
		message,
		false,
	}
}
