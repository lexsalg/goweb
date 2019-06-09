package config

import (
	"fmt"
)

type Config interface {
	url() string
}

type ServerConfig struct {
	host  string
	port  int
	debug bool
}

type DBConfig struct {
	username string
	password string
	host     string
	port     int
	database string
	debug    bool
}

var server *ServerConfig
var database *DBConfig

func init() {

	server = &ServerConfig{}
	server.host = StringEnv("HOST", "localhost")
	server.port = IntEnv("PORT", 5000)
	server.debug = BoolEnv("DEBUG", true)

	database = &DBConfig{}
	database.username = StringEnv("DBUSERNAME", "root")
	database.password = StringEnv("DBPASSWORD", "root")
	database.host = StringEnv("DBHOST", "localhost")
	database.port = IntEnv("DBPORT", 3306)
	database.database = StringEnv("DBNAME", "goweb")
	database.debug = BoolEnv("DBDEBUG", true)
}

func (s *ServerConfig) url() string {
	return fmt.Sprintf("%s:%d", s.host, s.port)
}

func (c *DBConfig) url() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", c.username, c.password, c.host, c.port, c.database)
}

func UrlServer() string {
	return server.url()
}

func ServerPort() int {
	return server.port
}

func DsnDB() string {
	return database.url()
}

func Debug() bool {
	return server.debug
}

func DirTemplate() string {
	return "templates/**/*.html"

}

func DirTemplateError() string {
	return "templates/error.html"
}

func DirAssets() string {
	return "assets"
}

func PrefixAssets() string {
	return "/assets/"

}
