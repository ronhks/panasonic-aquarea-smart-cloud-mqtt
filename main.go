package main

import (
	"time"

	"github.com/ronhks/panasonic-aquarea-smart-cloud-mqtt/src/config"
	"github.com/ronhks/panasonic-aquarea-smart-cloud-mqtt/src/data"
	httputils "github.com/ronhks/panasonic-aquarea-smart-cloud-mqtt/src/http"
	"github.com/ronhks/panasonic-aquarea-smart-cloud-mqtt/src/mqtt"
	"github.com/ronhks/panasonic-aquarea-smart-cloud-mqtt/src/panasonic"
	log "github.com/sirupsen/logrus"
)

func main() {
	initializeTheEnvironment()

	err := loginAndGetContract()
	if err != nil {
		log.Error(err)
		log.Error("Error while Login and get Contract")
		return
	}

	startQueryStatusData()
}

func startQueryStatusData() {
	maxTries := 3
	for maxTries > 0 {
		success := getStatusData()
		if success {
			maxTries = 3
			time.Sleep(config.GetConfig().RefreshInterval)
		} else {
			err := loginAndGetContract()
			if err != nil {
				log.Error(err)
				maxTries--
			}
			startQueryStatusData()
		}
	}
}

func initializeTheEnvironment() {
	httputils.InitHttpClient()
	mqtt.InitMqttConnection()
	setLogSetup()

}

func setLogSetup() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.SetReportCaller(true)
}

func loginAndGetContract() error {

	panasonic.Login()

	err := panasonic.GetContractAndSetGwidAndDeviceIdInCookie()
	if err != nil {
		log.Fatal(err)
		log.Error("Error while get Contract")

		return err
	}

	return nil
}

func getStatusData() bool {

	statusData, err := data.GetDeviceData()
	if err != nil || len(statusData.Status) == 0 {
		log.Error(err)
		log.Error("Error while get DeviceData")
		return false
	}
	mqtt.PublishStatus(statusData)
	log.Trace(statusData)

	return true
}
