/*
* BSD 3-Clause License
* Copyright © 2020. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
*/
package execgo

import (
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/stretchr/testify/assert"
)

var activityMetadata *activity.Metadata
var connectionData = ``

func getActivityMetadata() *activity.Metadata {
	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}
		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}
	return activityMetadata
}

func TestActivityRegistration(t *testing.T) {
	act := NewActivity(getActivityMetadata())
	if act == nil {
		t.Error("Activity Not Registered")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {
	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(act.Metadata())

	// *** for testing, replace all in <> with your Account Details!

	//setup attrs
	tc.SetInput("Input", "green")
	tc.SetInput("ScriptURL", "https://raw.githubusercontent.com/TIBCOSoftware/TCST-GOlang/master/scripting/logicscripts/dynifsample.go")

	_, err := act.Eval(tc)
	assert.Nil(t, err)

	result := tc.GetOutput("Output")
	assert.Contains(t, result, "25")

	//assert.Equal(t, result, "done")

	t.Log(result)
}
