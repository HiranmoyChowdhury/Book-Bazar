package model

type Error struct {
	Error_code string `json:"code"`
	Error_type string `json:"error_type"`
	Message    string `json:"message"`
}
