package panasonic

import (
	auth "github.com/ronhks/panasonic-aquarea-smart-cloud-mqtt/src/auth"
	conf "github.com/ronhks/panasonic-aquarea-smart-cloud-mqtt/src/config"
	httputils "github.com/ronhks/panasonic-aquarea-smart-cloud-mqtt/src/http"
	log "github.com/sirupsen/logrus"
)

func Login() {
	err := auth.Login()
	if err != nil {
		log.Error(err)
		return
	}
}

func Logout() {
	err := auth.Logout()
	if err != nil {
		log.Error(err)
		return
	}
}

func GetContractAndSetGwidAndDeviceIdInCookie() error {

	contractURL := conf.GetConfig().AquareaSmartCloudURL + "/remote/contract"

	_, err := httputils.PostREQ(contractURL)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
