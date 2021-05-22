package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	config *viper.Viper
}

func NewConfig() *Config {
	c := new(Config)
	c.config = readConfig()
	return c
}

func (c *Config) Get() *viper.Viper {
	if c.config == nil {
		log.Fatal("config not initialized")
	}
	return c.config
}

func readConfig() *viper.Viper {
	log.Info("reading environment variables")
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath(".")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("error reading config file or env variable '%s'", err.Error())
	}

	return v
}
