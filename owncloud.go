package ownCloud

import (
	"encoding/json"
	//"fmt"
	"io/ioutil"
	"log"
	"strings"
	"net/http"
)

type ownCloud struct {
	serverUrl string
	user string
	password string
}

func (sdk *ownCloud) Connect(serverUrl string) bool {
	sdk.serverUrl = serverUrl
	return sdk.IsConnected() && sdk.Version() != ""
}

func (sdk *ownCloud) Disconnect() bool {
	sdk.serverUrl = ""
	return !sdk.IsConnected()
}

func (sdk *ownCloud) IsConnected() bool {
	return sdk.serverUrl != ""
}

func (sdk *ownCloud) Login(user string, password string) bool {
	if !sdk.IsConnected() {
		return false
	}
	
	sdk.user = user
	sdk.password = password
	return sdk.IsLoggedIn()
}

func (sdk *ownCloud) Logout() bool {
	if !sdk.IsConnected() {
		return false
	}

	sdk.user = ""
	sdk.password = ""
	return !sdk.IsLoggedIn()
}

func (sdk *ownCloud) IsLoggedIn() bool {
	return sdk.IsConnected() && sdk.user != "" && sdk.password != ""
}

func (sdk *ownCloud) Capabilities() interface{} {
	//capabilities := make(map[string]string)
	var capabilities interface{}
	if !sdk.IsLoggedIn() {
		return capabilities
	}

	client := &http.Client{}

	/* Authenticate */
	req, err := http.NewRequest(http.MethodGet, sdk.serverUrl + "ocs/v1.php/cloud/capabilities?format=json", nil)
	req.SetBasicAuth(sdk.user, sdk.password)

	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return capabilities
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
			return capabilities
		}

		json.Unmarshal(contents, &capabilities)
		return capabilities
	}
}

func (sdk *ownCloud) getBasicAuthServerUrl() string {
	if !sdk.IsLoggedIn() {
		return ""
	}
	return strings.Replace(sdk.serverUrl, "://", "://" + sdk.user + ":" + sdk.password + "@", 1)
}

func (sdk *ownCloud) Status() map[string]string {
	status := make(map[string]string)
	if !sdk.IsConnected() {
		return status
	}
	
	client := &http.Client{}
	response, err := client.Get(sdk.serverUrl + "status.php")
	if err != nil {
		log.Fatal(err)
		return status
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
			return status
		}

		json.Unmarshal(contents, &status)
		return status
	}
}

func (sdk *ownCloud) Version() string {
	status := sdk.Status()
	return status["version"]
}
