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

func TestGetTypes(t *testing.T) {
	props, err := properties.ReadPropertiesFile("../../credentials.properties")
	if err != nil {
		panic("Error while reading properties file")
	}

	cookie, _ := idm.LoginOauth("cookie", props["region"], props["tenant"], props["username"], props["password"], props["clientid"], props["version"])

	queryParams := url.Values{}
	queryParams.Set("$sandbox", "3100")
	queryParams.Set("$filter", "applicationId eq 3495")

	_, types, _, _ := GetTypes(props["region"], cookie, queryParams)

	assert.NotEmpty(t, types)
}
