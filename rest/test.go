package main

import (
	"fmt"
	"github.com/lexsalg/goweb/rest/models"
)

func main() {
	models.CreateConnection()
	models.Ping()
	models.CreateTables()

	models.CreateUser("alexis1", "123", "xel.salg@gmail.com")
	models.CreateUser("alexis2", "123", "xel.salg@gmail.com")
	models.CreateUser("alexis3", "123", "xel.salg@gmail.com")

	users := models.GetUsers()
	fmt.Println(users)

	models.CloseConnection()

}
