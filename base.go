package gomonnify

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/jcobhams/gomonnify/testhelpers"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

//Base or shared methods

func newBase(config Config) *base {

	if tm, ok := os.LookupEnv("GOMONNIFY_TESTMODE"); ok && strings.ToUpper(tm) == "ON" {
		config.Environment = EnvTest
	}

	b := &base{
		Config:     &config,
		HTTPClient: &http.Client{Timeout: config.RequestTimeout},
	}

	switch config.Environment {
	case EnvSandbox:
		b.APIBaseUrl = APIBaseUrlSandbox
	case EnvLive:
		b.APIBaseUrl = APIBaseUrlLive
	case EnvTest:
		if testUrl, ok := os.LookupEnv("GOMONNIFY_TESTURL"); ok {
			b.APIBaseUrl = testUrl
			b.Config.SecretKey = testhelpers.SecretKey
		} else {
			panic("gomonnify is running in test mode but no test url provided. Please Set GOMONNIFY_TESTURL env var")
		}
	}

	return b
}

func (b *base) Login() (LoginResponse, error) {
	result := LoginResponse{}
	url := fmt.Sprintf("%v/v1/auth/login", b.APIBaseUrl)

	rawResponse, statusCode, err := b.postRequest(url, requestAuthTypeBasic, nil)
	if err != nil {
		return result, err
	}

	err = b.unmarshallJson(strings.NewReader(rawResponse), &result)
	if err != nil {
		return result, err
	}

	AuthToken = result.ResponseBody.AccessToken
	t := time.Second * time.Duration(result.ResponseBody.ExpiresIn)
	ExpiresIn = time.Now().UTC().Add(t)

	if statusCode != http.StatusOK {
		return result, failedRequestMessage(statusCode, result.ResponseCode, result.ResponseMessage)
	}
	return result, nil
}

func (b *base) unmarshallJson(body io.Reader, responseBody interface{}) error {
	rawResponseBody, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(rawResponseBody, &responseBody)
	if err != nil {
		return err
	}

	return nil
}

func (b *base) setBasicAuth(req *http.Request) {
	s := []byte(fmt.Sprintf("%v:%v", b.Config.APIKey, b.Config.SecretKey))
	req.Header.Set("Authorization", fmt.Sprintf("Basic %v", base64.StdEncoding.EncodeToString(s)))
}

func (b *base) setBearerAuth(req *http.Request) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", AuthToken))
}

func (b *base) isTokenExpired() bool {
	if time.Now().After(ExpiresIn) {
		return true
	}
	return false
}

func (b *base) postRequest(url string, authType requestAuthType, data interface{}) (string, int, error) {
	var payload io.Reader
	payload = nil

	if data != nil {
		p, err := json.Marshal(data)
		if err != nil {
			return "", 0, err
		}
		payload = bytes.NewReader(p)
	}

	return b.request("POST", url, authType, payload)
}

func (b *base) getRequest(url string, authType requestAuthType) (string, int, error) {
	return b.request("GET", url, authType, nil)
}

func (b *base) deleteRequest(url string, authType requestAuthType) (string, int, error) {
	return b.request("DELETE", url, authType, nil)
}

func (b *base) request(method, url string, authType requestAuthType, data io.Reader) (string, int, error) {
	req, err := http.NewRequest(method, url, data)
	if err != nil {
		return "", 0, err
	}

	switch authType {
	case requestAuthTypeBasic:
		b.setBasicAuth(req)
	case requestAuthTypeBearer:
		if b.isTokenExpired() {
			_, err := b.Login()
			if err != nil {
				return "", 0, err
			}
		}
		b.setBearerAuth(req)
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := b.HTTPClient.Do(req)
	if err != nil {
		return "", 0, err
	}

	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", resp.StatusCode, err
	}

	return string(body), resp.StatusCode, nil
}
