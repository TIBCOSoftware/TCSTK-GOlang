/*
* BSD 3-Clause License
* Copyright © 2020. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
 */
package authenticate

import (
	"github.com/TIBCOSoftware/TCSTK-GOlang/tibcocloud/idm"
	"github.com/project-flogo/core/activity"
)

func init() {
	_ = activity.Register(&Activity{}, New)
}

var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

// New function for the activity
func New(ctx activity.InitContext) (activity.Activity, error) {
	act := &Activity{}
	return act, nil
}

// Activity is an activity that is used to invoke a REST Operation
type Activity struct {
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
		logger.Debugf("Input params: %s", input.Username)
	}

	cookie, err := idm.LoginOauth("", input.Region, input.Tenant, input.Username, input.Password, input.ClientID, input.Version)

	cookieMap := make(map[string]interface{})
	cookieMap["tsc"] = cookie.Tsc
	cookieMap["domain"] = cookie.Domain
	cookieMap["region"] = cookie.Region

	output := &Output{Cookie: cookieMap}

	err = ctx.SetOutputObject(output)
	if err != nil {
		return false, err
	}

	return true, nil
}
