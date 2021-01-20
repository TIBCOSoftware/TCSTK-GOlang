/*
* BSD 3-Clause License
* Copyright © 2020. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
 */
package authenticate

import (
	"fmt"
	"testing"

	"github.com/TIBCOSoftware/TCSTK-common-mods/utils/properties"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&Activity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

func TestEval(t *testing.T) {

	props, err := properties.ReadPropertiesFile("../../../credentials.properties")
	if err != nil {
		panic("Error while reading properties file")
	}

	act := &Activity{}
	tc := test.NewActivityContext(act.Metadata())

	input := &Input{
		Region:   props["region"],
		Tenant:   props["tenant"],
		Username: props["username"],
		Password: props["password"],
		ClientID: props["clientid"],
		Version:  props["version"],
	}
	tc.SetInputObject(input)

	act.Eval(tc)

	output := &Output{}
	tc.GetOutputObject(output)
	fmt.Printf("*********** Cookie∫: %v\n", output.Cookie["tsc"])
}
