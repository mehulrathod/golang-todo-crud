package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"updated_structure/awesomeGoProject/Config"
	"updated_structure/awesomeGoProject/Models"
	"updated_structure/awesomeGoProject/Routes"
)

var err error

func main()  {
	// Creating a connection to the database
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("statuse: ", err)
	}

	defer Config.DB.Close()

	// run the migrations: todo struct
	Config.DB.AutoMigrate(&Models.Todo{})

	//setup routes
	r := Routes.SetupRouter()

	// running
	r.Run()
}
