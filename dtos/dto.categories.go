package dtos

type DTOCategories struct {
	Name string `json:"name"`
}

type DTOCategoriesPagination struct {
	Limit       string `validate:"required,numeric" json:"limit"`
	Offset      string `validate:"required,numeric" json:"offset"`
	Sort        string `validate:"required,alpha" json:"sort"`
	Count       int    `validate:"numeric" json:"count"`
	CurrentPage string `validate:"required,numeric" json:"current_page"`
	TotalPage   int    `validate:"numeric" json:"total_page"`
}

type DTOCategoriesId struct {
	Id int `json:"id"`
}
