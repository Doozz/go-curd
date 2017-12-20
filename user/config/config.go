package config

import (
	"database/sql"
	"errors"

	"github.com/go-redis/redis"
)

var (
	RedisHandle *redis.Client
	DBHandle    *sql.DB
	Config      *CmdConfig
)

// InitializeConn to initialize db and redis
func InitializeConn() (err error) {
	DBHandle, err = Config.DbConnection()
	if err != nil {
		return errors.New("Config.DbConnection() failed. Error info: " + err.Error())
	}

	RedisHandle = Config.RedisConnection()
	return nil
}
