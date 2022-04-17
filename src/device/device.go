package device

import (
	"data"
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	httputils "http"
	"net/http"
	"session"
)

func SetOperationOn() error {
	err := changeStatus(data.ON)
	if err != nil {
		log.Error(err)
		return err
	}

	log.Info("Device Operation set ON")

	return nil
}
func SetOperationOff() error {
	err := changeStatus(data.OFF)
	if err != nil {
		log.Error(err)
		return err
	}

	log.Info("Device Operation set OFF")

	return nil
}

func changeStatus(newStatus int) error {
	deviceGuid, deviceDataURLWithDeviceID := session.GetSessionInitData()

	statusDataInput := setOperationStatus(newStatus, deviceGuid)

	byteArray, err := json.Marshal(&statusDataInput)

	if err != nil {
		log.Error("Fail to create convert JSON, %v", err.Error())
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
	return nil
}

func setOperationStatus(newStatus int, deviceGuid string) data.SetStatusData {
	var status data.SetStatus
	status.DeviceGuid = deviceGuid
	status.OperationStatus = &newStatus
	var statusDataInput data.SetStatusData
	statusDataInput.Status = []data.SetStatus{status}

	return statusDataInput
}
