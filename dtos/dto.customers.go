package dtos

type DTOCustomersPagination struct {
	Limit       int    `json:"limit" default:"10"`
	Offset      int    `json:"offset" default:"0"`
	Sort        string `json:"sort" default:"asc"`
	Count       int    `json:"count"`
	CurrentPage int    `json:"current_page" default:"1"`
	TotalPage   int    `json:"total_page"`
}

type DTOCustomersRegister struct {
	Username    string `validate:"required,alphanum" json:"username"`
	Name        string `validate:"required,alpha" json:"name"`
	Email       string `validate:"required,email" json:"email"`
	Password    string `validate:"required,alphanumunicode" json:"password"`
	Active      bool   `validate:"required,boolean" json:"active"`
	Verified    bool   `validate:"required,boolean" json:"verified"`
	RoleId      *uint  `validate:"required,number" json:"role_id"`
	CategorieId *uint  `validate:"required,number" json:"categorie_id"`
}

type DTOCustomersLogin struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,alphanumunicode" json:"password"`
}

type DTOCustomersActivation struct {
	Token string `validate:"required,base64" json:"token"`
}

type DTOCustomersForgotPassword struct {
	Email string `validate:"required,email" json:"email"`
}

type DTOCustomersResendActivation struct {
	Email string `validate:"required,email" json:"email"`
}

type DTOCustomersResetPassword struct {
	Password  string `validate:"required,alphanumunicode,min=8" json:"password"`
	Cpassword string `validate:"required,alphanumunicode,min=8" json:"cpassword"`
}

type DTOCustomersChangePassword struct {
	Password  string `validate:"required,alphanumunicode,min=8" json:"password"`
	Cpassword string `validate:"required,alphanumunicode,min=8" json:"cpassword"`
}

type DTOCustomersGetProfileById struct {
	Id int `validate:"required,numeric,min=1" json:"id"`
}

type DTOCustomersUpdateProfileById struct {
	Email    string `validate:"required,email" json:"email"`
	Username string `validate:"required,alphanum" json:"username"`
	Name     string `validate:"required,alpha" json:"name"`
	Active   bool   `validate:"boolean" json:"active"`
}
