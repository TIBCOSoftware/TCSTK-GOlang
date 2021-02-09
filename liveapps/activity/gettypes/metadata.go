/*
* BSD 3-Clause License
* Copyright © 2020. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
 */
package gettypes

import (
	"github.com/TIBCOSoftware/TCSTK-GOlang/liveapps/casemanagement"
	"github.com/project-flogo/core/data/coerce"
)

// Settings for the activity
type Settings struct {
	Security string `md:"security"`
}

// Input for the activity
type Input struct {
	Region      string                 `md:"region"`      //
	Cookie      map[string]interface{} `md:"cookie"`      // TIBCO Cloud cookie
	AccessToken string                 `md:"accessToken"` // AccessToken for Outh 2.0
	SandboxID   string                 `md:"sandboxId"`
	Select      string                 `md:"select"` // A comma-separated list of identifiers defining the properties to be returned for each Type. If $select is not specified, all properties are returned. Possible values are: b, a, sa, s, js, c, ac
	Filter      string                 `md:"filter"` // A filter expression that defines the Types to be returned. The expression can contain one of the following operands: applicationId (eq and in), applicationName (eq) and isCase (eq)
	Skip        string                 `md:"skip"`   // The number of items to exclude from the results list, counting from the beginning of the list. The value must be 0 or greater. For example, $skip=80 results in the first 80 items in the results list being ignored. Subsequent items are returned, starting with the 81st item in the list.
	Top         string                 `md:"top"`    // The maximum number of items to be returned from the results list (with the first item determined by the value of the $skip parameter). The value of $top must be between 1 and 1000. For example, $top=20 returns 20 items from the results list, or all the results if the list contains 20 or fewer items.
	Count       bool                   `md:"count"`  // If set to 'TRUE’, returns the number of Types in the result, rather than the Types themselves.
}

// ToMap for Input
func (i *Input) ToMap() map[string]interface{} {

	return map[string]interface{}{
		"region":      i.Region,
		"cookie":      i.Cookie,
		"accessToken": i.AccessToken,
		"sandboxId":   i.SandboxID,
		"select":      i.Select,
		"filter":      i.Filter,
		"skip":        i.Skip,
		"top":         i.Top,
		"count":       i.Count,
	}
}

// FromMap for input
func (i *Input) FromMap(values map[string]interface{}) error {
	i.Region, _ = coerce.ToString(values["region"])
	i.Cookie, _ = coerce.ToObject(values["cookie"])
	i.AccessToken, _ = coerce.ToString(values["accessToken"])
	i.SandboxID, _ = coerce.ToString(values["sandboxId"])
	i.Select, _ = coerce.ToString(values["select"])
	i.Filter, _ = coerce.ToString(values["filter"])
	i.Skip, _ = coerce.ToString(values["skip"])
	i.Top, _ = coerce.ToString(values["top"])
	i.Count, _ = coerce.ToBool(values["count"])

	return nil
}

// Output for the activity
type Output struct {
	TypeResponseItems []casemanagement.GetTypeResponseItem `md:"typeResponseItems"`
}

// ToMap conver to object
func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"typeResponseItems": o.TypeResponseItems,
	}
}

// FromMap convert to object
func (o *Output) FromMap(values map[string]interface{}) error {

	o.TypeResponseItems, _ = values["typeResponseItems"].([]casemanagement.GetTypeResponseItem)

	return nil
}
