package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ronhks/panasonic-aquarea-smart-cloud-mqtt/src/config"
	"github.com/ronhks/panasonic-aquarea-smart-cloud-mqtt/src/http"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
)

type SetTemp struct {
	NewTemp int `json:"newTemp"`
}

const (
	ON                  int = 1
	OFF                 int = 0
	ZoneId1             int = 1
	ZoneId2             int = 2
	ModeHeatAndHotWater int = -1
	ModeOnlyHotWater    int = 0
)

type SetOpertaion struct {
	NewMode int `json:"newMode"`
}

type SetStatusData struct {
	Status []SetStatus `json:"status,omitempty"`
}

type SetStatus struct {
	OperationStatus *int            `json:"operationStatus,omitempty"`
	OperationMode   *int            `json:"operationMode,omitempty"`
	DeviceGuid      string          `json:"deviceGuid,omitempty"`
	ZoneStatus      []SetZoneStatus `json:"zoneStatus,omitempty"`
	TankStatus      []SetTankStatus `json:"tankStatus,omitempty"`
}

type SetZoneStatus struct {
	OperationStatus *int `json:"operationStatus,omitempty"`
	ZoneId          *int `json:"zoneId,omitempty"`
}

type SetTankStatus struct {
	OperationStatus *int `json:"operationStatus,omitempty"`
}
type Status struct {
	DeviceStatus    int `json:"deiceStatus,omitempty"`
	OperationStatus int `json:"operationStatus,omitempty"`
	SpecialStatus   []struct {
		SpecialMode     int `json:"specialMode"`
		OperationStatus int `json:"operationStatus"`
	} `json:"specialStatus,omitempty"`
	ZoneStatus    []ZoneStatus `json:"zoneStatus,omitempty"`
	OutdoorNow    int          `json:"outdoorNow,omitempty"`
	OperationMode int          `json:"operationMode,omitempty"`
	HolidayTimer  int          `json:"holidayTimer,omitempty"`
	DeviceGuid    string       `json:"deviceGuid,omitempty"`
	Bivalent      int          `json:"bivalent,omitempty"`
	TankStatus    []TankStatus `json:"tankStatus,omitempty"`
}

type ZoneStatus struct {
	OperationStatus int `json:"operationStatus,omitempty"`
	EcoHeat         int `json:"ecoHeat,omitempty"`
	CoolMin         int `json:"coolMin,omitempty"`
	HeatMin         int `json:"heatMin,omitempty"`
	ComfortCool     int `json:"comfortCool,omitempty"`
	TemparatureNow  int `json:"temparatureNow,omitempty"`
	CoolSet         int `json:"coolSet,omitempty"`
	ZoneId          int `json:"zoneId,omitempty"`
	ComfortHeat     int `json:"comfortHeat,omitempty"`
	HeatMax         int `json:"heatMax,omitempty"`
	CoolMax         int `json:"coolMax,omitempty"`
	EcoCool         int `json:"ecoCool,omitempty"`
	HeatSet         int `json:"heatSet,omitempty"`
}

type TankStatus struct {
	OperationStatus int `json:"operationStatus,omitempty"`
	TemparatureNow  int `json:"temparatureNow,omitempty"`
	HeatMax         int `json:"heatMax,omitempty"`
	HeatMin         int `json:"heatMin,omitempty"`
	HeatSet         int `json:"heatSet,omitempty"`
}

type StatusData struct {
	Status    []Status `json:"status,omitempty"`
	ErrorCode int      `json:"errorCode,omitempty"`
}

func GetDeviceData() (StatusData, error) {

	deviceDataURL := config.GetConfig().AquareaSmartCloudURL + "/remote/v1/api/devices/"

	var statusData StatusData

	deviceDataURLWithDeviceID := deviceDataURL + httputils.GetDeviceIdFromCookie()

	const referer = "https://aquarea-smart.panasonic.com/remote/a2wEnergyConsumption?keepState=true"
	response, err := httputils.GetREQ(deviceDataURLWithDeviceID, referer)
	if response.StatusCode != http.StatusOK {
		log.Error("HTTP call result code is:", response.StatusCode)
		errors.New("NOK HTTP response code")
		return statusData, err
	}

	if err != nil {
		return statusData, err
	}
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Error(err)
		return statusData, err
	}
	log.Trace(string(body))

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	err = json.Unmarshal(body, &statusData)
	if err != nil {
		log.Error(err, "Error while paring DataStruct", string(body))
		return statusData, err
	}

	if statusData.ErrorCode != 0 {
		err = errors.New(fmt.Sprint(statusData.ErrorCode))
		log.Error(err)
	}
	return statusData, err
}
