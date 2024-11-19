//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package test

import (
	"testing"

	"github.com/sbrueseke/cloudstack-go/v2/cloudstack"
)

func TestRegionService(t *testing.T) {
	service := "RegionService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testaddRegion := func(t *testing.T) {
		if _, ok := response["addRegion"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Region.NewAddRegionParams("endpoint", 0, "name")
		_, err := client.Region.AddRegion(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("AddRegion", testaddRegion)

	testlistRegions := func(t *testing.T) {
		if _, ok := response["listRegions"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Region.NewListRegionsParams()
		_, err := client.Region.ListRegions(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListRegions", testlistRegions)

	testremoveRegion := func(t *testing.T) {
		if _, ok := response["removeRegion"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Region.NewRemoveRegionParams(0)
		_, err := client.Region.RemoveRegion(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("RemoveRegion", testremoveRegion)

	testupdateRegion := func(t *testing.T) {
		if _, ok := response["updateRegion"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Region.NewUpdateRegionParams(0)
		_, err := client.Region.UpdateRegion(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("UpdateRegion", testupdateRegion)

}
