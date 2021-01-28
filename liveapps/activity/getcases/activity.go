/*
* BSD 3-Clause License
* Copyright © 2020. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
 */
package getcases

import (
	"net/url"

	"github.com/TIBCOSoftware/TCSTK-GOlang/liveapps/casemanagement"
	"github.com/TIBCOSoftware/TCSTK-GOlang/liveapps/utils"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/coerce"
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

	queryParams := url.Values{}
	queryParams.Set("$sandbox", input.SandboxID)
	if input.Filter != "" {
		queryParams.Set("$filter", url.QueryEscape(input.Filter))
		queryParams.Set("$filter", input.Filter)
	}
	if input.Select != "" {
		queryParams.Set("$select", input.Select)
	}
	if input.Skip != "" {
		queryParams.Set("$skip", input.Skip)
	}
	if input.Top != "" {
		queryParams.Set("$top", input.Top)
	}
	if input.Search != "" {
		queryParams.Set("$search", input.Search)
	}
	if input.Count {
		count, _ := coerce.ToString(input.Count)
		queryParams.Set("$count", count)
	}

	_, response, _, _ := casemanagement.GetCases(input.Region, authentication, queryParams)

	output := &Output{GetCaseResponseItem: response}
	err = ctx.SetOutputObject(output)
	if err != nil {
		return false, err
	}

	return true, nil
}
