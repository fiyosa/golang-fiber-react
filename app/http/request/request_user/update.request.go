package request_user

type Update struct {
	Name string `json:"name" form:"name" validate:"required,min=1" example:""`
}
