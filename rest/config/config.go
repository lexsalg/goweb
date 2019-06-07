package config

import (
	"fmt"
	"github.com/eduardogpg/gonv"
)

type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
}

var database *DBConfig

func init() {
	database = &DBConfig{}
	database.Username = gonv.GetStringEnv("USERNAME", "root")
	database.Password = gonv.GetStringEnv("PASSWORD", "root")
	database.Host = gonv.GetStringEnv("HOST", "localhost")
	database.Port = gonv.GetIntEnv("PORT", 3306)
	database.Database = gonv.GetStringEnv("DATABASE", "goweb")
}

func (c *DBConfig) dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", c.Username, c.Password, c.Host, c.Port, c.Database)
}

func GetDsnDB() string {
	return database.dsn()
}
