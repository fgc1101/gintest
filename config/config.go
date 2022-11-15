package config

type Config struct {
	Server     Server     `yaml:"server" json:"app"`
	DataSource DataSource `yaml:"data_source" json:"data_source"`
	QQEmail    QQEmail    `yaml:"qq_email" json:"qq_email"`
}

type DataSource struct {
	DriverName string `yaml:"driver_name" json:"driver_name"`
	Host       string `yaml:"host" json:"host"`
	Port       int    `yaml:"port" json:"port"`
	Database   string `yaml:"database" json:"database"`
	Username   string `yaml:"username" json:"username"`
	Password   string `yaml:"password" json:"password"`
	Charset    string `yaml:"charset" json:"charset"`
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
