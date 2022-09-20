package dtos

type DTOUsersPagination struct {
	Limit       uint   `validate:"required,numeric" json:"limit"`
	Offset      *uint  `validate:"required,numeric" json:"offset"`
	Sort        string `validate:"required,alpha" json:"sort"`
	Count       int    `validate:"numeric" json:"count"`
	CurrentPage uint   `validate:"required,numeric" json:"current_page"`
	TotalPage   int    `validate:"numeric" json:"total_page"`
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

type DTOUsersResetPasswordToken struct {
	Token string `validate:"required,base64" json:"token"`
}

type DTOUsersResetPassword struct {
	Password  string `validate:"required,alphanumunicode,min=8" json:"password"`
	Cpassword string `validate:"required,alphanumunicode,min=8" json:"cpassword"`
}

type DTOUsersChangePassword struct {
	Password  string `validate:"required,alphanumunicode,min=8" json:"password"`
	Cpassword string `validate:"required,alphanumunicode,min=8" json:"cpassword"`
}

type DTOUsersGetProfileById struct {
	Id uint `validate:"required,numeric,min=1" json:"id"`
}

type DTOUsersUpdateProfileById struct {
	Email    string `mod:"trim" validate:"required,email" json:"email"`
	Username string `mod:"trim" validate:"required,alphanum" json:"username"`
	Name     string `mod:"strip_num_unicode" validate:"required,alpha" json:"name"`
	Active   *bool  `validate:"boolean" json:"active"`
}

type DTOUsersCreate struct {
	Username    string `mod:"trim" validate:"required,alphanum" json:"username"`
	Name        string `mod:"strip_num_unicode" validate:"required,alpha" json:"name"`
	Email       string `mod:"trim" validate:"required,email" json:"email"`
	Password    string `validate:"required,alphanumunicode" json:"password"`
	RoleId      uint   `validate:"required,numeric" json:"role_id"`
	CategorieId uint   `validate:"required,numeric" json:"categorie_id"`
}

type DTOUsersById struct {
	Id uint `validate:"required,numeric,min=1" json:"id"`
}

type DTOUsersUpdate struct {
	Username    string `mod:"trim" validate:"required,alphanum" json:"username"`
	Name        string `mod:"strip_num_unicode" validate:"required,alpha" json:"name"`
	Email       string `mod:"trim" validate:"required,email" json:"email"`
	Active      *bool  `validate:"required,boolean" json:"active"`
	RoleId      uint   `validate:"required,numeric" json:"role_id"`
	CategorieId uint   `validate:"required,numeric" json:"categorie_id"`
}
