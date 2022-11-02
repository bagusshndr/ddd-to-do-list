package shared

import "encoding/json"

type Response struct {
	Status       bool        `json:"status"`
	Code         string      `json:"code"`
	Message      string      `json:"message"`
	ErrorMessage interface{} `json:"error_message"`
	Data         interface{} `json:"data"`
}

func (r *Response) JSON() []byte {
	result, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return result
}

func NewResponse(status bool, code, message string, errorMessage, data interface{}) *Response {
	return &Response{
		Status:       status,
		Code:         code,
		Message:      message,
		ErrorMessage: errorMessage,
		Data:         data,
	}
}
