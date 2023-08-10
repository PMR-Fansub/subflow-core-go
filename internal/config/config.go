package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Server struct {
		Type       string `yaml:"type" env:"SERVER_TYPE" env-default:"dev"`
		IpAddr     string `yaml:"ip" env:"SERVER_IP" env-default:"localhost"`
		Port       uint16 `yaml:"port" env:"SERVER_PORT" env-default:"8000"`
		SigningKey string `yaml:"signing-key" env:"SIGNING_KEY" env-default:"secret"`
		CookieHost string `yaml:"cookie-host" env:"COOKIE_HOST" env-default:"http://localhost:3000"`
	} `yaml:"server"`
	Datasource struct {
		Type            string `yaml:"type"`
		DSN             string `yaml:"dsn"`
		MaxIdleConn     int    `yaml:"max-idle-conn" env-default:"1"`
		MaxOpenConn     int    `yaml:"max-open-conn" env-default:"1"`
		ConnMaxLifeTime int    `yaml:"conn-max-life-time" env-default:"1"` // Hour
	} `yaml:"datasource"`
}

const (
	ServerTypeDev        = "dev"
	ServerTypeProd       = "prod"
	DatasourceTypeSQLite = "sqlite"
	DatasourceTypeMySQL  = "mysql"
)

const (
	DefaultConfigPath = "config.yaml"
)

func New() (*Config, error) {
	var cfg Config
	err := cleanenv.ReadConfig(DefaultConfigPath, &cfg)
	return &cfg, err
}
