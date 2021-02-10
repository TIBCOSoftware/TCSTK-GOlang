/*
* BSD 3-Clause License
* Copyright © 2020. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
 */
package contentmanagement

import (
	"encoding/base64"
	"net/url"

	"github.com/TIBCOSoftware/TCSTK-common-mods/utils/rest"
	"github.com/TIBCOSoftware/TCSTK-common-mods/utils/server"
)

// UploadToCaseFolder Uploads a new or updated Artifact to a particular Case Folder.
func UploadToCaseFolder(region string, authentication interface{}, pathParms map[string]string, queryParams url.Values, data string) (statusCode int, responseOK Status, responseKO Error, err error) {
	request := rest.Input{
		URL:            server.GetServer(region, "bpm") + ContentManagementV1 + "/caseFolders/:folderName/artifacts/:artifactName/upload/",
		HTTPMethod:     "POST",
		ContentType:    "application/multipart",
		Authentication: authentication,
		PathParams:     pathParms,
		QueryParams:    queryParams,
		Data:           data,
		Response:       &Status{},
		Error:          &Error{},
	}

	restResponse, err := rest.MakeCall(request)
	if err != nil {
		return -1, responseOK, responseKO, err
	}

	if restResponse.Status == 200 {
		responseOK = *(restResponse.Data).(*Status)
	} else {
		responseKO = *(restResponse.Data).(*Error)
	}

	return restResponse.Status, responseOK, responseKO, nil
}

// UploadToOrgFolder Uploads a new or updated Artifact to a particular Org Folder.
func UploadToOrgFolder() {

}

func GetCaseFolderArtifact(region string, authentication interface{}, pathParams map[string]string, queryParams url.Values) (statusCode int, contentType string, responseOK string, responseKO string, err error) {
	var response string
	request := rest.Input{
		URL:            server.GetServer(region, "bpm") + "/webresource/folders/:folderName/:sandbox/:artifactName",
		HTTPMethod:     "GET",
		ContentType:    "application/json",
		PathParams:     pathParams,
		QueryParams:    queryParams,
		Authentication: authentication,
		Response:       &response,
		Error:          &response,
	}

	restResponse, err := rest.MakeCall(request)
	if err != nil {
		return -1, responseOK, "", responseKO, err
	}

	if restResponse.Status == 200 {
		responseOK = base64.StdEncoding.EncodeToString([]byte(restResponse.Data.(string)))
	} else {
		responseKO = restResponse.Data.(string)
	}

	return restResponse.Status, restResponse.Headers["Content-Type"][0], responseOK, responseKO, nil
}

// GetOrgFolderArtifact to get the folder artifact
func GetOrgFolderArtifact(region string, authentication interface{}, pathParams map[string]string, queryParams url.Values) (statusCode int, contentType string, responseOK string, responseKO string, err error) {
	var response string
	request := rest.Input{
		URL:            server.GetServer(region, "bpm") + "/webresource/orgFolders/:folderName/:artifactName",
		HTTPMethod:     "GET",
		ContentType:    "application/json",
		PathParams:     pathParams,
		QueryParams:    queryParams,
		Authentication: authentication,
		Response:       &response,
		Error:          &response,
	}

	restResponse, err := rest.MakeCall(request)
	if err != nil {
		return -1, responseOK, "", responseKO, err
	}

	if restResponse.Status == 200 {
		responseOK = base64.StdEncoding.EncodeToString([]byte(restResponse.Data.(string)))
	} else {
		responseKO = restResponse.Data.(string)
	}

	return restResponse.Status, restResponse.Headers["Content-Type"][0], responseOK, responseKO, nil
}
