package panasonic

import (
	conf "config"
	httputils "http"
	"login"
)

func Login() {
	login.GetLogin()
}

func GetContractAndSetGwidAndDeviceIdInCookie() error {

	contractURL := conf.GetConfig().AquareaSmartCloudURL + "/remote/contract"

	_, err := httputils.PostREQ(contractURL)
	if err != nil {
		return err
	}
	return nil
}