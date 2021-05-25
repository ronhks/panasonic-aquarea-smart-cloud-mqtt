package main

import (
	"encoding/json"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)



var LastChecksum [16]byte

func main() {
	initializeTheEnvironment()

	loginAndGetContract()

	for {
		getStatusData()
		time.Sleep(PoolInterval)
	}
}

func initializeTheEnvironment() {
	readAndSetConfig()
	initHttpClient()
	initMqttConnection()
}

func loginAndGetContract() error {
	err := GetLogin()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error while logging in.")
		return err
	}

	err = GetContractAndSetGwidAndDeviceIdInCookie()
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
		fmt.Println("Error while get Contract")

		return err
	}

	return nil
}

func getStatusData() bool {

	statusData,err := GetDeviceData()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error while get DeviceData")

		return false
	}

	statusDataJson, err := json.Marshal(statusData)
	PublishLog("/status", string(statusDataJson))
	PublishLog("/outdoor/temp", fmt.Sprintf("%d", statusData.Status[0].OutdoorNow))
	fmt.Println(statusData)

	return true
}
