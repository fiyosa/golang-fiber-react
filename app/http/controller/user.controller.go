package controller

import (
	"encoding/json"
	"go-fiber-react/app/helper"
	"go-fiber-react/app/http/request/request_user"
	"go-fiber-react/app/http/resource/resource_user"
	"go-fiber-react/app/model"
	"go-fiber-react/app/repository"
	"go-fiber-react/config"
	"go-fiber-react/lang"
	"time"

	"github.com/gofiber/fiber/v2"
)

var User userController

type userController struct{}

func (*userController) Auth(c *fiber.Ctx) error {
	user := c.Locals("user").(model.User)
	roles := c.Locals("roles").([]string)
	permissions := c.Locals("permissions").([]string)

	id, _ := helper.Hash.EncodeId(user.Id)
	return helper.Res.SendData(
		c,
		lang.L.Convert(lang.L.Get().RETRIEVED_SUCCESSFULLY, fiber.Map{"operator": lang.L.Get().USER}),
		&resource_user.Show{
			Id:          id,
			Username:    user.Username,
			Name:        user.Name,
			Roles:       roles,
			Permissions: permissions,
			CreatedAt:   helper.Time2Str(user.CreatedAt),
			UpdatedAt:   helper.Time2Str(user.UpdatedAt),
		},
	)
}

func (*userController) Index(c *fiber.Ctx) error {
	query := helper.Req.QueryStr(c)

	rolesSubQuery := config.G.
		Table("roles as r").
		Select("COALESCE(json_agg(DISTINCT r.name), '[]')").
		Joins("left join user_has_roles as uhr on uhr.role_id = r.id").
		Where("uhr.user_id = u.id")

	rolesSubQuery2 := config.G.
		Table("roles as r").
		Select("DISTINCT r.name").
		Joins("left join user_has_roles as uhr on uhr.role_id = r.id").
		Where("uhr.user_id = u.id")

	permissionsSubQuery := config.G.
		Table("permissions as p").
		Select("COALESCE(json_agg(DISTINCT  p.name), '[]')").
		Joins("left join role_has_permissions as rhp on rhp.permission_id = p.id").
		Joins("left join roles as r on r.id = rhp.role_id").
		Joins("left join user_has_roles as uhr on uhr.user_id = u.id").
		Where("uhr.user_id = u.id and r.name in (?)", rolesSubQuery2)

	users := &[]map[string]any{}
	if err := config.G.Table("users as u").
		Select("u.*, (?) as roles, (?) as permissions", rolesSubQuery, permissionsSubQuery).
		Limit(query.Limit).
		Offset(helper.Req.Offset(query.Page, query.Limit)).
		Scan(users).Error; err != nil {
		return helper.Res.SendException(c, err)
	}

	newUsers := []resource_user.Show{}
	for _, v := range *users {
		id, _ := helper.Hash.EncodeId(int(v["id"].(int64)))

		var convertRoles []string
		if err := json.Unmarshal([]byte(v["roles"].(string)), &convertRoles); err != nil {
			convertRoles = []string{}
		}

		var convertPermissions []string
		if err := json.Unmarshal([]byte(v["permissions"].(string)), &convertPermissions); err != nil {
			convertPermissions = []string{}
		}

		newUsers = append(newUsers, resource_user.Show{
			Id:          id,
			Username:    v["username"].(string),
			Name:        v["name"].(string),
			Roles:       convertRoles,
			Permissions: convertPermissions,
			CreatedAt:   helper.Time2Str(v["created_at"].(time.Time)),
			UpdatedAt:   helper.Time2Str(v["updated_at"].(time.Time)),
		})
	}

	var countUsers int64
	config.G.
		Model(&model.User{}).
		Where("username like ?", "%"+query.Keyword+"%").
		Or("name like ?", "%"+query.Keyword+"%").
		Count(&countUsers)

	pagination := helper.Paginate{
		Page:  query.Page,
		Limit: query.Limit,
		Total: countUsers,
	}

	return helper.Res.SendDatas(
		c,
		lang.L.Convert(lang.L.Get().RETRIEVED_SUCCESSFULLY, fiber.Map{"operator": lang.L.Get().USER}),
		newUsers,
		pagination,
	)
}

func (*userController) Show(c *fiber.Ctx) error {
	query := helper.Req.QueryStr(c)

	rolesSubQuery := config.G.
		Table("roles as r").
		Select("COALESCE(json_agg(DISTINCT r.name), '[]')").
		Joins("left join user_has_roles as uhr on uhr.role_id = r.id").
		Where("uhr.user_id = u.id")

	rolesSubQuery2 := config.G.
		Table("roles as r").
		Select("DISTINCT r.name").
		Joins("left join user_has_roles as uhr on uhr.role_id = r.id").
		Where("uhr.user_id = u.id")

	permissionsSubQuery := config.G.
		Table("permissions as p").
		Select("COALESCE(json_agg(DISTINCT  p.name), '[]')").
		Joins("left join role_has_permissions as rhp on rhp.permission_id = p.id").
		Joins("left join roles as r on r.id = rhp.role_id").
		Joins("left join user_has_roles as uhr on uhr.user_id = u.id").
		Where("uhr.user_id = u.id and r.name in (?)", rolesSubQuery2)

	getParam := c.Params("id")
	user_id, err := helper.Hash.DecodeId(getParam)
	if err != nil {
		helper.Res.SendErrorMsg(c, lang.L.Convert(lang.L.Get().NOT_FOUND, fiber.Map{"operator": lang.L.Get().USER}))
	}

	users := &[]map[string]any{}
	if err := config.G.Table("users as u").
		Select("u.*, (?) as roles, (?) as permissions", rolesSubQuery, permissionsSubQuery).
		Where("id = ?", user_id).
		Scan(users).Error; err != nil {
		return helper.Res.SendException(c, err)
	}

	newUsers := []resource_user.Show{}
	for _, v := range *users {
		id, _ := helper.Hash.EncodeId(int(v["id"].(int64)))

		var convertRoles []string
		if err := json.Unmarshal([]byte(v["roles"].(string)), &convertRoles); err != nil {
			convertRoles = []string{}
		}

		var convertPermissions []string
		if err := json.Unmarshal([]byte(v["permissions"].(string)), &convertPermissions); err != nil {
			convertPermissions = []string{}
		}

		newUsers = append(newUsers, resource_user.Show{
			Id:          id,
			Username:    v["username"].(string),
			Name:        v["name"].(string),
			Roles:       convertRoles,
			Permissions: convertPermissions,
			CreatedAt:   helper.Time2Str(v["created_at"].(time.Time)),
			UpdatedAt:   helper.Time2Str(v["updated_at"].(time.Time)),
		})
	}

	var countUsers int64
	config.G.
		Model(&model.User{}).
		Where("username like ?", "%"+query.Keyword+"%").
		Or("name like ?", "%"+query.Keyword+"%").
		Count(&countUsers)

	return helper.Res.SendData(
		c,
		lang.L.Convert(lang.L.Get().RETRIEVED_SUCCESSFULLY, fiber.Map{"operator": lang.L.Get().USER}),
		newUsers,
	)
}

func (*userController) Update(c *fiber.Ctx) error {
	validated := &request_user.Update{}
	if err, isOk := helper.Validate(c, validated); !isOk {
		return err
	}

	getParam := c.Params("id")
	user_id, err := helper.Hash.DecodeId(getParam)
	if err != nil {
		return helper.Res.SendException(c, err)
	}

	user := &model.User{}
	if err := repository.User.First(c, user_id, user); err != nil {
		return helper.Res.SendException(c, err)
	}
	if user.Id == 0 {
		return helper.Res.SendErrorMsg(c, lang.L.Convert(lang.L.Get().NOT_FOUND, fiber.Map{"operator": lang.L.Get().USER}))
	}

	user.Name = validated.Name
	if err := repository.User.Update(user); err != nil {
		return helper.Res.SendException(c, err)
	}

	roles := &[]string{}
	if err := repository.Role.GetMany(user.Id, roles); err != nil {
		return helper.Res.SendException(c, err)
	}

	permissions := &[]string{}
	if err := repository.Permission.GetManyByUserId(user.Id, roles); err != nil {
		return helper.Res.SendException(c, err)
	}

	id, _ := helper.Hash.EncodeId(user.Id)
	newUser := resource_user.Show{
		Id:          id,
		Username:    user.Username,
		Name:        user.Name,
		Roles:       *roles,
		Permissions: *permissions,
		CreatedAt:   helper.Time2Str(user.CreatedAt),
		UpdatedAt:   helper.Time2Str(user.UpdatedAt),
	}

	return helper.Res.SendData(c, lang.L.Convert(lang.L.Get().UPDATED_SUCCESSFULLY, fiber.Map{"operator": lang.L.Get().USER}), newUser)
}
