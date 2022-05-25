/*
* BSD 3-Clause License
* Copyright © 2022. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
*/
package urlshortener

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
	Provider      string `md:"provider"`
	SecurityToken string `md:"securityToken"`
}

// Input for the activity
type Input struct {
	LongDynamicLink string `md:"longDynamicLink"`
}

// ToMap for Input
func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"longDynamicLink": i.LongDynamicLink,
	}
}

// FromMap for input
func (i *Input) FromMap(values map[string]interface{}) error {
	i.LongDynamicLink, _ = coerce.ToString(values["longDynamicLink"])

	return nil
}

// Output for the activity
type Output struct {
	ShortLink string `md:"shortLink"`
}

// ToMap conver to object
func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"shortLink": o.ShortLink,
	}
}

// FromMap convert to object
func (o *Output) FromMap(values map[string]interface{}) error {
	o.ShortLink, _ = coerce.ToString(values["shortLink"])

	return nil
}
