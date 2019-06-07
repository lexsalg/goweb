package main

import (
	"github.com/lexsalg/goweb/rest/models"
)

func main() {
	models.CreateConnection()
	models.Ping()
	models.CreateTables()
	models.CloseConnection()

}
