package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App AppConfig `yaml:"app"`
	DB  DBConfig  `yaml:"db"`
}

type AppConfig struct {
	Name     string      `yaml:"name"`
	AuthPort string      `yaml:"auth_port"`
	Creds    CredsConfig `yaml:"creds"`
	JWT      JWTConfig   `yaml:"jwt"`
}

type JWTConfig struct {
	Secret   string `yaml:"secret"`
	Duration int16  `yaml:"duration"`
}

type DBConfig struct {
	Host           string           `yaml:"host"`
	Port           string           `yaml:"port"`
	User           string           `yaml:"user"`
	Password       string           `yaml:"password"`
	Name           string           `yaml:"name"`
	SSLMode        string           `yaml:"ssl_mode"`
	ConnectionPool DBConnectionPool `yaml:"connection_pool"`
}

type CredsConfig struct {
	SlatPassword int `yaml:"salt_password"`
}

type DBConnectionPool struct {
	MaxOpenConnection int `yaml:"max_open_connection"`
	MaxIdleConnection int `yaml:"max_idle_connection"`

	// in seconds
	MaxLifeTime int `yaml:"max_life_time"`
	MaxIdleTime int `yaml:"max_idle_time"`
}

var Cfg *Config

func LoadConfig(filename string) (err error) {
	fileByte, err := os.ReadFile(filename)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(fileByte, &Cfg)
	if err != nil {
		return
	}

	if Cfg.DB.SSLMode == "" {
		Cfg.DB.SSLMode = "disable"
	}
	return
}
