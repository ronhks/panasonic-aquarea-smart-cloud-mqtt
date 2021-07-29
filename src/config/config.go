package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"time"
)


var AquareaTimeout time.Duration
var MqttKeepalive time.Duration
var PoolInterval time.Duration

var config *Config

func GetConfig() *Config{
	if config == nil {
		readAndSetConfig()
	}

	return config
}

func readAndSetConfig(){
	var configfileName = "config"
	readConfig(configfileName)

	setTimeouts()
}

func readConfig(configfileName string) {
	_, err := os.Stat(configfileName)
	if err != nil {
		log.Fatal("Config file is missing: ", configfileName)
	}

	var config Config
	if _, err := toml.DecodeFile(configfileName, &config); err != nil {
		log.Fatal(err)
	}
}

func setTimeouts(){
	AquareaTimeout = time.Second * time.Duration(config.AquareaTimeout)
	MqttKeepalive = time.Second * time.Duration(config.MqttKeepalive)
	PoolInterval = time.Second * time.Duration(config.PoolInterval)
}

type Config struct {
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