/*
* BSD 3-Clause License
* Copyright © 2020. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
 */
package casemanagement

import (
	"net/url"

	"github.com/TIBCOSoftware/TCSTK-common-mods/utils/rest"
	"github.com/TIBCOSoftware/TCSTK-common-mods/utils/server"
)

const (
	// TYPES endpoint
	TYPES = "/types"
)

// GetTypes Returns Types that match the specified query parameters.
func GetTypes(region string, authentication interface{}, queryParams url.Values) (statusCode int, responseOK []GetTypeResponseItem, responseKO Error, err error) {
	request := rest.Input{
		URL:            server.GetServer(region, "bpm") + CaseManagementV1 + TYPES,
		HTTPMethod:     "GET",
		ContentType:    "application/json",
		Authentication: authentication,
		QueryParams:    queryParams,
		Response:       &[]GetTypeResponseItem{},
		Error:          &Error{},
	}

	restResponse, err := rest.MakeCall(request)
	if err != nil {
		return -1, responseOK, responseKO, err
	}

	if restResponse.Status == 200 {
		responseOK = *(restResponse.Data).(*[]GetTypeResponseItem)
	} else {
		responseKO = *(restResponse.Data).(*Error)
	}

	return restResponse.Status, responseOK, responseKO, nil
}
