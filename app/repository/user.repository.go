package repository

import (
	"go-fiber-react/app/model"
	"go-fiber-react/config"

	"github.com/gofiber/fiber/v2"
)

var User userRepository

type userRepository struct{}

func (*userRepository) First(c *fiber.Ctx, user_id int, u *model.User) error {
	return config.G.Model(&model.User{}).Where("id = ?", user_id).Scan(&u).Error
}

func (*userRepository) Create(u *model.User) error {
	return config.G.Create(&u).Error
}

func (*userRepository) Update(u *model.User) error {
	return config.G.Save(&u).Error
}
