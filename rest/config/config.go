package config

import (
	"fmt"
	"github.com/eduardogpg/gonv"
)

type DBConfig struct {
	username string
	password string
	host     string
	port     int
	database string
	debug    bool
}

var database *DBConfig

func init() {
	database = &DBConfig{}
	database.username = gonv.GetStringEnv("USERNAME", "root")
	database.password = gonv.GetStringEnv("PASSWORD", "root")
	database.host = gonv.GetStringEnv("HOST", "localhost")
	database.port = gonv.GetIntEnv("PORT", 3306)
	database.database = gonv.GetStringEnv("DATABASE", "goweb")
	database.debug = gonv.GetBoolEnv("DEBUG", true)
}

func (c *DBConfig) dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", c.username, c.password, c.host, c.port, c.database)
}

func GetDsnDB() string {
	return database.dsn()
}

func GetDebug() bool {
	return database.debug
}
