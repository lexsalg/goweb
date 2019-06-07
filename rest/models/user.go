package models

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

const userSchema string = `CREATE TABLE users(
    	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    	username VARCHAR(30) NOT NULL,
    	password VARCHAR(64) NOT NULL,
    	email VARCHAR(40),
    	created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP )`

type Users []User

// constructor
func NewUser(username, password, email string) *User {
	user := &User{Username: username, Password: password, Email: email}
	return user
}

func CreateUser(username, password, email string) *User {
	user := NewUser(username, password, email)
	user.Save()
	return user
}

func GetUsers() Users {
	users := Users{}
	sql := "SELECT id, username, password, email FROM users"
	rows, _ := Query(sql)

	for rows.Next() {
		o := User{}
		_ = rows.Scan(&o.Id, &o.Username, &o.Password, &o.Email)
		users = append(users, o)
	}
	return users
}

func GetUser(id int) *User {
	u := NewUser("", "", "")
	sql := "SELECT id, username, password, email FROM users WHERE id=?"
	rows, _ := Query(sql, id)

	for rows.Next() {
		_ = rows.Scan(&u.Id, &u.Username, &u.Password, &u.Email)
	}

	return u
}

func (u *User) Save() {
	if u.Id == 0 {
		u.insert()
	} else {
		u.update()
	}
}

func (u *User) Delete() {
	sql := `DELETE FROM users WHERE id=?`
	_, _ = Exec(sql, u.Id)
}

func (u *User) insert() {
	sql := `INSERT users SET username=?, password=?, email=?`
	res, _ := Exec(sql, u.Username, u.Password, u.Email)
	u.Id, _ = res.LastInsertId()
}

func (u *User) update() {
	sql := `UPDATE users SET username=?, password=?, email=?`
	_, _ = Exec(sql, u.Username, u.Password, u.Email)
	//res, _ := Exec(sql, u.Username, u.Password, u.Email)
	//u.Id, _ = res.LastInsertId()
}
