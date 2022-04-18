package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"time"
)

var config *Config

func GetConfig() *Config {
	if config == nil {
		readAndSetConfig()
	}

	return config
}

func readAndSetConfig() {
	var configFilename = "etc/config"
	readConfig(configFilename)

	setTimeouts()
}

func readConfig(configFilename string) {
	_, err := os.Stat(configFilename)
	if err != nil {
		log.Fatal("Config file is missing: ", configFilename)
	}

	if _, err := toml.DecodeFile(configFilename, &config); err != nil {
		log.Fatal(err)
	}
}

func setTimeouts() {
	config.HttpTimeout = time.Second * config.HttpTimeout
	config.MqttKeepalive = time.Second * config.MqttKeepalive
	config.RefreshInterval = time.Second * config.RefreshInterval
}

type Config struct {
	AquareaSmartCloudURL string
	Username             string
	Password             string
	HttpTimeout          time.Duration
	MqttServer           string
	MqttPort             string
	MqttLogin            string
	MqttPass             string
	MqttTopicRoot        string
	MqttClientID         string
	MqttKeepalive        time.Duration
	RefreshInterval      time.Duration
	LogSecOffset         int64
}

func GetDeviceDataURL() string {
	return GetConfig().AquareaSmartCloudURL + "/remote/v1/api/devices/"
}
