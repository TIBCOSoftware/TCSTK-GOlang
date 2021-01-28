/*
* BSD 3-Clause License
* Copyright © 2020. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
*/
package utils

import (
	"github.com/TIBCOSoftware/TCSTK-GOlang/tibcocloud/cookie"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/coerce"
)

// GenerateAuthentication Generate authentication for the API request
func GenerateAuthentication(security string, cookieInput map[string]interface{}, token string) (authentication interface{}, error error) {
	if security == "Cookie" {
		cookieParam, _ := coerce.ToObject(cookieInput)
		var cookieObject cookie.Details
		cookieObject.Tsc, _ = coerce.ToString(cookieParam["tsc"])
		cookieObject.Domain, _ = coerce.ToString(cookieParam["domain"])
		if cookieObject.Tsc == "" || cookieObject.Domain == "" {
			err := activity.NewError("Cookie not defined", "ERROR-0001", cookieInput)
			return nil, err
		}
		authentication = cookieObject
	}

	if security == "AccessToken" {
		if token == "" {
			err := activity.NewError("Token not defined", "ERROR-0002", nil)
			return nil, err
		}
		authentication = "Bearer " + token
	}
	return authentication, nil
}
