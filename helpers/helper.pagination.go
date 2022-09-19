package helpers

import (
	"encoding/json"
	"math"
	"strconv"
)

type paginationOptions struct {
	Limit       int    `json:"limit" default:"10"`
	Offset      int    `json:"offset" default:"0"`
	Sort        string `json:"sort" default:"asc"`
	Count       int    `json:"count"`
	CurrentPage int    `json:"current_page" default:"1"`
	TotalPage   int    `json:"total_page"`
}

func Pagination(pagination interface{}, count int) paginationOptions {
	payload := make(map[string]any)
	pagin := paginationOptions{}

	jsn, _ := json.Marshal(pagination)
	json.Unmarshal(jsn, &payload)

	limit, _ := strconv.Atoi(payload["limit"].(string))
	offset, _ := strconv.Atoi(payload["offset"].(string))
	current_page, _ := strconv.Atoi(payload["current_page"].(string))

	pagin.Limit = limit
	pagin.Offset = offset
	pagin.Sort = payload["sort"].(string)
	pagin.Count = count
	pagin.CurrentPage = current_page

	if pagin.CurrentPage > 1 {
		pagin.Offset = (pagin.Limit * pagin.CurrentPage) - pagin.Limit
	} else {
		pagin.Offset = 0
	}

	pagin.TotalPage = int(math.Abs(math.Ceil(float64(pagin.Count) / float64(pagin.Limit))))

	return pagin
}
