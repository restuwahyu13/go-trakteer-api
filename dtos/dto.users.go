package dtos

type DTOUsersPagination struct {
	Limit       int    `json:"limit" default:"10"`
	Offset      int    `json:"offset" default:"0"`
	Sort        string `json:"sort" default:"asc"`
	Count       int    `json:"count"`
	CurrentPage int    `json:"current_page" default:"1"`
	TotalPage   int    `json:"total_page"`
}

type DTOUsersLogin struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,alphanumunicode" json:"password"`
}

type DTOUsersActivation struct {
	Token string `validate:"required,base64" json:"token"`
}

type DTOUsersForgotPassword struct {
	Email string `validate:"required,email" json:"email"`
}

type DTOUsersResetPassword struct {
	Token     string `validate:"required,base64" json:"token"`
	Password  string `validate:"required,alphanumunicode,min=8" json:"password"`
	Cpassword string `validate:"required,alphanumunicode,min=8" json:"cpassword"`
}

type DTOUsersChangePassword struct {
	Password  string `validate:"required,alphanumunicode,min=8" json:"password"`
	Cpassword string `validate:"required,alphanumunicode,min=8" json:"cpassword"`
}

type DTOUsersGetProfileById struct {
	Id int `validate:"required,numeric,min=1" json:"id"`
}

type DTOUsersUpdateProfileById struct {
	Email    string `validate:"required,email" json:"email"`
	Username string `validate:"required,alphanum" json:"username"`
	Name     string `validate:"required,alpha" json:"name"`
	Active   bool   `validate:"boolean" json:"active"`
}

type DTOUsersCreate struct {
	Username    string `validate:"required,alphanum" json:"username"`
	Name        string `validate:"required,alpha" json:"name"`
	Email       string `validate:"required,email" json:"email"`
	Password    string `validate:"required,alphanumunicode" json:"password"`
	Active      bool   `validate:"required,boolean" json:"active"`
	Verified    bool   `validate:"required,boolean" json:"verified"`
	RoleId      *uint  `validate:"required,number" json:"role_id"`
	CategorieId *uint  `validate:"required,number" json:"categorie_id"`
}

type DTOUsersById struct {
	Id int `validate:"required,numeric,min=1" json:"id"`
}

type DTOUsersUpdate struct {
	Username    string `validate:"required,alphanum" json:"username"`
	Name        string `validate:"required,alpha" json:"name"`
	Email       string `validate:"required,email" json:"email"`
	Password    string `validate:"required,alphanumunicode" json:"password"`
	Active      bool   `validate:"required,boolean" json:"active"`
	Verified    bool   `validate:"required,boolean" json:"verified"`
	RoleId      *uint  `validate:"required,number" json:"role_id"`
	CategorieId *uint  `validate:"required,number" json:"categorie_id"`
}
