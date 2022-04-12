package config

import (
	"github.com/ghodss/yaml"
	"io/ioutil"
	"os"
)

type Config struct {
	Mysql struct {
		IP       string `json:"ip" yaml:"ip"`
		Port     int    `json:"port" yaml:"port"`
		Username string `json:"username" yaml:"username"`
		Password string `json:"password" yaml:"password"`
		DBName   string `json:"DBName" yaml:"DBName"`
	} `json:"mysql" yaml:"mysql"`
}

func Init() (*Config, error) {
	file, err := os.Open("internal/config/config.yml")
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	cfg := &Config{}
	if err = yaml.Unmarshal(b, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
