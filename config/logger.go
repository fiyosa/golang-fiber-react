package config

import (
	"fmt"
	"io"
	"os"

	"github.com/gofiber/fiber/v2/log"
)

const out = "console" // "file" or "console"

// example:
// config.log("info")
// config.logf("info %v", "data")

func Log(v ...interface{}) {
	if APP_ENV == "development" {
		log.Info(v...)
	}
}

func Logf(format string, v ...interface{}) {
	if APP_ENV == "development" {
		log.Infof(format, v...)
	}
}

func Logger() {
	if APP_ENV == "development" {
		dirLogs := "./storage/logs"

		_, err := os.Stat(dirLogs)
		if os.IsNotExist(err) {
			err := os.MkdirAll(dirLogs, 0755)
			if err != nil {
				fmt.Printf("Error creating directory logs: %v \n\n", err.Error())
				os.Exit(0)
			}
		}

		file, err := os.OpenFile(dirLogs+"/fiber.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			fmt.Printf("Error opening file: %v \n\n", err.Error())
			os.Exit(0)
		}

		if out == "file" {
			log.SetOutput(file)
		} else {
			iw := io.MultiWriter(os.Stdout, file)
			log.SetOutput(iw)
		}
	}
}
