package orm

import "time"

type User struct {
	Id        int64     `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
type Users []User

func CreateUser(username, password, email string) *User {
	user := NewUser(username, password, email)
	user.Save()
	return user
}

func GetUsers() Users {
	users := Users{}
	db.Find(&users)
	return users
}

func GetUser(id int) *User {
	u := &User{}
	db.Where("id=?", id).First(u)
	return u
}

func NewUser(username, password, email string) *User {
	user := &User{Username: username, Password: password, Email: email}
	return user

}

func (u *User) Save() {
	if u.Id == 0 {
		db.Create(&u)
	} else {
		u.update()
	}
}

func (u *User) Delete() {
	db.Delete(&u)
}

func (u *User) update() {
	user := User{Username: u.Username, Password: u.Password, Email: u.Email}
	db.Model(&u).UpdateColumns(user)
}
