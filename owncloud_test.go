/*
Copyright 2014 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package ownCloud

import (
	"testing"
	"fmt"
)

func TestConnect(t *testing.T) {
	sdk := ownCloud{}
	if !sdk.Connect("http://localhost:8081/ownCloud/master/core/") {
		t.Error("Can not connect to server")
	}
}

func TestLogin(t *testing.T) {
	sdk := ownCloud{}
	if !sdk.Connect("http://localhost:8081/ownCloud/master/core/") {
		t.Skip("Can not connect to server")
	}
	if !sdk.Login("admin", "admin") {
		t.Errorf("Can not login to server: %s as user: %s with password: %s", sdk.serverUrl, sdk.user, sdk.password)
	}
}

func TestCapabilities(t *testing.T) {
	sdk := ownCloud{}
	if !sdk.Connect("http://localhost:8081/ownCloud/master/core/") {
		t.Skip("Can not connect to server")
	}
	if !sdk.Login("admin", "admin") {
		t.Skip("Can not login to server: %s as user: %s with password: %s", sdk.serverUrl, sdk.user, sdk.password)
	}
	var capabilities = sdk.Capabilities()
	m := capabilities.(map[string]interface{})
	ocs := m["ocs"].(map[string]interface{})
	fmt.Println(ocs)
	meta := ocs["meta"].(map[string]interface{})
	fmt.Println(meta["status"])

	if len(meta) == 0 || len(meta) != 0 {
		t.Error("Can not get capabilities", meta["status"])
	}
}
