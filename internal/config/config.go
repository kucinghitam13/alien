package config

import (
	"log"
	"os"

	gcfg "gopkg.in/gcfg.v1"
)

type AppConfig struct {
	Config *Config
}

func InitConfig() *AppConfig {
	config := &Config{}
	environment := os.Getenv("ENV")
	if environment == "" {
		environment = "development"
	}
	configFile := "files/etc/alien/" + "alien" + "." + environment + ".ini"
	if err := gcfg.ReadFileInto(config, configFile); err != nil {
		configFile = "../../" + configFile
		if err = gcfg.ReadFileInto(config, configFile); err != nil {
			log.Fatalf("Error reading config .ini file from %s: %s", configFile, err)
			panic("Error reading config .ini")
		}
	}

	return &AppConfig{
		Config: config,
	}
}
