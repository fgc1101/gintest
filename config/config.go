package config

type Config struct {
	Server  Server  `yaml:"server" json:"app"`
	QQEmail QQEmail `yaml:"qq_email" json:"qq_email"`
}

type Server struct {
	Port int `yaml:"port" json:"port"`
}

type QQEmail struct {
	User     string `yaml:"user" json:"user"`
	AuthCode string `yaml:"authCode" json:"authCode"`
	Host     string `yaml:"host" json:"host"`
	Port     int    `yaml:"port" json:"port"`
}
