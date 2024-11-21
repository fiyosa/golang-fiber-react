package resource_auth

import (
	"go-fiber-react/app/http/resource/resource_user"
)

type Register struct {
	Data    resource_user.Show `json:"data"`
	Message string             `json:"message" example:""`
}
