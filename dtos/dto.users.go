package dtos

type DTOUsers struct {
	Username    string `validate:"required,alphanum" json:"username"`
	Name        string `validate:"required,alpha" json:"name"`
	Email       string `validate:"required,email" json:"email"`
	Password    string `validate:"required,alphanumunicode" json:"password"`
	RoleID      int    `validate:"required,number" json:"role_id"`
	CategorieID int    `validate:"required,number" json:"categorie_id"`
}
