package httputils

import (
	"config"
	"crypto/tls"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

var client http.Client
var	cookieJar cookiejar.Jar

func InitHttpClient() {
	cookieJar, _ := cookiejar.New(nil)
	client = http.Client{
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
		Jar:       cookieJar,
		Timeout: config.GetConfig().HttpTimeout,
	}
}

func PostREQ(url string) (*http.Response, error) {
	return PostREQWithReferer(url,config.GetConfig().AquareaSmartCloudURL,nil)
}
func PostREQWithParam(url string, uv url.Values) (*http.Response, error) {
	return PostREQWithReferer(url,config.GetConfig().AquareaSmartCloudURL,uv)

}
func PostREQWithReferer(url string, referer string, uv url.Values) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(uv.Encode()))
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
	url,_ := url.Parse(config.GetConfig().AquareaSmartCloudURL)

	selectedDeviceId := ""
	for _,cookie := range client.Jar.Cookies(url) {
		if "selectedDeviceId" == cookie.Name {
			selectedDeviceId = cookie.Value
			break
		}
	}
	return selectedDeviceId
}