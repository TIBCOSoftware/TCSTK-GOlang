/*
* BSD 3-Clause License
* Copyright © 2020. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
 */
package idm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/TIBCOSoftware/TCSTK-common-mods/utils/properties"
)

func TestCreateCookie(t *testing.T) {

	props, err := properties.ReadPropertiesFile("../../credentials.properties")
	if err != nil {
		panic("Error while reading properties file")
	}

	cookie, _ := LoginOauth("kkk", props["region"], props["tenant"], props["username"], props["password"], props["clientid"], props["version"])

	assert.NotEmpty(t, cookie.Domain)
}
