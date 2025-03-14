package config

import (
	_ "embed"
	"gopkg.in/yaml.v3"
)

//go:embed config.yaml
var config []byte

type Config struct {
	Bot struct {
		Token string `yaml:"token"`
	} `yaml:"bot"`

	Redis struct {
		Port        string `yaml:"port"`
		Host        string `yaml:"host"`
		CachePrefix string `yaml:"cache_prefix"`
	} `yaml:"redis"`

	Http struct {
		Port string `yaml:"port"`
	} `yaml:"http"`

	DB struct {
		Port     string `yaml:"port"`
		Host     string `yaml:"host"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DbName   string `yaml:"dbname"`
		SslMode  string `yaml:"sslmode"`
	} `yaml:"db"`

	TG struct {
		ApiId   int32  `yaml:"api_id"`
		ApiHash string `yaml:"api_hash"`
	} `yaml:"tg"`

	Kafka struct {
		ProducerTopics struct {
			NewMessage string `yaml:"new_message"`
		} `yaml:"producer_topics"`
	} `yaml:"kafka"`
	//WebHookPort string
}

func (c *Config) Parse() error {
	return yaml.Unmarshal(config, c)
}

func NewConfig() *Config {
	return &Config{}
}
