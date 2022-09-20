package helpers

import (
	"encoding/json"
	"math"
	"strconv"
)

type paginationOptions struct {
	Limit       uint   `json:"limit" default:"10"`
	Offset      uint   `json:"offset" default:"0"`
	Sort        string `json:"sort" default:"asc"`
	Count       uint   `json:"count"`
	CurrentPage uint   `json:"current_page" default:"1"`
	TotalPage   uint   `json:"total_page"`
}

func Pagination(pagination interface{}, count int) paginationOptions {
	payload := make(map[string]any)
	pagin := paginationOptions{}

	jsn, _ := json.Marshal(pagination)
	json.Unmarshal(jsn, &payload)

	limit, _ := strconv.Atoi(payload["limit"].(string))
	offset, _ := strconv.Atoi(payload["offset"].(string))
	current_page, _ := strconv.Atoi(payload["current_page"].(string))

	pagin.Limit = uint(limit)
	pagin.Offset = uint(offset)
	pagin.Sort = payload["sort"].(string)
	pagin.Count = uint(count)
	pagin.CurrentPage = uint(current_page)

	if pagin.CurrentPage > 1 {
		pagin.Offset = (pagin.Limit * pagin.CurrentPage) - pagin.Limit
	} else {
		pagin.Offset = 0
	}

	pagin.TotalPage = uint(math.Abs(math.Ceil(float64(pagin.Count) / float64(pagin.Limit))))

	return pagin
}
