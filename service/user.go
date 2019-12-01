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

func (c *AuthApiService) Login(username string, password string) (string, error) {
	loginPath := "http://localhost:18000/api/v1/login"
	var values url.Values
	var resp LoginVo
	var body []byte
	values = url.Values{}
	values.Set("username", username)
	values.Set("password", password)
	fmt.Println(username)
	fmt.Println(password)
	response, err := c.PostForm(loginPath, values)
	fmt.Println(response)
	if err != nil {
		return "", err
	} else if response.StatusCode == 200 {
		body, err = ioutil.ReadAll(response.Body)
		if err != nil {
			return "", err
		}
		_ = json.Unmarshal(body, &resp)
		return resp.Token, err
	} else if response.StatusCode == 404 {
		return "", err
	} else if response.StatusCode == 500 {
		body, err = ioutil.ReadAll(response.Body)
		fmt.Println(string(body))
	} else {
		var errMsg map[string]string
		body, err = ioutil.ReadAll(response.Body)
		_ = json.Unmarshal(body, &errMsg)
		fmt.Println(string(body))
		return errMsg["detail"], err
	}
	return "", err
}

func (c *AuthApiService) Register(name string, password string) string {
	return ""
}
