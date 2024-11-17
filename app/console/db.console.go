package console

import (
	"fmt"
	"go-fiber-react/app/model"
	"go-fiber-react/database/seeder"
	"os"

	"gorm.io/gorm"
)

type DB struct{}

var models = []interface{}{
	&model.User{},
	&model.Auth{},
	&model.Role{},
	&model.Permission{},
	&model.UserHasRole{},
	&model.RoleHasPermission{},
}

func (*DB) Seed(g *gorm.DB) {
	if err := seeder.Seed(g); err != nil {
		fmt.Printf("Error seeder: %v \n\n", err.Error())
		os.Exit(1)
	}
	fmt.Println("Seeder successfully")
}

func (*DB) Migrate(g *gorm.DB) {
	if err := g.AutoMigrate(models...); err != nil {
		fmt.Printf("Error migration: %v \n\n", err.Error())
		os.Exit(1)
	}
	fmt.Println("Migrate successfully")
}

func (*DB) Drop(g *gorm.DB) {
	if err := g.Migrator().DropTable(models...); err != nil {
		fmt.Printf("Error drop all table: %v \n\n", err.Error())
		os.Exit(1)
	}
	fmt.Println("Deleted all table successfully.")
}
