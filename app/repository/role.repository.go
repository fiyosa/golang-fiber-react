package repository

import (
	"go-fiber-react/config"
)

var Role roleRepository

type roleRepository struct{}

func (*roleRepository) GetMany(user_id int, roles *[]string) error {
	return config.G.
		Table("users AS u").
		Distinct("r.name").
		Joins("LEFT JOIN user_has_roles AS uhr ON uhr.user_id = u.id").
		Joins("LEFT JOIN roles AS r ON r.id = uhr.role_id").
		Where("u.id = ?", user_id).
		Scan(roles).
		Error
}
