package water

import (
	"data"
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	httputils "http"
	"net/http"
	"session"
)

func SetWaterTemp(newTemp int) error {
	deviceGuid, deviceDataURLWithDeviceID := session.GetSessionInitData()

	var tankStatus data.TankStatus
	tankStatus.HeatSet = newTemp
	var status data.Status
	status.DeviceGuid = deviceGuid
	status.TankStatus = []data.TankStatus{tankStatus}
	var statusDataInput data.StatusData
	statusDataInput.Status = []data.Status{status}

	byteArray, err := getJsonByteArray(statusDataInput)
	if err != nil {
		log.Error(err)
		return err
	}

	response, err := httputils.PostREQWithJsonBody(deviceDataURLWithDeviceID, byteArray)

	if err != nil {
		log.Error(err)
		return err
	}
	if response.StatusCode != http.StatusOK {
		log.Error("HTTP call result code is:", response.StatusCode)
		return errors.New("NOK HTTP response code")
	}

	log.Info("Water temp set to ", newTemp, " C")

	return nil
}

func SetOperationOn() error {
	deviceGuid, deviceDataURLWithDeviceID := session.GetSessionInitData()

	setStatusDataInput := setTankStatus(data.ON, data.ON, deviceGuid)

	byteArray, err := getJsonByteArrayFromSetStatus(setStatusDataInput)
	if err != nil {
		log.Error(err)
		return err
	}

	response, err := httputils.PostREQWithJsonBody(deviceDataURLWithDeviceID, byteArray)

	if err != nil {
		log.Error(err)
		return err
	}
	if response.StatusCode != http.StatusOK {
		log.Error("HTTP call result code is:", response.StatusCode)
		return errors.New("NOK HTTP response code")
	}

	log.Info("Water operation set ON")

	return nil
}

func SetOperationOff() error {

	deviceGuid, deviceDataURLWithDeviceID := session.GetSessionInitData()

	statusDataInput := setTankStatus(data.OFF, data.ON, deviceGuid)

	byteArray, err := json.Marshal(&statusDataInput)

	if err != nil {
		log.Errorf("Fail to create convert JSON, %v", err.Error())
		return err
	}

	response, err := httputils.PostREQWithJsonBody(deviceDataURLWithDeviceID, byteArray)

	if err != nil {
		log.Error(err)
		return err
	}
	if response.StatusCode != http.StatusOK {
		log.Error("HTTP call result code is:", response.StatusCode)
		return errors.New("NOK HTTP response code")
	}

	log.Info("Water operation set OFF")

	return nil
}

func setTankStatus(newStatus int, deviceNewStatus int, deviceGuid string) data.SetStatusData {
	var tankStatus data.SetTankStatus
	tankStatus.OperationStatus = &newStatus
	var status data.SetStatus
	status.TankStatus = []data.SetTankStatus{tankStatus}
	status.DeviceGuid = deviceGuid
	status.OperationStatus = &deviceNewStatus
	var statusDataInput data.SetStatusData
	statusDataInput.Status = []data.SetStatus{status}

	return statusDataInput
}

func getJsonByteArray(statusDataInput data.StatusData) (jsonByteArray []byte, error error) {
	byteArray, err := json.Marshal(&statusDataInput)

	if err != nil {
		log.Errorf("Fail to create convert JSON, %v", err.Error())
		return nil, err
	}
	return byteArray, err
}

func getJsonByteArrayFromSetStatus(setStatusDataInput data.SetStatusData) ([]byte, error) {
	byteArray, err := json.Marshal(&setStatusDataInput)

	if err != nil {
		log.Errorf("Fail to create convert JSON, %v", err.Error())
		return nil, err
	}
	return byteArray, err
}
