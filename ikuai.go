package ikuai

import (
	"crypto/md5"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/wuxs/ikuai/action"
)

type IKuai struct {
	client   *resty.Client
	debug    bool
	Url      string
	Username string
	Password string

	session string
}

func NewIKuai(url string, username string, password string, insecureSkipVerify, autoLogin bool) *IKuai {
	client := resty.New()

	if insecureSkipVerify {
		client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}

	i := &IKuai{
		client:   client,
		Url:      url,
		Username: username,
		Password: password,
	}

	if autoLogin {
		client.SetRetryCount(2)
		client.AddRetryCondition(func(response *resty.Response, err error) bool {
			body := response.Body()
			var result action.Result
			rErr := json.Unmarshal(body, &result)
			if rErr != nil {
				log.Printf("Unmarshal error: %v", rErr)
				return false
			}

			if result.Result == 10014 {
				log.Printf("session timeout: try to login")
				_, err := i.Login()
				if err != nil {
					return false
				}

				log.Printf("successfully login, re-try to meter states")

				return true
			}

			return false
		})
	}

	return i
}

type LoginRequest struct {
	Username string `json:"username"`
	Passwd   string `json:"passwd"`
}

func getMD5(password string) string {
	hash := md5.New()
	hash.Write([]byte(password))
	sum := hash.Sum(nil)

	return fmt.Sprintf("%x", sum)
}

func (i *IKuai) Login() (string, error) {
	var result action.Result

	response, err := i.client.R().
		SetBody(&LoginRequest{
			Username: i.Username,
			Passwd:   getMD5(i.Password),
		}).
		SetResult(&result).
		Post(i.Url + "/Action/login")

	if err != nil {
		return "", err
	}

	for _, cookie := range response.Cookies() {
		if cookie.Name == "sess_key" {
			i.session = cookie.Value
			return cookie.Value, nil
		}
	}

	log.Printf("login error: %s", response.Body())

	return "", errors.New(fmt.Sprintf("login error: %s, no cookies", result.ErrMsg))
}

func (i *IKuai) Run(session string, action *action.Action, result interface{}) (string, error) {
	url := i.Url + "/Action/call"

	response, err := i.client.R().
		SetHeader("Content-Type", "application/json").
		SetCookie(&http.Cookie{Name: "sess_key", Value: session}).
		SetBody(action).
		SetResult(result).
		Post(url)

	if err != nil {
		return "", err
	}

	if i.debug {
		log.Printf("POST %s, request: %v, response: %s", url, action, response.String())
	}

	return response.String(), nil
}

func (i *IKuai) Debug() {
	i.debug = true
}
