package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	log "github.com/sirupsen/logrus"
	"sync"
)

//	type ConfigBot struct {
//		TokenD    string `yaml:"token_d" env:"TOKEN_DS"`
//		TokenT    string `yaml:"token_t" env:"TOKEN_TG"`
//		LogToken  string `yaml:"logToken" env:"LOGTOKEN"`
//		LogChatId int64  `yaml:"logChatId" env:"LOGCHATID"`
//
//		DBHostname   string `yaml:"dbhostname" env:"DBHOSTNAME" env-default:"127.0.0.1:3306"`
//		Dbname       string `yaml:"dbname" env:"DBNAME" env-default:"rsbot"`
//		Dbusername   string `yaml:"dbusername" env:"DBUSERNAME" env-default:"root"`
//		DbPassword   string `yaml:"dbpassword" env:"DBPASSWORD"env-default:"root"`
//		DBmode       string `yaml:"DBmode"  env:"DBMODE" env-default:"postgres"`
//		BotMode      string `yaml:"botMode" env:"BOTMOD" env-default:"server"` //reserve || server
//		ServerAdrr   string `yaml:"serverAdrr" env:"SERVERADRR" env-default:"braut.com.ua:7733"`
//		Debug        bool   `yaml:"debug" env:"DEBUG" env-default:"false"`
//		SupabasePass string `yaml:"supabasePass" env:"SUPABASEPASS"`
//	}
type ConfigBot struct {
	IsDebug    bool   `yaml:"is_debug" env-default:"false"`
	BotMode    string `yaml:"bot_mode" env-default:"server"` //reserve || server
	ServerAdrr string `yaml:"server_adrr"  env-default:"braut.com.ua:7733"`
	Token      struct {
		TokenDiscord   string `yaml:"token_discord"`
		TokenTelegram  string `yaml:"token_telegram"`
		NameDbWhatsapp string `yaml:"name_db_whatsapp"`
	} `yaml:"token"`
	Logger struct {
		Token  string `yaml:"token"`
		ChatId int64  `yaml:"chat_id"`
	} `yaml:"logger"`
	Postgress struct {
		Host     string `yaml:"host" env-default:"127.0.0.1:3306"`
		Name     string `yaml:"name" env-default:"rsbot"`
		Username string `yaml:"username" env-default:"root"`
		Password string `yaml:"password" env-default:"root"`
	} `yaml:"postgress"`
	Supabase struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Name     string `yaml:"name"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
}

var Instance *ConfigBot
var once sync.Once

func InitConfig() *ConfigBot {
	once.Do(func() {
		Instance = &ConfigBot{}
		err := cleanenv.ReadConfig("config/config.yml", Instance)
		if err != nil {
			help, _ := cleanenv.GetDescription(Instance, nil)
			log.Println(help)
			log.Fatal(err)
		}
	})
	return Instance
}
