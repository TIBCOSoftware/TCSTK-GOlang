/*
* BSD 3-Clause License
* Copyright © 2020. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
 */
package getartifact

import (
	"github.com/TIBCOSoftware/TCSTK-GOlang/liveapps/contentmanagement"
	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
	Security   string `md:"security"`
	FolderType string `md:"folderType"`
}

// Input for the activity
type Input struct {
	Region       string                 `md:"region"`      //
	Cookie       map[string]interface{} `md:"cookie"`      // TIBCO Cloud cookie
	AccessToken  string                 `md:"accessToken"` // AccessToken for Outh 2.0
	FolderName   string                 `md:"folderName"`
	ArtifactName string                 `md:"artifactName"`
	SandboxID    string                 `md:"sandboxId"`
	Version      string                 `md:"version"`
}

// ToMap for Input
func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"region":       i.Region,
		"cookie":       i.Cookie,
		"accessToken":  i.AccessToken,
		"folderName":   i.FolderName,
		"artifactName": i.ArtifactName,
		"sandboxId":    i.SandboxID,
		"version":      i.Version,
	}
}

// FromMap for input
func (i *Input) FromMap(values map[string]interface{}) error {
	i.Region, _ = coerce.ToString(values["region"])
	i.Cookie, _ = coerce.ToObject(values["cookie"])
	i.AccessToken, _ = coerce.ToString(values["accessToken"])
	i.FolderName, _ = coerce.ToString(values["folderName"])
	i.ArtifactName, _ = coerce.ToString(values["artifactName"])
	i.SandboxID, _ = coerce.ToString(values["sandboxId"])
	i.Version, _ = coerce.ToString(values["version"])

	return nil
}

// Output for the activity
type Output struct {
	StatusCode  int                     `md:"statusCode"`
	ContentType string                  `md:"contentType"`
	Content     string                  `md:"content"`
	Error       contentmanagement.Error `md:"error"`
}

// ToMap conver to object
func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"statusCode":  o.StatusCode,
		"contentType": o.ContentType,
		"content":     o.Content,
		"error":       o.Error,
	}
}

// FromMap convert to object
func (o *Output) FromMap(values map[string]interface{}) error {
	var err error
	o.StatusCode, err = coerce.ToInt(values["statusCode"])
	if err != nil {
		return err
	}
	o.ContentType, _ = coerce.ToString(values["contentType"])
	o.Content, _ = coerce.ToString(values["content"])
	o.Error, _ = values["error"].(contentmanagement.Error)

	return nil
}
