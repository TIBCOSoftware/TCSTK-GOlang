/*
* BSD 3-Clause License
* Copyright © 2020. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
 */
package processmanagement

import (
	"github.com/TIBCOSoftware/TCSTK-GOlang/tibcocloud/idm"
	"github.com/TIBCOSoftware/TCSTK-common-mods/utils/rest"
)

// PostProcesses Creates an instance of a specified Case Creator or Case Action process.
func PostProcesses(region string, authentication interface{}, body ProcessDetails) (statusCode int, responseOK ProcessDetails, responseKO Error, err error) {
	request := rest.Input{
		URL:            idm.GetServer(region, "bpm") + ProcessManagementV1 + "/processes",
		HTTPMethod:     "POST",
		ContentType:    "application/json",
		Authentication: authentication,
		Data:           body,
		Response:       &ProcessDetails{},
		Error:          &Error{},
	}

	restResponse, err := rest.MakeCall(request)
	if err != nil {
		return -1, responseOK, responseKO, err
	}

	if restResponse.Status == 200 {
		responseOK = *(restResponse.Data).(*ProcessDetails)
	} else {
		responseKO = *(restResponse.Data).(*Error)
	}

	return restResponse.Status, responseOK, responseKO, nil
}
