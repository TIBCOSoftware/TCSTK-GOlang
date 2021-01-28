/*
* BSD 3-Clause License
* Copyright © 2020. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
 */
package startcase

import (
	"testing"

	"github.com/TIBCOSoftware/TCSTK-GOlang/tibcocloud/idm"
	"github.com/TIBCOSoftware/TCSTK-common-mods/utils/properties"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/mapper"
	"github.com/project-flogo/core/data/resolve"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&Activity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

func TestEval(t *testing.T) {

	props, err := properties.ReadPropertiesFile("../../credentials.properties")
	if err != nil {
		panic("Error while reading properties file")
	}

	settings := &Settings{
		Security: "Cookie",
	}

	cookie, _ := idm.LoginOauth("cookie", props["region"], props["tenant"], props["username"], props["password"], props["clientid"], props["version"])
	cookieObject, _ := coerce.ToObject("{\"tsc\": \"" + cookie.Tsc + "\", \"domain\": \"" + cookie.Domain + "\"}")

	mf := mapper.NewFactory(resolve.GetBasicResolver())
	iCtx := test.NewActivityInitContext(settings, mf)
	act, err := New(iCtx)
	assert.Nil(t, err)

	tc := test.NewActivityContext(act.Metadata())

	input := &Input{
		Region:        props["region"],
		Cookie:        cookieObject,
		ID:            "16352",
		ApplicationID: "3495",
		SandboxID:     "3100",
		Data:          "{\"MCTestOneTimeClick\":{\"state\":\"Created\",\"field1\":\"From activity tester 2\",\"field2\":\"2\",\"field3\":\"3\"}}",
	}
	tc.SetInputObject(input)

	act.Eval(tc)

	output := &Output{}
	tc.GetOutputObject(output)
	assert.NotEmpty(t, output)
}
