package httputils

import (
	"bytes"
	"crypto/tls"
	"github.com/ronhks/panasonic-aquarea-smart-cloud-mqtt/src/config"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

var client http.Client

func InitHttpClient() {
	cookieJar, _ := cookiejar.New(nil)
	client = http.Client{
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
		Jar:       cookieJar,
		Timeout:   config.GetConfig().HttpTimeout,
	}
}

func PostREQ(url string) (*http.Response, error) {
	return PostREQWithReferer(url, config.GetConfig().AquareaSmartCloudURL, nil, "")
}

func PostREQWithURLParam(url string, uv url.Values) (*http.Response, error) {
	return PostREQWithReferer(url, config.GetConfig().AquareaSmartCloudURL, strings.NewReader(uv.Encode()), "")
}

func PostREQWithJsonBody(url string, jsonPayload []byte) (*http.Response, error) {
	return PostREQWithReferer(url, config.GetConfig().AquareaSmartCloudURL, bytes.NewBuffer(jsonPayload), "application/json")
}

func PostREQWithReferer(url string, referer string, body io.Reader, contentType string) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, body)
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:74.0) Gecko/20100101 Firefox/74.0")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Encoding", "deflate, br")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	if len(contentType) == 0 {
		contentType = "application/x-www-form-urlencoded"
	}
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Referer", referer)

	resp, err := client.Do(req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
func GetREQ(url string, referer string) (*http.Response, error) {
	return GetREQWithParam(url, referer, nil)
}
func GetREQWithParam(url string, referer string, uv url.Values) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, strings.NewReader(uv.Encode()))
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:74.0) Gecko/20100101 Firefox/74.0")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Encoding", "deflate, br")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Referer", referer)

	resp, err := client.Do(req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func GetDeviceIdFromCookie() string {
	smartCloudUrl, _ := url.Parse(config.GetConfig().AquareaSmartCloudURL)

	selectedDeviceId := ""
	for _, cookie := range client.Jar.Cookies(smartCloudUrl) {
		if "selectedDeviceId" == cookie.Name {
			selectedDeviceId = cookie.Value
			break
		}
	}
	return selectedDeviceId
}

func HandleResponse(response *http.Response) {
	if response.StatusCode != 200 {
		log.Error("HTTP call result code is:", response.StatusCode)
	}
}
