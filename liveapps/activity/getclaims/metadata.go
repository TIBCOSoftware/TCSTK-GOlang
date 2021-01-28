/*
* BSD 3-Clause License
* Copyright © 2020. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
 */
package getclaims

import (
	"github.com/TIBCOSoftware/TCSTK-GOlang/liveapps/authorization"
	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
	Security string `md:"security"`
}

// Input for the activity
type Input struct {
	Region      string                 `md:"region"`      //
	Cookie      map[string]interface{} `md:"cookie"`      // TIBCO Cloud cookie
	AccessToken string                 `md:"accessToken"` // AccessToken for Outh 2.0
}

// ToMap for Input
func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"region":      i.Region,
		"cookie":      i.Cookie,
		"accessToken": i.AccessToken,
	}
}

// FromMap for input
func (i *Input) FromMap(values map[string]interface{}) error {
	i.Region, _ = coerce.ToString(values["region"])
	i.Cookie, _ = coerce.ToObject(values["cookie"])
	i.AccessToken, _ = coerce.ToString(values["accessToken"])

	return nil
}

// Output for the activity
type Output struct {
	StatusCode int                  `md:"statusCode"`
	Claims     authorization.Claims `md:"claims"`
	Error      authorization.Error  `md:"error"`
}

// ToMap conver to object
func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"statusCode": o.StatusCode,
		"claims":     o.Claims,
		"error":      o.Error,
	}
}

// FromMap convert to object
func (o *Output) FromMap(values map[string]interface{}) error {
	var err error
	o.StatusCode, err = coerce.ToInt(values["statusCode"])
	if err != nil {
		return err
	}
	o.Claims, _ = values["claims"].(authorization.Claims)
	o.Error, _ = values["error"].(authorization.Error)

	return nil
}
