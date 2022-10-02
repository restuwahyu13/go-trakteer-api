package dtos

type DTODonationType struct {
	Name string `mod:"strip_num_unicode" validate:"required,alpha" json:"name" form:"name"`
	Icon string `validate:"required,url" json:"icon" form:"icon"`
}

type DTODonationTypePagination struct {
	Limit       uint   `validate:"required,numeric" json:"limit"`
	Offset      *uint  `validate:"required,numeric" json:"offset"`
	Sort        string `validate:"required,alpha" json:"sort"`
	Count       uint   `validate:"numeric" json:"count"`
	CurrentPage uint   `validate:"required,numeric" json:"current_page"`
	TotalPage   uint   `validate:"numeric" json:"total_page"`
}

type DTODonationTypeId struct {
	Id uint `validate:"required,numeric" json:"id"`
}
