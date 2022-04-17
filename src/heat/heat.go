package heat

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
	deviceGuid, deviceDataURLWithDeviceID := session.GetSessionInitData()

	statusDataInput := setNewStatus(data.ON, data.ON, data.ModeHeatAndHotWater, deviceGuid)

	byteArray, err := json.Marshal(&statusDataInput)

	if err != nil {
		log.Error("Fail to create convert JSON, %v", err.Error())
		return errors.New("NOK HTTP response code")
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

	log.Info("Heat is set ON")

	return nil
}

func SetOperationOff() error {
	deviceGuid, deviceDataURLWithDeviceID := session.GetSessionInitData()

	statusDataInput := setNewStatus(data.OFF, data.ON, data.ModeOnlyHotWater, deviceGuid)

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

	log.Info("Heat is set OFF")

	return nil
}

func setNewStatus(newStatus int, deviceNewStatus int, operationMode int, deviceGuid string) data.SetStatusData {
	var status data.SetStatus

	zoneStatus1 := setZoneStatus(data.ZoneId1, newStatus)
	zoneStatus2 := setZoneStatus(data.ZoneId2, data.OFF)

	status.ZoneStatus = []data.SetZoneStatus{zoneStatus1, zoneStatus2}
	status.OperationStatus = &deviceNewStatus

	status.OperationMode = &operationMode

	status.DeviceGuid = deviceGuid
	var statusDataInput data.SetStatusData
	statusDataInput.Status = []data.SetStatus{status}
	return statusDataInput
}

func setZoneStatus(zoneId int, newStatus int) data.SetZoneStatus {
	var zoneStatus data.SetZoneStatus
	zoneStatus.ZoneId = &zoneId
	zoneStatus.OperationStatus = &newStatus
	return zoneStatus
}
