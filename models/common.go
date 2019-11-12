package models

type CommonResp struct {
	Message   string      `json:"message"`
	Timestamp int64       `json:"timestamp"`
	Data      interface{} `json:"data,omitempty"`
}
