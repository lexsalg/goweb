package test

import (
	"fmt"
	"github.com/VividCortex/mysqlerr"
	"github.com/go-sql-driver/mysql"
	"github.com/lexsalg/goweb/models"
	"math/rand"
	"testing"
	"time"
)

const (
	id           = 1
	username     = "alexis-test"
	password     = "pasword"
	passwordHash = "$2a$10$2dJ9CVxw02V0OUGun7txMOM9pGuP9uu/QflgSCDFDqyPheKnxGeb2"
	email        = "xel.salg@gmail.com"
	createdDate  = "2020-06-07"
)

var user *models.User

func randomUsername() string {
	return fmt.Sprintf("%s/%d", username, rand.Intn(1000))
}

func equalsUser(user *models.User) bool {
	return user.Username == username && user.Email == email
}

func equalsCreatedDate(date time.Time) bool {
	t, _ := time.Parse("2006-01-02", createdDate)
	return t == date
}

func TestNewUser(t *testing.T) {
	_, err := models.NewUser(username, password, email)
	if err != nil {
		t.Error("no es posible crear el objeto", nil)
	}
}
func TestValidEmail(t *testing.T) {
	if err := models.ValidEmail(email); err != nil {
		t.Error("Validacion erronea en el email", err)
	}
}

func TestInvalidEmail(t *testing.T) {
	if err := models.ValidEmail("ccddc*@@dfsdf.com"); err == nil {
		t.Error("Validacion erronea en el email")
	}
}

func TestUsernameLength(t *testing.T) {
	name := username
	for i := 0; i < 10; i++ {
		name += name
	}
	if _, err := models.NewUser(name, password, email); err == nil {
		t.Error("Es posible generar username mayor a 30 caracteres")
	}
}

func TestPassword(t *testing.T) {
	user, _ := models.NewUser(username, password, email)
	if user.Password == password || len(user.Password) != 60 {
		t.Error("no es posible cifrar el password", nil)
	}
}

func TestLogin(t *testing.T) {
	if valid := models.Login(username, password); !valid {
		t.Error("no es posible realizar el login", nil)
	}
}

func TestNoLogin(t *testing.T) {
	if valid := models.Login(randomUsername(), password); valid {
		t.Error("es posible realizar el login con parametros erroneos", nil)
	}
}

func TestSave(t *testing.T) {
	user, _ := models.NewUser(randomUsername(), password, email)
	if err := user.Save(); err != nil {
		t.Error("no es posible crear usuario", err)
	}
}

func TestCreateUser(t *testing.T) {
	if _, err := models.CreateUser(randomUsername(), password, email); err != nil {
		t.Error("no es posible insertar usuario", err)
	}
}

func TestUniqueUsername(t *testing.T) {
	if _, err := models.CreateUser(username, password, email); err == nil {
		t.Error("es posible insertar usuarios con usernames duplicados", err)
	}
}

func TestDuplicateUsername(t *testing.T) {
	_, err := models.CreateUser(username, password, email)
	//if err
	if driverErr, ok := err.(*mysql.MySQLError); ok {
		if driverErr.Number != mysqlerr.ER_DUP_ENTRY {
			t.Error("es posible tener username duplicado en la bd", err)
		}
	}
}

func TestGetUser(t *testing.T) {
	user := models.GetUserById(id)
	//t.Log(user.GetCreatedData())
	if !equalsUser(user) || !equalsCreatedDate(user.GetCreatedData()) {
		t.Error("no es posible obtener el usuario")
	}
}

func TestGetUsers(t *testing.T) {
	users := models.GetUsers()
	if len(users) == 0 {
		t.Error("no es posible obtener los usuarios")
	}
}

func TestDeleteUser(t *testing.T) {
	if err := user.Delete(); err != nil {
		t.Error("no es posible eliminar al usuario", err)
	}
}
