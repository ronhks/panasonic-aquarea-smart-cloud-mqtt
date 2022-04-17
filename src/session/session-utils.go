package session

import (
	"config"
	httputils "http"
)

func GetSessionInitData() (deviceGuid string, deviceDataURLWithDeviceID string) {
	deviceGuid = httputils.GetDeviceIdFromCookie()
	deviceDataURLWithDeviceID = config.GetDeviceDataURL() + deviceGuid

	return deviceGuid, deviceDataURLWithDeviceID
}
