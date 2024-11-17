package routes

import (
	"flag"
	"fmt"
	"go-fiber-react/app/console"
	"go-fiber-react/config"
	"os"
)

func Console() {
	if config.APP_ENV != "development" {
		fmt.Println("Cannot access cmd while mode production")
		os.Exit(1)
	}

	dropFlag := flag.Bool("drop", false, "Drop the database tables")
	seedFlag := flag.Bool("seed", false, "Seed the database with initial data")
	migrateFlag := flag.Bool("migrate", false, "Run database migrations")

	flag.Parse()
	status := false

	cdb := console.DB{}

	if *dropFlag {
		cdb.Drop(config.G)
		status = true
	}

	if *migrateFlag {
		cdb.Migrate(config.G)
		status = true
	}

	if *seedFlag {
		cdb.Seed(config.G)
		status = true
	}

	if status {
		os.Exit(0)
	}

}
