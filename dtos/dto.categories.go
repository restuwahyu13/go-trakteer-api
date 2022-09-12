package dtos

type DTOCategories struct {
	Name string `json:"name"`
}

type DTOCategoriesPagination struct {
	Limit       int    `json:"limit" default:"10"`
	Offset      int    `json:"offset" default:"0"`
	Sort        string `json:"sort" default:"asc"`
	Count       int    `json:"count"`
	Perpage     int    `json:"per_page" default:"10"`
	CurrentPage int    `json:"current_page" default:"1"`
	TotalPage   int    `json:"total_page"`
}
