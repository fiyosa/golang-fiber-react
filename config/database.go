package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var G *gorm.DB

func DB() {
	dsn := fmt.Sprintf(
		`host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=Asia/Jakarta`,
		APP_DB_HOST,
		APP_DB_USER,
		APP_DB_PASS,
		APP_DB_NAME,
		APP_DB_PORT,
		APP_DB_SSLMODE,
	)

	var setLogger logger.Interface
	if APP_ENV != "development" {
		setLogger = logger.Default.LogMode(logger.Silent)
	} else {
		setLogger = gormLogger()
	}

	connect, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: APP_DB_SCHEMA + ".",
			// SingularTable: true,
		},
		SkipDefaultTransaction: true,
		Logger:                 setLogger,
		NowFunc: func() time.Time {
			return time.Now().Local() // timestamps
		},
	})

	if err != nil {
		fmt.Printf("Error access database: %v \n\n", err.Error())
		os.Exit(1)
	}

	G = connect
}

func gormLogger() logger.Interface {
	dirLogs := "./storage/logs"

	_, err := os.Stat(dirLogs)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(dirLogs, 0755); err != nil {
			fmt.Printf("Error creating directory logs: %v \n\n", err.Error())
			os.Exit(1)
		}
	}

	file, err := os.OpenFile(dirLogs+"/gorm.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Error opening file: %v \n\n", err.Error())
		os.Exit(1)
	}

	return logger.New(
		log.New(file, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             0,           // Disable slow threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
}
