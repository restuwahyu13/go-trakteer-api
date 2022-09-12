package helpers

import (
	"encoding/json"
	"math"
)

type paginationOptions struct {
	Limit       int    `json:"limit" default:"10"`
	Offset      int    `json:"offset" default:"0"`
	Sort        string `json:"sort" default:"asc"`
	Count       int    `json:"count"`
	Perpage     int    `json:"per_page" default:"10"`
	CurrentPage int    `json:"current_page" default:"1"`
	TotalPage   int    `json:"total_page"`
}

func Pagination(pagination interface{}, count int) paginationOptions {
	payload := make(map[string]any)
	pagin := paginationOptions{}

	jsn, _ := json.Marshal(pagination)
	json.Unmarshal(jsn, &payload)

	pagin.Limit = int(payload["limit"].(float64))
	pagin.Offset = int(payload["offset"].(float64))
	pagin.Sort = payload["sort"].(string)
	pagin.Perpage = int(payload["per_page"].(float64))
	pagin.Count = count
	pagin.CurrentPage = int(payload["current_page"].(float64))

	if pagin.CurrentPage > 1 {
		pagin.Offset = pagin.Offset + pagin.CurrentPage - 1
	} else {
		pagin.Offset = 0
	}

	pagin.TotalPage = int(math.Abs(math.Ceil(float64(pagin.Count) / float64(pagin.Perpage))))

	return pagin
}
