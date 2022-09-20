package dtos

type DTOCategories struct {
	Name string `mod:"strip_num_unicode" validate:"required,alpha" json:"name"`
}

type DTOCategoriesPagination struct {
	Limit       uint   `validate:"required,numeric" json:"limit"`
	Offset      *uint  `validate:"required,numeric" json:"offset"`
	Sort        string `validate:"required,alpha" json:"sort"`
	Count       uint   `validate:"numeric" json:"count"`
	CurrentPage uint   `validate:"required,numeric" json:"current_page"`
	TotalPage   uint   `validate:"numeric" json:"total_page"`
}

type DTOCategoriesId struct {
	Id uint `validate:"required,numeric" json:"id"`
}
