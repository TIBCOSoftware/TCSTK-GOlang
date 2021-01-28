/*
* BSD 3-Clause License
* Copyright © 2020. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
 */
package processmanagement

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/TIBCOSoftware/TCSTK-GOlang/tibcocloud/idm"
	"github.com/TIBCOSoftware/TCSTK-common-mods/utils"
	"github.com/project-flogo/core/data/coerce"
	"github.com/stretchr/testify/assert"
)

func TestPostProcessesAction(t *testing.T) {

	props, err := utils.ReadPropertiesFile("../../credentials.properties")
	if err != nil {
		panic("Error while reading properties file")
	}

	cookie, _ := idm.LoginOauth("cookie", props["region"], props["tenant"], props["username"], props["password"], props["clientid"], props["version"])

	id := "\"id\": \"16353\""
	applicationID := "\"applicationId\": \"3495\""
	sandboxID := "\"sandboxID\": \"3100\""
	caseReference := "\"caseReference\": \"427641\""
	data := "\"data\": \"{\\\"MCTestOneTimeClick\\\":{\\\"field1\\\":\\\"desde tester2\\\",\\\"field2\\\":\\\"2\\\",\\\"field3\\\":\\\"3\\\"}}\""
	body := "{" + id + "," + applicationID + "," + sandboxID + "," + caseReference + "," + data + "}"
	var action ProcessDetails
	b, _ := coerce.ToBytes(body)
	json.Unmarshal(b, &action)

	_, workitem, _, _ := PostProcesses(props["region"], cookie, action)
	fmt.Printf("Response: %v\n", workitem)

	assert.Equal(t, true, true)
}

func TestPostProcessesStart(t *testing.T) {

	props, err := utils.ReadPropertiesFile("../../credentials.properties")
	if err != nil {
		panic("Error while reading properties file")
	}

	cookie, _ := idm.LoginOauth("cookie", props["region"], props["tenant"], props["username"], props["password"], props["clientid"], props["version"])

	id := "\"id\": \"16352\""
	applicationID := "\"applicationId\": \"3495\""
	sandboxID := "\"sandboxID\": \"3100\""
	data := "\"data\": \"{\\\"MCTestOneTimeClick\\\":{\\\"state\\\":\\\"Created\\\",\\\"field1\\\":\\\"desde tester\\\",\\\"field2\\\":\\\"2\\\",\\\"field3\\\":\\\"3\\\"}}\""
	body := "{" + id + "," + applicationID + "," + sandboxID + "," + data + "}"
	var action ProcessDetails
	b, _ := coerce.ToBytes(body)
	json.Unmarshal(b, &action)

	_, processDetails, _, _ := PostProcesses(props["region"], cookie, action)
	fmt.Printf("Response: %v\n", processDetails)

	assert.Equal(t, true, true)
}
