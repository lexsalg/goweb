package models

import (
	"database/sql"
	"fmt"
	"github.com/lexsalg/goweb/rest/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func CreateConnection() {
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
	if err != nil {
		log.Println(err)
	}
	return result, err
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		log.Println(err)
	}
	return rows, err
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
		truncateTable(tableName)
	}

}

func truncateTable(tableName string) {
	query := fmt.Sprintf("TRUNCATE %s", tableName)
	_, _ = Exec(query)

}

/*---------------------------------------------------------------------------------------------*/
