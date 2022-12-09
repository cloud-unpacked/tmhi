package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

func gatewayGet(endpoint string, authNeeded bool) (*http.Response, error) {

	return sendGatewayRequest("GET", endpoint, authNeeded, nil)
}

func gatewayPost(endpoint string, authNeeded bool, data []byte) (*http.Response, error) {

	return sendGatewayRequest("POST", endpoint, authNeeded, data)
}

func getAuthToken() (string, error) {

	resp, err := gatewayPost("auth/login", false, []byte(`{"username": "`+viper.GetString("username")+`", "password": "`+viper.GetString("password")+`"}`))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var respJSON loginResp

	err = json.Unmarshal(body, &respJSON)
	if err != nil {
		return "", err
	}

	return respJSON.Auth.Token, nil

}

func sendGatewayRequest(verb string, endpoint string, authNeeded bool, data []byte) (*http.Response, error) {

	var postData io.Reader
	if len(data) > 0 {
		postData = bytes.NewBuffer(data)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, verb, "http://192.168.12.1/TMI/v1/"+endpoint, postData)
	if err != nil {
		return nil, err
	}

	if authNeeded {

		token, err := getAuthToken()
		if err != nil {
			return nil, errors.New("This request requires authentication however an auth token was unable to be retrieved.")
		}

		req.Header.Add("Authorization", "Bearer "+token)
	}

	if postData != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil

}
