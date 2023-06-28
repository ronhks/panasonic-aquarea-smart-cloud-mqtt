package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ronhks/panasonic-aquarea-smart-cloud-mqtt/src/config"
	httputils "github.com/ronhks/panasonic-aquarea-smart-cloud-mqtt/src/http"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ResponseStruct struct {
	AgreementStatus struct {
		Contract      bool `json:"contract"`
		CookiePolicy  bool `json:"cookiePolicy"`
		PrivacyPolicy bool `json:"privacyPolicy"`
	} `json:"agreementStatus"`
	ErrorCode int `json:"errorCode"`
}

func Login() error {

	var loginResponseStruct ResponseStruct
	loginURL := config.GetConfig().AquareaSmartCloudURL + "/remote/v1/api/auth/login"

	uv := url.Values{
		"var.loginId":   {config.GetConfig().Username},
		"var.password":  {config.GetConfig().Password},
		"var.inputOmit": {"false"},
	}

	response, err := httputils.PostREQWithURLParam(loginURL, uv)
	if err != nil {
		log.Error(err)
		return err
	}
	if err != nil {
		log.Error(err)
		return err
	}
	if response.StatusCode != http.StatusOK {
		log.Error("HTTP call result code is:", response.StatusCode)
		return err
	}

	body, err := getBodyFromResponse(err, response)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &loginResponseStruct)
	if err != nil {
		log.Error(err, "Error while parsing Login Response JSON", string(body))
	}

	if isPanasonicResponseHasError(loginResponseStruct) {
		err = errors.New("Internal Panasonic Error. ErrorCode: " + fmt.Sprint(loginResponseStruct.ErrorCode))
	}

	log.Info("Login success")

	if err != nil {
		log.Error(err)
		return err

	}
	return nil
}

func getBodyFromResponse(err error, response *http.Response) ([]byte, error) {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Trace(string(body))
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Error(err)
			return
		}
	}(response.Body)
	return body, nil
}

func isPanasonicResponseHasError(loginStruct ResponseStruct) bool {
	return loginStruct.ErrorCode != 0
}

func Logout() error {
	logoutURL := config.GetConfig().AquareaSmartCloudURL + "/remote/v1/api/auth/logout"

	response, err := httputils.PostREQ(logoutURL)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		log.Error("HTTP call result code is:", response.StatusCode)
		return err
	}

	return nil
}
