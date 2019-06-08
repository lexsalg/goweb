package test

import (
	"fmt"
	"github.com/lexsalg/goweb/rest/models"
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	beforeTest()
	result := m.Run() //ejecuta todas las pruebas ubitariass
	afterTest()
	os.Exit(result)
}

func beforeTest() {
	//fmt.Println("antes de las pruebas")
	models.CreateConnection()
	models.CreateTables()
	createDefaultUser()
}
func afterTest() {
	//fmt.Println("despues de las pruebas")
	models.CloseConnection()
}

func createDefaultUser() {
	sql := fmt.Sprintf("INSERT users SET id='%d', username='%s', password='%s' ,email='%s',created_date='%s'", id, username, passwordHash, email, createdDate)
	if _, err := models.Exec(sql); err != nil {
		panic(err)
	}
	user = &models.User{
		Id:       id,
		Username: username,
		Password: password,
		Email:    email,
	}
}
