package models

type JsonResponsePayload struct {
	Status  string      `json:"status"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}
