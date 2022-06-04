package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type ConfigBot struct {
	TokenD   string `yaml:"token_d" env:"TOKEN_DS"`
	TokenT   string `yaml:"token_t" env:"TOKEN_TG"`
	Port     string `yaml:"port" env:"PORT" env-default:"3306"`
	Host     string `yaml:"host" env:"HOST" env-default:"127.0.0.1"`
	Name     string `yaml:"name" env:"NAME" env-default:"rsbot"`
	User     string `yaml:"user" env:"USER" env-default:"root"`
	Password string `yaml:"password" env:"PASSWORD"env-default:"root"`
}

var cfg ConfigBot

func InitConfig() ConfigBot {
	err := cleanenv.ReadConfig(".env", &cfg)
	if err != nil {
		panic(err)

	}
	return cfg

}
