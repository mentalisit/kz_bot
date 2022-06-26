package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type ConfigBot struct {
	TokenD    string `yaml:"token_d" env:"TOKEN_DS"`
	TokenT    string `yaml:"token_t" env:"TOKEN_TG"`
	LogToken  string `yaml:"logToken" env:"LOGTOKEN"`
	LogChatId int64  `yaml:"logChatId" env:"LOGCHATID"`

	DBHostname string `yaml:"dbhostname" env:"DBHOSTNAME" env-default:"127.0.0.1:3306"`
	Dbname     string `yaml:"dbname" env:"DBNAME" env-default:"rsbot"`
	Dbusername string `yaml:"dbusername" env:"DBUSERNAME" env-default:"root"`
	DbPassword string `yaml:"dbpassword" env:"DBPASSWORD"env-default:"root"`
}

var cfg ConfigBot

func InitConfig() ConfigBot {
	err := cleanenv.ReadConfig(".env", &cfg)
	if err != nil {
		panic(err)

	}
	return cfg

}
