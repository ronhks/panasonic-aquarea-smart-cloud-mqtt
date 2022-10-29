package heat

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/ronhks/panasonic-aquarea-smart-cloud-mqtt/src/data"
	httputils "github.com/ronhks/panasonic-aquarea-smart-cloud-mqtt/src/http"
	"github.com/ronhks/panasonic-aquarea-smart-cloud-mqtt/src/session"
	log "github.com/sirupsen/logrus"
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

	var debugByteArray = string(byteArray)
	log.Info("SetOperationON  JSON bytearray: ", debugByteArray)

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

	var debugByteArray = string(byteArray)
	log.Info("SetOperationOff  JSON bytearray: ", debugByteArray)

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
	zoneStatus2 := setZoneStatus(data.ZoneId2, newStatus)

	status.ZoneStatus = []data.SetZoneStatus{zoneStatus1, zoneStatus2}
	status.OperationStatus = &deviceNewStatus

	status.OperationMode = &operationMode

	status.DeviceGuid = deviceGuid
	var statusDataInput data.SetStatusData
	statusDataInput.Status = []data.SetStatus{status}
	var result = statusDataInput

	return result
}

func setZoneStatus(zoneId int, newStatus int) data.SetZoneStatus {
	var zoneStatus data.SetZoneStatus
	zoneStatus.ZoneId = &zoneId
	zoneStatus.OperationStatus = &newStatus

	var result = zoneStatus

	var logMsg = "Set heating for zone: " + strconv.Itoa(zoneId) + " to status: " + strconv.Itoa(newStatus)
	log.Info(logMsg)

	return result
}

func SetHeatTemp(newTemp int) error {

	var zoneStatus1 = setZoneHeatTemp(1, newTemp)
	var zoneStatus2 = setZoneHeatTemp(2, newTemp)

	log.Info("Zone status for zone 2: ", zoneStatus1)
	log.Info("Zone status for zone 2: ", zoneStatus2)

	return nil
}

func setZoneHeatTemp(zoneId int, newTemp int) data.ZoneStatus {
	deviceGuid, deviceDataURLWithDeviceID := session.GetSessionInitData()

	//Offset for min an max temp
	var minTemp = newTemp - 3
	var maxTemp = newTemp + 3

	var zoneStatus data.ZoneStatus
	zoneStatus.ZoneId = zoneId
	zoneStatus.HeatSet = newTemp
	zoneStatus.HeatMin = minTemp
	zoneStatus.HeatMax = maxTemp

	var status data.Status
	status.DeviceGuid = deviceGuid
	status.ZoneStatus = []data.ZoneStatus{zoneStatus}
	var statusDataInput data.StatusData
	statusDataInput.Status = []data.Status{status}

	byteArray, err := getJsonByteArray(statusDataInput)

	var infoJsonDataRequest = string(byteArray)
	log.Info("Heat temperature JSON data: ", infoJsonDataRequest)

	var result = zoneStatus

	if err != nil {
		log.Error(err)
		return result
	}

	response, err := httputils.PostREQWithJsonBody(deviceDataURLWithDeviceID, byteArray)

	if err != nil {
		log.Error(err)
		return result
	}

	if response.StatusCode != http.StatusOK {
		log.Error("HTTP call result code is:", response.StatusCode)
		return result
	}

	log.Info("Heat temp set to ", newTemp, " C")
	return result

}

func getJsonByteArray(statusDataInput data.StatusData) (jsonByteArray []byte, error error) {
	byteArray, err := json.Marshal(&statusDataInput)

	if err != nil {
		log.Errorf("Fail to create convert JSON, %v", err.Error())
		return nil, err
	}
	return byteArray, err
}
