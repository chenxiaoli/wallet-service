package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	AUTH_URL = "http://www.baiud.com/"
)

type AuthApiService struct {
	http.Client
}

type LoginVo struct {
	Token string `json:"token"`
}

func (c *AuthApiService) Login(name string, password string) (string, error) {
	loginPath := "http://localhost:18000/login"
	var values url.Values
	var resp LoginVo
	response, err := c.PostForm(loginPath, values)
	fmt.Println("fxxxxxxxxxxxxxxx")
	if err != nil {
		return "", err
	} else if response.StatusCode == 200 {
		var body []byte
		body, err = ioutil.ReadAll(response.Body)
		if err != nil {
			return "", err
		}
		_ = json.Unmarshal(body, &resp)
	}
	return resp.Token, err
}

func (c *AuthApiService) Register(name string, password string) string {
	return ""
}
