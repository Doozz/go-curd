package config

// ApiConf is stored the api listen host and port
type ApiConf struct {
	Host string `default:""`
	Port int    `default:"9090"`
}
