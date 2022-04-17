package panasonic

import (
	conf "config"
	log "github.com/sirupsen/logrus"
	httputils "http"
	"login"
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
