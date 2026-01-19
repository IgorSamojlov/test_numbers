package main

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DBDSL string `envconfig:"dbdsl" required:"true"`
}

func NewConfig() (Config, error) {
	var c Config

	err := envconfig.Process("numbers", &c)

	if err != nil {
		return Config{}, err
	}

	return c, nil
}
