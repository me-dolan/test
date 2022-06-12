package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Port      string `yaml:"port"`
	MongoUrl  string `yaml:"mongoUrl"`
	SecretKey string `yaml:"secretKey"`
}

func NewConfig(path string) (*Config, error) {
	c := new(Config)
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
