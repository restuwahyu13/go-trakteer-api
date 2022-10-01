package dtos

type DTOCustomersPagination struct {
	Limit       uint   `validate:"required,numeric" json:"limit"`
	Offset      uint   `validate:"required,numeric" json:"offset"`
	Sort        string `validate:"required,alpha" json:"sort"`
	Count       uint   `validate:"numeric" json:"count"`
	CurrentPage uint   `validate:"required,numeric" json:"current_page"`
	TotalPage   uint   `validate:"numeric" json:"total_page"`
}

type DTOCustomersRegister struct {
	Username    string `validate:"required,alphanum" json:"username"`
	Name        string `mod:"strip_num_unicode" validate:"required,alpha" json:"name"`
	Email       string `validate:"required,email" json:"email"`
	Password    string `validate:"required,alphanumunicode" json:"password"`
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

type DTOCustomerResetPasswordToken struct {
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
	Id uint `validate:"required,numeric,min=1" json:"id"`
}

type DTOCustomersUpdateProfileById struct {
	Email    string `validate:"required,email" json:"email"`
	Username string `validate:"required,alphanum" json:"username"`
	Name     string `mod:"strip_num_unicode" validate:"required,alpha" json:"name"`
	Active   *bool  `validate:"boolean" json:"active"`
}

type DTOCustomersSocialLink struct {
	GithubLink    string `json:"github_link"`
	DribbleLink   string `json:"dribble_link"`
	YoutubeLink   string `json:"youtube_link"`
	FacebookLink  string `json:"facebook_link"`
	InstagramLink string `json:"instagram_link"`
	LinkendinLink string `json:"linkendin_link"`
}

type DTOCustomersById struct {
	Id uint `validate:"required,numeric,min=1" json:"id"`
}

type DTOCustomersHealthToken struct {
	Token string `validate:"required" json:"access_token"`
}

type DTOCustomersRefreshToken struct {
	AccessToken string `validate:"required" json:"access_token"`
}
