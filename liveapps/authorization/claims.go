/*
* BSD 3-Clause License
* Copyright © 2020. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
 */
package authorization

import (
	"github.com/project-flogo/core/data/coerce"
	"github.com/TIBCOSoftware/TCSTK-common-mods/utils/rest"
	"github.com/TIBCOSoftware/TCSTK-common-mods/utils/server"
)

// CLAIMS endpoint
const (
	CLAIMS = "/claims"
)

// GetClaims Returns the details for each User in the current subscription whose name matches the specified query parameters.
func GetClaims(region string, authentication interface{}) (statusCode int, responseOK Claims, responseKO Error, err error) {

	request := rest.Input{
		URL:            server.GetServer(region, "bpm") + AuthorizationV1 + CLAIMS,
		HTTPMethod:     "GET",
		ContentType:    "application/json",
		Authentication: authentication,
		Response:       &Claims{},
		Error:          &Error{},
	}

	restResponse, err := rest.MakeCall(request)
	if err != nil {
		return -1, responseOK, responseKO, err
	}

	if restResponse.Status == 200 {
		responseOK = *(restResponse.Data).(*Claims)
	} else if restResponse.Status == 401 {
		errorText, _ := coerce.ToString(restResponse.Data)
		responseKO = Error{
			ErrorMsg:  errorText,
			ErrorCode: "401",
		}
	} else {
		responseKO = *(restResponse.Data).(*Error)
	}

	return restResponse.Status, responseOK, responseKO, nil
}
