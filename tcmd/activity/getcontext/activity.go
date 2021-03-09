/*
* BSD 3-Clause License
* Copyright © 2020. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
 */
package getcontext

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/TIBCOSoftware/TCSTK-common-mods/utils/server"
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

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	server := server.GetServer(input.Region, "metadata")
	req, err := http.NewRequest("GET", server, nil)
	req.Header.Set("Authorization", "Bearer "+input.AccessToken)
	resp, err := client.Do(req)

	u, err := url.Parse(resp.Header.Get("Location"))
	context := strings.Split(u.Path, "/")
	defer resp.Body.Close()
	output := &Output{
		Context1: context[1],
		Context2: context[2],
	}
	err = ctx.SetOutputObject(output)
	if err != nil {
		return false, err
	}

	return true, nil
}
