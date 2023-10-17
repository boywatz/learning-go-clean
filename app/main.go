package main

import (
	"boywatz/go-clean/databases"
	"boywatz/go-clean/deliveries/routes"
	"boywatz/go-clean/models"
	"fmt"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var err error

func main() {
	databases.DB, err = gorm.Open("mysql", databases.ConnectionString(databases.BuildDBConfig()))

	if err != nil {
		fmt.Println("status: ", err)
	}

	defer databases.DB.Close()
	databases.DB.AutoMigrate(&models.Todo{})

	log.Info("This is sample log with logrus.")
	log.Info("add this line for test about pre-commit git hook.")

	r := routes.SetupRouter()
	if err := r.Run(":9000"); err != nil {
		log.Error(err)
	}
}
