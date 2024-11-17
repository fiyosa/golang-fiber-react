package repository

import (
	"go-fiber-react/config"

	"gorm.io/gorm"
)

type Role struct{}

func (*Role) GetRoles(user_id int, roles *[]string) *gorm.DB {
	return config.G.
		Table("users AS u").
		Select("r.name").
		Joins("LEFT JOIN user_has_roles AS uhr ON uhr.user_id = u.id").
		Joins("LEFT JOIN roles AS r ON r.id = uhr.role_id").
		Where("u.id = ?", user_id).
		Scan(&roles)
}
