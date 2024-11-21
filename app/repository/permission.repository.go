package repository

import (
	"go-fiber-react/config"
)

var Permission permissionRepository

type permissionRepository struct{}

func (*permissionRepository) GetManyByRoles(roles []string, permissions *[]string) error {
	return config.G.
		Table("permissions AS p").
		Distinct("p.name").
		Joins("LEFT JOIN role_has_permissions AS rhp ON rhp.permission_id = p.id").
		Joins("LEFT JOIN roles AS r ON r.id = rhp.role_id").
		Where("r.name IN ?", roles).
		Scan(permissions).
		Error
}

func (*permissionRepository) GetManyByUserId(user_id int, permissions *[]string) error {
	return config.G.
		Table("permissions AS p").
		Distinct("p.name").
		Joins("LEFT JOIN role_has_permissions AS rhp ON rhp.permission_id = p.id").
		Joins("LEFT JOIN roles AS r ON r.id = rhp.role_id").
		Joins("LEFT JOIN user_has_roles AS uhr ON uhr.role_id = r.id").
		Joins("LEFT JOIN users AS u ON u.id = uhr.user_id").
		Where("u.id = ?", user_id).
		Scan(permissions).
		Error
}
