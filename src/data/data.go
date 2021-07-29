package data

import (
	"config"
	"encoding/json"
	"errors"
	"fmt"
	httputils "http"
	"io/ioutil"
)

type StatusData struct {
	Status[] struct {
		DeviceStatus     	 int `json:"deiceStatus"`
		OperationStatus 	 int `json:"operationStatus"`
		SpecialStatus[] struct {
			SpecialMode 	 	int `json:"specialMode"`
			OperationStatus 	int `json:"operationStatus"`
		} `json:"specialStatus"`
		ZoneStatus[] struct {
			OperationStatus 	int 	`json:"operationStatus"`
			EcoHeat 		 	int 	`json:"ecoHeat"`
			CoolMin 		 	int 	`json:"coolMin"`
			HeatMin 		 	int 	`json:"heatMin"`
			ComfortCool 		int 	`json:"comfortCool"`
			TemparatureNow 		int  	`json:"temparatureNow"`
			CoolSet 			int 	`json:"coolSet"`
			ZoneId 				int  	`json:"zoneId"`
			ComfortHeat 		int 	`json:"comfortHeat"`
			HeatMax 			int 	`json:"heatMax"`
			CoolMax 			int 	`json:"coolMax"`
			EcoCool 			int 	`json:"ecoCool"`
			HeatSet 			int 	`json:"heatSet"`
		} `json:"zoneStatus"`
		OutdoorNow     	 		int 	`json:"outdoorNow"`
		OperationMode 	 		int 	`json:"operationMode"`
		HolidayTimer 	 		int 	`json:"holidayTimer"`
		DeviceGuid 	 			string 	`json:"deviceGuid"`
		Bivalent 	 			int 	`json:"bivalent"`
	//	TankStatus struct {
		//	SpecialMode 	 	int `json:"specialMode"`
	//		OperationStatus 	bool `json:"operationStatus"`
	//	} `json:"specialStatus"`

	} `json:"status"`
	ErrorCode int `json:"errorCode"`
}

func GetDeviceData() (StatusData, error) {

	deviceDataURL := config.GetConfig().AquareaSmartCloudURL + "/remote/v1/api/devices/"

	var statusData StatusData

	deviceDataURLWithDeviceID := deviceDataURL + httputils.GetDeviceIdFromCookie()

	const referer = "https://aquarea-smart.panasonic.com/remote/a2wEnergyConsumption?keepState=true"
	resp, err := httputils.GetREQ(deviceDataURLWithDeviceID, referer)
	if err != nil {
		return statusData, err
	}
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return statusData,err
	}
	fmt.Println(string(body))
	defer resp.Body.Close()

	err = json.Unmarshal(body, &statusData)
	if err != nil {
		fmt.Println(err, "Error while paring DataStruct", string(body))
		return statusData,err
	}

	if statusData.ErrorCode != 0 {
		err = errors.New(fmt.Sprint(statusData.ErrorCode))
		fmt.Println(err)
	}
	return statusData,err
}
