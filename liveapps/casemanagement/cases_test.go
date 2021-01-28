/*
* BSD 3-Clause License
* Copyright © 2020. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
 */
package casemanagement

import (
	"net/url"
	"testing"

	"github.com/TIBCOSoftware/TCSTK-GOlang/tibcocloud/idm"
	"github.com/TIBCOSoftware/TCSTK-common-mods/utils/properties"
	"github.com/stretchr/testify/assert"
)

func TestGetCases(t *testing.T) {
	props, err := properties.ReadPropertiesFile("../../credentials.properties")
	if err != nil {
		panic("Error while reading properties file")
	}

	cookie, _ := idm.LoginOauth("cookie", props["region"], props["tenant"], props["username"], props["password"], props["clientid"], props["version"])

	queryParams := url.Values{}
	queryParams.Set("$sandbox", "31")
	queryParams.Set("$filter", "applicationId eq 3406 and typeId eq 1 and purgeable eq TRUE and modificationTimestamp le 2020-03-17T13:18:28.423Z")
	// queryParams.Set("$select", "")
	// queryParams.Set("$skip", "0")
	// queryParams.Set("$top", "")
	// queryParams.Set("$search", "")
	// queryParams.Set("$count", "");
	// queryParams.Set("$user", "")

	_, cases, _, _ := GetCases(props["region"], cookie, queryParams)

	assert.NotEmpty(t, cases)
}

func TestDeleteCases(t *testing.T) {
	props, err := properties.ReadPropertiesFile("../../credentials.properties")
	if err != nil {
		panic("Error while reading properties file")
	}

	cookie, _ := idm.LoginOauth("cookie", props["region"], props["tenant"], props["username"], props["password"], props["clientid"], props["version"])

	queryParams := url.Values{}
	queryParams.Set("$sandbox", "31")
	queryParams.Set("$filter", "applicationId eq 3406 and typeId eq 1 and purgeable eq TRUE and modificationTimestamp le 2020-03-17T13:18:28.423Z")

	_, casesDeleted, _, _ := DeleteCases(props["region"], cookie, queryParams)

	assert.NotEmpty(t, casesDeleted)
}
