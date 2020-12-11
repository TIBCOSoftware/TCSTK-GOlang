/*
* BSD 3-Clause License
* Copyright © 2020. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
*/
package execgo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/novalagung/golpal"
)

const (
	ivInput     = "Input"
	ivScriptURL = "ScriptURL"
	ovOutput    = "Output"
)

var activityLog = logger.GetLogger("Scripting-activity-GO")

type execgoActivity struct {
	metadata *activity.Metadata
}

//NewActivity TCI Wi Activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &execgoActivity{metadata: metadata}
}

func (a *execgoActivity) Metadata() *activity.Metadata {
	return a.metadata
}
func (a *execgoActivity) Eval(context activity.Context) (done bool, err error) {
	activityLog.Info("Executing GO Scripting activity")
	//Read Inputs
	if context.GetInput(ivInput) == nil {
		// ivInput string is not configured
		// return error to the engine
		return false, activity.NewError("Scripting ivInput string is not configured", "Scripting-GO-4001", nil)
	}
	input := context.GetInput(ivInput).(string)

	if context.GetInput(ivScriptURL) == nil {
		// APIsecret string is not configured
		// return error to the engine
		return false, activity.NewError("Scripting URL string is not configured", "Scripting-GO-4002", nil)
	}
	scriptURL := context.GetInput(ivScriptURL).(string)

	// execute validation - Start

	// load Script from Web
	res, err := http.Get(scriptURL)
	if err != nil {
		return false, activity.NewError("Scripting URL not found", "Scripting-GO-4003", nil)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var script = string(body)

	// execute
	output := rungo(input, script)

	// Output
	context.SetOutput(ovOutput, output)

	return true, nil
}

func rungo(in string, script string) string {

	script = strings.Replace(script, "{input}", in, 1)

	output, err := golpal.New().ExecuteSimple(script)
	if err != nil {
		fmt.Println(err)
	}

	ret := output
	return ret
}
