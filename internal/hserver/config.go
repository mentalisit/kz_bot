package hserver

type Config struct {
	BindAdrr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
}

func NewConfig() *Config {
	return &Config{
		BindAdrr: ":8080",
		LogLevel: "debug",
	}

}
