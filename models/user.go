package models

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"time"
)

type User struct {
	Id          int64  `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	createdDate time.Time
}

const userSchema string = `CREATE TABLE users(
    	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    	username VARCHAR(30) NOT NULL UNIQUE,
    	password VARCHAR(60) NOT NULL,
    	email VARCHAR(40),
    	created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP )`

var emailRegExp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type Users []User

// constructor
func NewUser(username, password, email string) (*User, error) {
	user := &User{Username: username, Email: email}
	if err := user.Valid(); err != nil {
		return &User{}, err
	}
	err := user.SetPassword(password)
	return user, err
}

func CreateUser(username, password, email string) (*User, error) {
	user, err := NewUser(username, password, email)
	if err != nil {
		return &User{}, err
	}
	err = user.Save()
	return user, err
}

func Login(username, password string) bool {
	user := GetUserByUsername(username)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

func ValidUsername(username string) error {
	if len(username) > 30 {
		return errors.New("Username debe tener maximo 30 caracteres")
	}
	return nil
}

func ValidEmail(email string) error {
	if !emailRegExp.MatchString(email) {
		return errors.New("Email no valido")
	}
	return nil
}

func GetUserByUsername(username string) *User {
	sql := "SELECT id, username, password, email, created_date FROM users WHERE username=?"
	return GetUser(sql, username)
}

func GetUserById(id int) *User {
	sql := "SELECT id, username, password, email, created_date FROM users WHERE id=?"
	return GetUser(sql, id)
}

func GetUser(sql string, condition interface{}) *User {
	u := &User{}
	rows, err := Query(sql, condition)
	if err != nil {
		return u
	}

	for rows.Next() {
		_ = rows.Scan(&u.Id, &u.Username, &u.Password, &u.Email, &u.createdDate)
	}
	return u
}

func GetUsers() Users {
	users := Users{}
	sql := "SELECT id, username, password, email, created_date FROM users"
	rows, _ := Query(sql)

	for rows.Next() {
		o := User{}
		_ = rows.Scan(&o.Id, &o.Username, &o.Password, &o.Email, &o.createdDate)
		users = append(users, o)
	}
	return users
}

func (u *User) Valid() error {
	if err := ValidEmail(u.Email); err != nil {
		return err
	}
	if err := ValidUsername(u.Username); err != nil {
		return err
	}
	return nil
}

func (u *User) Save() error {
	if u.Id == 0 {
		return u.insert()
	} else {
		return u.update()
	}
}

func (u *User) Delete() error {
	sql := `DELETE FROM users WHERE id=?`
	_, err := Exec(sql, u.Id)
	return err
}

func (u *User) insert() error {
	sql := `INSERT users SET username=?, password=?, email=?`
	id, err := InsertData(sql, u.Username, u.Password, u.Email)
	u.Id = id
	return err
}

func (u *User) update() error {
	sql := `UPDATE users SET username=?, password=?, email=? WHERE id=?`
	_, err := Exec(sql, u.Username, u.Password, u.Email, u.Id)
	return err
}

func (u *User) GetCreatedData() time.Time {
	return u.createdDate
}

func (u *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("no es posible cifrar la contraseña")
	}
	u.Password = string(hash)
	return nil
}
