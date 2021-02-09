/*
* BSD 3-Clause License
* Copyright © 2020. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
 */
package runaction

import (
	"github.com/project-flogo/core/data/coerce"
)

// Settings for the activity
type Settings struct {
	Security string `md:"security"`
}

// Input for the activity
type Input struct {
	Region        string                 `md:"region"`
	Cookie        map[string]interface{} `md:"cookie"`      // TIBCO Cloud cookie
	AccessToken   string                 `md:"accessToken"` // AccessToken for Outh 2.0
	ID            string                 `md:"id"`
	ApplicationID string                 `md:"applicationId"`
	SandboxID     string                 `md:"sandboxd"`
	CaseReference string                 `md:"caseReference"`
	Data          string                 `md:"data"`
}

// ToMap for Input
func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"region":        i.Region,
		"cookie":        i.Cookie,
		"accessToken":   i.AccessToken,
		"id":            i.ID,
		"applicationId": i.ApplicationID,
		"sandboxId":     i.SandboxID,
		"caseReference": i.CaseReference,
		"data":          i.Data,
	}
}

// FromMap for input
func (i *Input) FromMap(values map[string]interface{}) error {
	i.Region, _ = coerce.ToString(values["region"])
	i.Cookie, _ = coerce.ToObject(values["cookie"])
	i.AccessToken, _ = coerce.ToString(values["accessToken"])
	i.ID, _ = coerce.ToString(values["id"])
	i.ApplicationID, _ = coerce.ToString(values["applicationId"])
	i.SandboxID, _ = coerce.ToString(values["sandboxId"])
	i.CaseReference, _ = coerce.ToString(values["caseReference"])
	i.Data, _ = coerce.ToString(values["data"])

	return nil
}

// Output for the activity
type Output struct {
	ID            string `md:"id"`
	ApplicationID string `md:"applicationId"`
	SandboxID     string `md:"sandboxd"`
	CaseReference string `md:"caseReference"`
	Data          string `md:"data"`
}

// ToMap conver to object
func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":            o.ID,
		"applicationId": o.ApplicationID,
		"sandboxId":     o.SandboxID,
		"caseReference": o.CaseReference,
		"data":          o.Data,
	}
}

// FromMap convert to object
func (o *Output) FromMap(values map[string]interface{}) error {
	o.ID, _ = coerce.ToString(values["id"])
	o.ApplicationID, _ = coerce.ToString(values["applicationId"])
	o.SandboxID, _ = coerce.ToString(values["sandboxId"])
	o.CaseReference, _ = coerce.ToString(values["caseReference"])
	o.Data, _ = coerce.ToString(values["data"])
	return nil
}
