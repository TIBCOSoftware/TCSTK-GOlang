/*
* BSD 3-Clause License
* Copyright © 2020. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
 */
package runaction

import (
	"github.com/TIBCOSoftware/TCSTK-GOlang/liveapps/processmanagement"
	"github.com/TIBCOSoftware/TCSTK-GOlang/liveapps/utils"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
)

func init() {
	_ = activity.Register(&Activity{}, New)
}

var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

// New function for the activity
func New(ctx activity.InitContext) (activity.Activity, error) {
	s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}

	act := &Activity{settings: s}

	return act, nil
}

// Activity is an activity that is used to invoke a REST Operation
type Activity struct {
	settings *Settings
}

// Metadata for the activity
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Create the hash
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return false, err
	}

	logger := ctx.Logger()

	if logger.DebugEnabled() {
		logger.Debugf("Input params: %s", input)
	}

	authentication, error := utils.GenerateAuthentication(a.settings.Security, input.Cookie, input.AccessToken)
	if error != nil {
		return false, error
	}

	var body processmanagement.ProcessDetails
	body.ID = input.ID
	body.ApplicationID = input.ApplicationID
	body.SandboxID = input.SandboxID
	body.CaseReference = input.CaseReference
	body.Data = input.Data

	_, response, _, _ := processmanagement.PostProcesses(input.Region, authentication, body)

	output := &Output{ID: response.ID, ApplicationID: response.ApplicationID, SandboxID: response.SandboxID, CaseReference: response.CaseReference, Data: response.Data}
	err = ctx.SetOutputObject(output)
	if err != nil {
		return false, err
	}

	return true, nil
}
