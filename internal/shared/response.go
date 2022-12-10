package shared

import (
	"github.com/labstack/echo/v4"
)

type Response struct {
	Status       string      `json:"status"`
	Code         int         `json:"-"`
	Message      string      `json:"message"`
	ErrorMessage interface{} `json:"-"`
	Data         interface{} `json:"data"`
}

func (r *Response) JSON(c echo.Context) error {
	return c.JSON(r.Code, r)
}

func NewResponse(status string, code int, message string, errorMessage, data interface{}) *Response {
	return &Response{
		Status:       status,
		Code:         code,
		Message:      message,
		ErrorMessage: errorMessage,
		Data:         data,
	}
}
