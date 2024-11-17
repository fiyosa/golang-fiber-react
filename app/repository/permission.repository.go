package repository

import (
	"go-fiber-react/config"

	"gorm.io/gorm"
)

type Permission struct{}

func (*Permission) GetPermissions(roles []string, permissions *[]string) *gorm.DB {
	return config.G.
		Table("permissions AS p").
		Distinct("p.name").
		Joins("LEFT JOIN role_has_permissions AS rhp ON rhp.permission_id = p.id").
		Joins("LEFT JOIN roles AS r ON r.id = rhp.role_id").
		Where("r.name IN ?", roles).
		Scan(&permissions)
}
