package main

import (
	"fmt"
	"myapp/models"
)

func main() {
	db := models.GetDB()

	var dbModels []interface{}
	dbModels = append(dbModels, &models.User{})

	var err error
	for _, m := range dbModels {
		err = db.AutoMigrate(m)
		if err != nil {
			fmt.Println(err)
		}
	}
}
