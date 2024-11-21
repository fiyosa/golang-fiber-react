package request_auth

type Login struct {
	Username string `json:"username" form:"username" validate:"required,min=3" example:""`
	Password string `json:"password" form:"password" validate:"required,min=3" example:""`
}
