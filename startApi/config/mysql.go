package config

type MysqlConf struct {
	Host      string `default:"192.168.7.204"`
	Port      int    `default:"23306"`
	User      string `default:"root"`
	Pass      string `default:"123456"`
	Name      string `default:"test"`
	Charset   string `default:"utf8"`
	IdleConns int    `default:"200"`
	OpenConns int    `default:"500"`
}
