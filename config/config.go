package config

type Config struct {
	Bot struct {
		Token string
	}

	Redis struct {
		Port string
		Host string
	}

	Http struct {
		Port string
	}
	//WebHookPort string
}

func (c *Config) Parse() error {
	return nil
}

func NewConfig() *Config {
	return &Config{}
}
