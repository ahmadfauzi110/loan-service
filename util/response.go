package util

type Response struct {
	Message string      `json:"message"`
	Error   interface{} `json:"error,omitempty"`
	Data    interface{} `json:"data"`
}
