package main

import (
	"fmt"
	"github.com/lexsalg/goweb/rest/orm"
)

func main() {

	orm.CreateConnection()
	orm.CreateTables()

	user := orm.NewUser("alexis orm", "123", "xel.salg@gmail.com")
	user.Save()

	users := orm.GetUsers()
	fmt.Println(users)

	user = orm.GetUser(1)
	user.Username = "cambio nombre"
	user.Save()
	fmt.Println(user)
	user.Delete()
	orm.CloseConnection()

}
