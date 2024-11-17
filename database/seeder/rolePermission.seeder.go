package seeder

import (
	"fmt"
	"go-fiber-react/app/model"
	"os"

	"gorm.io/gorm"
)

var roles = []string{
	"admin",
	"user",
}

var permissions = []string{
	"user_index",
	"user_show",
}

func RolePermissionSeeder(g *gorm.DB) {
	tx := g.Begin()

	createRoles := []*model.Role{}
	for _, v := range roles {
		createRoles = append(createRoles, &model.Role{Name: v})
	}

	createPermissions := []*model.Permission{}
	for _, v := range permissions {
		createPermissions = append(createPermissions, &model.Permission{Name: v})
	}

	if err := tx.Create(&createRoles).Error; err != nil {
		tx.Rollback()
		fmt.Printf("Error seeder role: %v \n\n", err.Error())
		os.Exit(1)
	}
	if err := tx.Create(&createPermissions).Error; err != nil {
		tx.Rollback()
		fmt.Printf("Error seeder permission: %v \n\n", err.Error())
		os.Exit(1)
	}

	createRoleHasPermissions := []*model.RoleHasPermission{}

	createRoleHasPermissions = append(createRoleHasPermissions, createAdmin(createRoles, createPermissions)...)
	createRoleHasPermissions = append(createRoleHasPermissions, createUser(createRoles, createPermissions)...)

	if err := tx.Create(&createRoleHasPermissions).Error; err != nil {
		tx.Rollback()
		fmt.Printf("Error seeder role has permission: %v \n\n", err.Error())
		os.Exit(1)
	}

	if err := tx.Commit().Error; err != nil {
		fmt.Println("Transaction commit failed role permission:", err)
		return
	}

	fmt.Println("Seeder: role permission created successfully.")
}

func createAdmin(cr []*model.Role, cp []*model.Permission) []*model.RoleHasPermission {
	roleName := "admin"

	var roleID int
	for _, v := range cr {
		if v.Name == roleName {
			roleID = v.Id
			break
		}
	}

	crhp := []*model.RoleHasPermission{}
	for _, v := range cp {
		crhp = append(crhp, &model.RoleHasPermission{
			RoleId:       roleID,
			PermissionId: v.Id,
		})
	}
	return crhp
}

func createUser(cr []*model.Role, cp []*model.Permission) []*model.RoleHasPermission {
	roleName := "user"
	permissions := []string{
		"user_show",
	}

	var roleID int
	for _, v := range cr {
		if v.Name == roleName {
			roleID = v.Id
			break
		}
	}

	crhp := []*model.RoleHasPermission{}
	for _, v := range cp {
		for _, p := range permissions {
			if p == v.Name {
				crhp = append(crhp, &model.RoleHasPermission{
					RoleId:       roleID,
					PermissionId: v.Id,
				})
			}
		}
	}
	return crhp
}
