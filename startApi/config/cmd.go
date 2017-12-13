package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-redis/redis"
)

type CmdConfig struct {
	ApiConf
	MysqlConf
	RedisConf
}

// DbConnection initialize the mysql
func (c *CmdConfig) DbConnection() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		c.MysqlConf.User,
		c.MysqlConf.Pass,
		c.MysqlConf.Host,
		c.MysqlConf.Port,
		c.MysqlConf.Name,
		c.MysqlConf.Charset,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(c.IdleConns)
	db.SetMaxOpenConns(c.OpenConns)

	return db, nil
}

// RedisConnection connection the redis
func (c *CmdConfig) RedisConnection() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.RedisConf.Host, c.RedisConf.Port),
		Password: "",
		DB:       c.RedisConf.DB,
	})

	return client
}
