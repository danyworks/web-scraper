package models

import (
	"time"
)

type Response struct {
	Status    string      `json:"status,omitempty,required" `
	Message   string      `json:"message,omitempty,required" `
	Data      interface{} `json:"data,omitempty,required"`
	Timestamp string      `json:"timestamp,omitempty,required"`
}

func (r *Response) IsEmpty() bool {
	return r.Status == "" || r.Message == "" || r.Data == nil
}

func NewResponse(status, message string, data interface{}) Response {
	return Response{
		Status:    status,
		Message:   message,
		Data:      data,
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
	}
}
