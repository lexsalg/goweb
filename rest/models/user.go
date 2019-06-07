package models

import "github.com/pkg/errors"

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

const userSchema string = `CREATE TABLE users(
    	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    	username VARCHAR(30) NOT NULL,
    	password VARCHAR(64) NOT NULL,
    	email VARCHAR(40),
    	created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP )`

type Users []User

var users = make(map[int]User)

func SetDefaultUser() {
	user := User{1, "lexsalg", "123"}
	users[user.Id] = user
}

func GetUsers() Users {
	list := Users{}
	for _, user := range users {
		list = append(list, user)
	}
	return list
}

func GetUser(id int) (User, error) {
	if user, ok := users[id]; ok {
		return user, nil
	}
	return User{}, errors.New("Usuario No existe en map.")
}

// SaveUser Guarda un usuario
func SaveUser(user User) User {
	user.Id = len(users) + 1
	users[user.Id] = user
	return user
}

// UpdateUser Actualiza un usuario
func UpdateUser(user User, username, password string) User {
	user.Username = username
	user.Password = password
	users[user.Id] = user
	return user
}

// DeleteUser Elimina un usuario
func DeleteUser(userId int) {
	delete(users, userId)
}
