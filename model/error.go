package model

type Error struct {
	Error_code string `json:"code"`
	Error_type string `json:"errorType"`
	Message    string `json:"message"`
}
