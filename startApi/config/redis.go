package config

type RedisConf struct {
	Host string `default:"192.168.7.204"`
	Port int    `default:"26379"`
	DB   int    `default:"6"`
}
