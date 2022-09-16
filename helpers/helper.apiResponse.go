package helpers

import (
	"encoding/json"
	"log"
)

type APIResponse struct {
	StatCode   uint   `json:"stat_code"`
	StatMsg    string `json:"stat_msg"`
	Data       any    `json:"data,omitempty"`
	Pagination any    `json:"pagination,omitempty"`
	QueryError error  `json:"sql_error,omitempty"`
}

func ApiResponse(data APIResponse) []byte {
	res, err := json.Marshal(data)
	if err != nil {
		log.Printf("Formatting json response error: %v", err)
	}
	return res
}
