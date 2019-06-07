package main

import "github.com/lexsalg/goweb/rest/models"

func main() {
	//_ = os.Setenv("HOST", "localhost")
	//_ = os.Unsetenv("HOST" +
	//	"")
	//env := os.Getenv("HOST")
	//fmt.Println(env)
	models.CreateConnection()
	models.Ping()
	models.CloseConnection()

}
