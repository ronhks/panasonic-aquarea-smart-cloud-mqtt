package main

import (
	"encoding/json"
	"fmt"
	"time"

	"config"
	"data"
	log "github.com/sirupsen/logrus"
	httputils "http"
	"mqtt"
	"panasonic"
)



var LastChecksum [16]byte

func main() {
	initializeTheEnvironment()

	loginAndGetContract()

	startQueryStatusData()
}

func startQueryStatusData() {
	for {
		getStatusData()
		time.Sleep(config.GetConfig().RefreshInterval)
	}
}

func initializeTheEnvironment() {
	httputils.InitHttpClient()
	mqtt.InitMqttConnection()
}

func loginAndGetContract() error {

	panasonic.Login()

	err := panasonic.GetContractAndSetGwidAndDeviceIdInCookie()
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
		fmt.Println("Error while get Contract")

		return err
	}

	return nil
}

func getStatusData() bool {

	statusData,err := data.GetDeviceData()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error while get DeviceData")

		return false
	}

	statusDataJson, err := json.Marshal(statusData)
	mqtt.PublishLog("/status", string(statusDataJson))
	mqtt.PublishLog("/outdoor/temp", fmt.Sprintf("%d", statusData.Status[0].OutdoorNow))
	fmt.Println(statusData)

	return true
}
