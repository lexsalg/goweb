package models

import (
	"database/sql"
	"fmt"
	"github.com/lexsalg/goweb/rest/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var debug bool

func init() {
	CreateConnection()
	debug = config.GetDebug()
}

func CreateConnection() {

	if GetConnection() != nil {
		return
	}

	dsn := config.GetDsnDB()
	if conn, err := sql.Open("mysql", dsn); err != nil {
		panic(err)
	} else {
		db = conn
		fmt.Println("db open connection")
	}
}
func CloseConnection() {
	_ = db.Close()
	fmt.Println("db close connection")
}

func GetConnection() *sql.DB {
	return db
}

func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

func CreateTables() {
	createTable("users", userSchema)
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil && !debug {
		log.Println(err)
	}
	return result, err
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	if err != nil && !debug {
		log.Println(err)
	}
	return rows, err
}

func InsertData(query string, args ...interface{}) (int64, error) {
	if result, err := Exec(query, args...); err != nil {
		return int64(0), err
	} else {
		id, err := result.LastInsertId()
		return id, err
	}
}

/*---------------------------------------------------------------------------------------------*/

func existsTable(tableName string) bool {
	query := fmt.Sprintf("SHOW TABLES LIKE '%s' ", tableName)
	rows, _ := Query(query)
	return rows.Next()
}

func createTable(tableName, schema string) {
	if !existsTable(tableName) {
		_, _ = Exec(schema)
	} else {
		truncateTable(tableName) // esto borrar, porque sino tus tabasl se van a borrar
	}

}

func truncateTable(tableName string) {
	query := fmt.Sprintf("TRUNCATE %s", tableName)
	_, _ = Exec(query)

}

/*---------------------------------------------------------------------------------------------*/
