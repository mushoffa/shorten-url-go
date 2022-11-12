package model

type BaseResponse struct {
	Data interface{} `json:"data,omitempty"`
	Error interface{} `json:"error,omitempty"`
}