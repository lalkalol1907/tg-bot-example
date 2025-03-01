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
		Port string `yaml:"port"`
		Host string `yaml:"host"`
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
			MessageToBot    string `yaml:"message_to_bot"`
			MessageToSender string `yaml:"message_to_sender"`
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
