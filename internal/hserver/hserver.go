package hserver

/*
import (
	"net/http"

	"kz_bot/pkg/logrus"
)

type ServerBot struct{
	config *Config
	logger *logrus.Logger
}

func New(config *Config) *ServerBot {
	return &ServerBot{
		config: config,
		logger: logrus.New(),
	}
}

func (s *ServerBot) Start() error {
	if err:=s.configureLogger();err!=nil{
		return err
	}

	s.logger.Info("server started")

	return nil
}
func (s *ServerBot) configureLogger() error {
	level,err:=logrus.ParseLevel(s.config.LogLevel)
	if err!=nil{
		return err
	}
	s.logger.SetLevel(level)
	return nil
}
func (s *ServerBot) name()  {
	mux:=http.NewServeMux()


}

*/
