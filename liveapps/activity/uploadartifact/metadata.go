/*
* BSD 3-Clause License
* Copyright © 2020. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
 */
package uploadartifact

import (
	"github.com/TIBCOSoftware/TCSTK-GOlang/liveapps/contentmanagement"
	"github.com/project-flogo/core/data/coerce"
)

// Settings for the activity
type Settings struct {
	Security string `md:"security"`
}

// Input for the activity
type Input struct {
	Region           string                 `md:"region"`           //
	Cookie           map[string]interface{} `md:"cookie"`           // TIBCO Cloud cookie
	AccessToken      string                 `md:"accessToken"`      // AccessToken for Outh 2.0
	FolderName       string                 `md:"folderName"`       // The name of the Case Folder to which the Artifact is to be uploaded.
	ArtifactName     string                 `md:"artifactName"`     // The name of the Artifact to be created or updated. Note that this name does not have to match the name of the file being uploaded.
	Sandbox          string                 `md:"sandbox"`          // The id of the Sandbox that contains the Case Folder.
	Description      string                 `md:"description"`      // A suitable description for the new or updated Artifact.
	ArtifactContents string                 `md:"artifactContents"` // The file to be used to update the Artifact.
}

// ToMap for Input
func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"region":           i.Region,
		"cookie":           i.Cookie,
		"accessToken":      i.AccessToken,
		"folderName":       i.FolderName,
		"artifactName":     i.ArtifactName,
		"sandbox":          i.Sandbox,
		"description":      i.Description,
		"artifactContents": i.ArtifactContents,
	}
}

// FromMap for input
func (i *Input) FromMap(values map[string]interface{}) error {
	i.Region, _ = coerce.ToString(values["region"])
	i.Cookie, _ = coerce.ToObject(values["cookie"])
	i.AccessToken, _ = coerce.ToString(values["accessToken"])
	i.FolderName, _ = coerce.ToString(values["folderName"])
	i.ArtifactName, _ = coerce.ToString(values["artifactName"])
	i.Sandbox, _ = coerce.ToString(values["sandbox"])
	i.Description, _ = coerce.ToString(values["description"])
	i.ArtifactContents, _ = coerce.ToString(values["artifactContents"])

	return nil
}

// Output for the activity
type Output struct {
	StatusCode int                      `md:"statusCode"`
	Status     contentmanagement.Status `md:"status"`
	Error      contentmanagement.Error  `md:"error"`
}

// ToMap conver to object
func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"statusCode": o.StatusCode,
		"status":     o.Status,
		"error":      o.Error,
	}
}

// FromMap convert to object
func (o *Output) FromMap(values map[string]interface{}) error {
	o.StatusCode, _ = coerce.ToInt(values["statusCode"])
	o.Status = values["status"].(contentmanagement.Status)
	o.Error = values["error"].(contentmanagement.Error)

	return nil
}
