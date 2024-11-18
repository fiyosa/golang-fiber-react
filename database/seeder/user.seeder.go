package seeder

import (
	"fmt"
	"go-fiber-react/app/helper"
	"go-fiber-react/app/model"
	"os"

	"gorm.io/gorm"
)

func UserSeeder(g *gorm.DB) {

	password, _ := helper.Hash.Create("Password")
	tx := g.Begin()

	users := []*model.User{
		{Username: "admin", Name: "Admin", Password: password},
		{Username: "user", Name: "User", Password: password},
	}

	if err := tx.Create(&users).Error; err != nil {
		tx.Rollback()
		fmt.Printf("Error seeder users: %v \n\n", err.Error())
		os.Exit(1)
	}

	roles := []*model.Role{}
	g.Find(&roles)

	userHasRoles := []*model.UserHasRole{}
	for _, u := range users {
		for _, r := range roles {
			if u.Username == "admin" && r.Name == "admin" {
				userHasRoles = append(userHasRoles, &model.UserHasRole{UserId: u.Id, RoleId: r.Id})
			}
			if u.Username == "user" && r.Name == "user" {
				userHasRoles = append(userHasRoles, &model.UserHasRole{UserId: u.Id, RoleId: r.Id})
			}
		}
	}

	if err := tx.Create(&userHasRoles).Error; err != nil {
		tx.Rollback()
		fmt.Printf("Error seeder user has permission: %v \n\n", err.Error())
		os.Exit(1)
	}

	if err := tx.Commit().Error; err != nil {
		fmt.Println("Transaction commit failed user:", err)
		return
	}

	fmt.Println("Seeder: user created successfully.")
}
