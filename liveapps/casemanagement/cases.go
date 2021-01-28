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

// CASES endpoint
const (
	CASES = "/cases"
)

// GetCases Returns Cases that match the specified query parameters.
func GetCases(region string, authentication interface{}, queryParams url.Values) (statusCode int, responseOK []GetCaseResponseItem, responseKO Error, err error) {
	request := rest.Input{
		URL:            server.GetServer(region, "bpm") + CaseManagementV1 + CASES,
		HTTPMethod:     "GET",
		ContentType:    "application/json",
		Authentication: authentication,
		QueryParams:    queryParams,
		Response:       &[]GetCaseResponseItem{},
		Error:          &Error{},
	}

	restResponse, err := rest.MakeCall(request)
	if err != nil {
		return -1, responseOK, responseKO, err
	}

	if restResponse.Status == 200 {
		responseOK = *(restResponse.Data).(*[]GetCaseResponseItem)
	} else {
		responseKO = *(restResponse.Data).(*Error)
	}

	return restResponse.Status, responseOK, responseKO, nil
}

// DeleteCases Deletes all purgeable Cases that match the specified query parameters.
func DeleteCases(region string, authentication interface{}, queryParams url.Values) (statusCode int, responseOK DeleteCasesResponse, responseKO Error, err error) {
	request := rest.Input{
		URL:            server.GetServer(region, "bpm") + CaseManagementV1 + CASES,
		HTTPMethod:     "DELETE",
		ContentType:    "application/json",
		Authentication: authentication,
		QueryParams:    queryParams,
		Response:       &DeleteCasesResponse{},
		Error:          &Error{},
	}

	restResponse, err := rest.MakeCall(request)
	if err != nil {
		return -1, responseOK, responseKO, err
	}

	if restResponse.Status == 200 {
		responseOK = *(restResponse.Data).(*DeleteCasesResponse)
	} else {
		responseKO = *(restResponse.Data).(*Error)
	}

	return restResponse.Status, responseOK, responseKO, nil
}
