package config

import (
	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
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
	configEnvVariable := os.Getenv("PANASONIC_AQUAREA_SMART_CLOUD_MQTT_CONFIG")

	if len(configEnvVariable) == 0 {
		configEnvVariable = "etc/config"
	}

	var configFilename = configEnvVariable
	readConfig(configFilename)

	setTimeouts()
}

func readConfig(configFilename string) {
	_, err := os.Stat(configFilename)
	if err != nil {
		log.Error("Missing config file. You can define it in environment variable or use the default: etc/config")
		log.Fatal("Config file is missing: \"", configFilename, "\"")
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
