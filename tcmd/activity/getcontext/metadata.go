/*
* BSD 3-Clause License
* Copyright © 2020. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
 */
package getcontext

import (
	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
}

// Input for the activity
type Input struct {
	Region      string `md:"region"`      //
	AccessToken string `md:"accessToken"` // AccessToken for Outh 2.0
}

// ToMap for Input
func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"region":      i.Region,
		"accessToken": i.AccessToken,
	}
}

// FromMap for input
func (i *Input) FromMap(values map[string]interface{}) error {
	i.Region, _ = coerce.ToString(values["region"])
	i.AccessToken, _ = coerce.ToString(values["accessToken"])

	return nil
}

// Output for the activity
type Output struct {
	Context1 string `md:"context1"`
	Context2 string `md:"context2"`
}

// ToMap conver to object
func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"context1": o.Context1,
		"context2": o.Context2,
	}
}

// FromMap convert to object
func (o *Output) FromMap(values map[string]interface{}) error {
	o.Context1, _ = coerce.ToString(values["context1"])
	o.Context2, _ = coerce.ToString(values["context2"])

	return nil
}
