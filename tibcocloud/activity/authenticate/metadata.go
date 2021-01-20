/*
* BSD 3-Clause License
* Copyright © 2020. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
*/
package authenticate

import (
	"github.com/project-flogo/core/data/coerce"
)

// Settings for the activity
type Settings struct {
}

// Input for the activity
type Input struct {
	Region   string `md:"region"`   // Data for the hash
	Tenant   string `md:"tenant"`   // LiveApps element: allowed values are action and workitem
	Username string `md:"username"` // LiveApps case reference
	Password string `md:"password"` // LiveApps application ID
	ClientID string `md:"clientID"` // LiveApps application ID
	Version  string `md:"version"`  // LiveApps application ID
}

// ToMap for input structure
func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"region":   i.Region,
		"tenant":   i.Tenant,
		"username": i.Username,
		"password": i.Password,
		"clientID": i.ClientID,
		"version":  i.Version,
	}
}

// FromMap for input structure
func (i *Input) FromMap(values map[string]interface{}) error {
	i.Region, _ = coerce.ToString(values["region"])
	i.Tenant, _ = coerce.ToString(values["tenant"])
	i.Username, _ = coerce.ToString(values["username"])
	i.Password, _ = coerce.ToString(values["password"])
	i.ClientID, _ = coerce.ToString(values["clientID"])
	i.Version, _ = coerce.ToString(values["version"])

	return nil
}

// Output for the activity
type Output struct {
	Cookie map[string]interface{} `md:"cookie"`
}

// ToMap conver to object
func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"cookie": o.Cookie,
	}
}

// FromMap convert to object
func (o *Output) FromMap(values map[string]interface{}) error {
	o.Cookie, _ = coerce.ToObject(values["cookie"])

	return nil
}
