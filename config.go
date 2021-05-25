package main

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
)

func ReadConfig() Config {
	var configfileName = "config"
	_, err := os.Stat(configfileName)
	if err != nil {
		log.Fatal("Config file is missing: ", configfileName)
	}

	var config Config
	if _, err := toml.DecodeFile(configfileName, &config); err != nil {
		log.Fatal(err)
	}
	return config
}

type Config struct {
	AquareaServiceCloudURL      string
	AquareaSmartCloudURL        string
	AquareaSmartCloudLogin      string
	AquareaSmartCloudPassword   string
	AquareaTimeout              int
	MqttServer                  string
	MqttPort                    string
	MqttLogin                   string
	MqttPass                    string
	MqttTopicRoot               string
	MqttClientID                string
	MqttKeepalive               int
	PoolInterval                int
	LogSecOffset                int64
}