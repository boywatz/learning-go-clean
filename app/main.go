package main

import (
	"boywatz/go-clean/databases"
	"boywatz/go-clean/deliveries/routes"
	"boywatz/go-clean/models"
	"fmt"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	databases.DB, err = gorm.Open("mysql", databases.ConnectionString(databases.BuildDBConfig()))

	if err != nil {
		fmt.Println("status: ", err)
	}

	defer databases.DB.Close()
	databases.DB.AutoMigrate(&models.Todo{})

	r := routes.SetupRouter()
	r.Run(":9000")
}
