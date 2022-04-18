package panasonic

import (
	conf "github.com/ronhks/panasonic-aquarea-smart-cloud-mqtt/src/config"
	httputils "github.com/ronhks/panasonic-aquarea-smart-cloud-mqtt/src/http"
	"github.com/ronhks/panasonic-aquarea-smart-cloud-mqtt/src/login"
	log "github.com/sirupsen/logrus"
)

func Login() {
	err := login.GetLogin()
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
