package models

import "github.com/pkg/errors"

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserXml struct {
	Id       int    `xml:"id"`
	Username string `xml:"username"`
	Password string `xml:"password"`
}

var users = make(map[int]User)

func SetDefaultUser() {
	user := User{1, "lexsalg", "123"}
	users[user.Id] = user
}

func GetUser(id int) (User, error) {
	if user, ok := users[id]; ok {
		return user, nil
	}
	return User{}, errors.New("Usuario No existe en map.")
}
