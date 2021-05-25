package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
)

type LoginStruct struct {
	AgreementStatus struct {
		Contract      bool `json:"contract"`
		CookiePolicy  bool `json:"cookiePolicy"`
		PrivacyPolicy bool `json:"privacyPolicy"`
	} `json:"agreementStatus"`
	ErrorCode int `json:"errorCode"`
}

func GetLogin() error {

	var loginResponseStruct LoginStruct
	loginURL := config.AquareaSmartCloudURL + "/remote/v1/api/auth/login"

	uv := url.Values{
		"var.loginId":         {config.AquareaSmartCloudLogin},
		"var.password":        {config.AquareaSmartCloudPassword},
		"var.inputOmit":       {"false"},
	}

	response, err := PostREQWithParam(loginURL, client, uv)
	if err != nil {
		return err

	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {

		fmt.Println(err)
		return err

	}
	fmt.Println(string(body))
	defer response.Body.Close()

	err = json.Unmarshal(body, &loginResponseStruct)
	fmt.Println(err, "Error while parsing Login Response JSON", string(body))

	if isPanasonicResponseHasError(loginResponseStruct) {
		err = errors.New("Internal Panasonic Error. ErrorCode: " + fmt.Sprint(loginResponseStruct.ErrorCode))
	}

	if err != nil {
		fmt.Println(err)
		return err

	}
	return nil
}

func isPanasonicResponseHasError(loginStruct LoginStruct) bool {
	return loginStruct.ErrorCode != 0
}