package config

type RpcConf struct {
	Host string `default:"0.0.0.0"`
	Port int    `default:"9311"`
}

