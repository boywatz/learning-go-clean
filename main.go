package main

import (
	"boywatz/go-clean/databases"
	"boywatz/go-clean/deliveries/routes"
	"boywatz/go-clean/models"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var err error

func main() {
	log.Info("This is sample log with logrus.")

	appEnv := os.Getenv("APP_ENV")
	log.Info("environment value_" + appEnv)

	databases.DB, err = gorm.Open("mysql", databases.ConnectionString(databases.BuildDBConfig()))

	if err != nil {
		fmt.Println("database error: ", err)
	}

	defer databases.DB.Close()
	databases.DB.AutoMigrate(&models.Todo{})

	r := routes.SetupRouter()
	if err := r.Run(":9000"); err != nil {
		log.Error(err)
	}
}
