package request

type LoginAuthReq struct {
	Username string `json:"username" form:"username" validate:"required,min=3" example:""`
	Password string `json:"password" form:"password" validate:"required,min=3" example:""`
}

type LoginAuthRes struct {
	Token string `json:"access_token" example:""`
}

type RegisterReq struct {
	Username string `json:"username" form:"username" from:"username" validate:"required,min=3" example:""`
	Password string `json:"password" form:"password" from:"password" validate:"required,min=3" example:""`
	Name     string `json:"name" form:"name" from:"name" validate:"required,min=3" example:""`
}

type RegisterRes struct {
	Data    UserShowRes `json:"data"`
	Message string      `json:"message" example:""`
}
