/*
* BSD 3-Clause License
* Copyright © 2020. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
 */
package authorization

import (
	"fmt"
	"testing"

	"github.com/TIBCOSoftware/TCSTK-GOlang/tibcocloud/idm"
	"github.com/TIBCOSoftware/TCSTK-common-mods/utils/properties"
	"github.com/stretchr/testify/assert"
)

func TestClaims(t *testing.T) {

	props, err := properties.ReadPropertiesFile("../../credentials.properties")
	if err != nil {
		panic("Error while reading properties file")
	}

	cookie, _ := idm.LoginOauth("cookie", props["region"], props["tenant"], props["username"], props["password"], props["clientid"], props["version"])
	statusCode, claims, _, _ := GetClaims(props["region"], cookie)
	fmt.Printf("Response: %v %v\n", statusCode, claims)
	assert.Equal(t, statusCode, 200)
}

func TestClaimsAccessToken(t *testing.T) {

	token := ""

	_, claims, _, _ := GetClaims("US", token)
	fmt.Printf("Response: %v\n", claims)

	assert.Equal(t, true, true)
}
