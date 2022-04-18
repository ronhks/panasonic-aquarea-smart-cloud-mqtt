package session

import (
	"github.com/ronhks/panasonic-aquarea-smart-cloud-mqtt/src/config"
	httputils "github.com/ronhks/panasonic-aquarea-smart-cloud-mqtt/src/http"
)

func GetSessionInitData() (deviceGuid string, deviceDataURLWithDeviceID string) {
	deviceGuid = httputils.GetDeviceIdFromCookie()
	deviceDataURLWithDeviceID = config.GetDeviceDataURL() + deviceGuid

	return deviceGuid, deviceDataURLWithDeviceID
}
