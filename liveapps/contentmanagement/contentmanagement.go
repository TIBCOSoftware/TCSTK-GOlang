/*
* BSD 3-Clause License
* Copyright © 2020. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
*/
package contentmanagement

const (
	// ContentManagementV1 Base URL for ContentManagement V1 API
	ContentManagementV1 = "/webresource/v1"
)

// Application The definition of an Application, System Application, Case Folder or Org Folder. (This object is used for the definition of all four resources.)
type Application struct {
	ID                 string `json:"id"`                 // The internal id of the object (used by the Web Resource Provisioner Service).
	Name               string `json:"name"`               // The name (and unique identifier) of the object.
	Owner              string `json:"owner"`              // The id of the User who owns the object.
	OwnerSub           string `json:"ownerSub"`           // The id of the subscription to which the object belongs.
	OwnerSandbox       string `json:"ownerSandbox"`       // The id of the Sandbox to which the object belongs.
	CreationDate       string `json:"creationDate"`       // The date on which the object was created.
	LastModifiedDate   string `json:"lastModifiedDate"`   // The date on which the object was last modified.
	LastModifiedBy     string `json:"lastModifiedBy"`     // The id of the User who last modified the object.
	PublishedVersion   string `json:"publishedVersion"`   // The currently published Version of the object.
	LatestVersion      string `json:"latestVersion"`      // The latest Version of the object.
	PublishedVersionID string `json:"publishedVersionId"` // The internal id of the published version of the object (used by the Web Resource Provisioner Service).
	LatestVersionID    string `json:"latestVersionId"`    // The internal id of the latest version of the object (used by the Web Resource Provisioner Service).
	Checksum           string `json:"checksum"`           // The publish checksum, used to determine if the object has changed.
	ExtRef             string `json:"extRef"`             // The external reference to the object.
	Tags               []*Tag `json:"tags"`
	Descriptor         string `json:"descriptor"` // The contents of the Application’s or System Application’s {appName}.app.desc.json descriptor file (if it has one). This property is not used for a Case Folder.
}

// ApplicationVersion The Definition of a particular version of an Application.
type ApplicationVersion struct {
	ID               string `json:"id"`               // The internal id of the ApplicationVersion (used by the Web Resource Provisioner Service).
	OwnerApp         string `json:"ownerApp"`         // The id of the Application that owns the ApplicationVersion.
	Version          string `json:"version"`          // The version number (and unique identifier) of the ApplicationVersion.
	Name             string `json:"name"`             // The name of the ApplicationVersion.
	CreationDate     string `json:"creationDate"`     // The date on which the ApplicationVersion was created.
	LastModifiedDate string `json:"lastModifiedDate"` // The date on which the ApplicationVersion was last modified.
	LastModifiedBy   string `json:"lastModifiedBy"`   // The id of the User who last modified the application.
	Permissions      int    `json:"permissions"`      // The permissions for the ApplicationVersion (not currently used).
	Label            string `json:"label"`            // The label to display for this ApplicationVersion.
}

// Artifact The definition of an Artifact. An Artifact can be:
// a file or folder within an Application
// a file within a Case Folder or Org Folder
type Artifact struct {
	ID               string `json:"id"`               // The internal id of the Artifact (used by the Web Resource Provisioner Service).
	Name             string `json:"name"`             // The name (and unique identifier) of the Artifact. For an Application, this name is the full path to the Artifact.
	ArtifactVersion  string `json:"artifactVersionn"` // The version of the Artifact.
	Author           string `json:"author"`           //The id of the User who created the Artifact (by first uploading it).
	CreationDate     string `json:"creationDate"`     // The date on which the Artifact was created.
	LastModifiedDate string `json:"lastModifiedDate"` // The date on which the Artifact was last modified.
	LastModifiedBy   string `json:"lastModifiedBy"`   // The id of the User who last modified the Artifact.
	ArtifactRef      string `json:"artifactRef"`      // The external reference to the actual Artifact.
	ArtifactCheckSum string `json:"artifactCheckSum"` // The checksum used to determine if the Artifact already exists.
	OwnerApp         string `json:"ownerApp"`         // The id of the Application or Case Folder to which the Artifact belongs.
	Size             string `json:"size"`             // The size of the Artifact.
	MimeType         string `json:"mimeType"`         // The MIME type of the Artifact.
	Description      string `json:"description"`      // The description of the Artifact.
}

// CaseFolder The definition of an Application, System Application, Case Folder or Org Folder. (This object is used for the definition of all four resources.)
type CaseFolder struct {
	ID                 string `json:"id"`                 // The internal id of the object (used by the Web Resource Provisioner Service).
	Name               string `json:"name"`               // The name (and unique identifier) of the object.
	Owner              string `json:"owner"`              // The id of the User who owns the object.
	OwnerSub           string `json:"ownerSub"`           // The id of the subscription to which the object belongs.
	OwnerSandbox       string `json:"ownerSandbox"`       // The id of the Sandbox to which the object belongs.
	CreationDate       string `json:"creationDate"`       // The date on which the object was created.
	LastModifiedDate   string `json:"lastModifiedDate"`   // The date on which the object was last modified.
	LastModifiedBy     string `json:"lastModifiedBy"`     // The id of the User who last modified the object.
	PublishedVersion   string `json:"publishedVersion"`   // The currently published Version of the object.
	LatestVersion      string `json:"latestVersion"`      // The latest Version of the object.
	PublishedVersionID string `json:"publishedVersionId"` // The internal id of the published version of the object (used by the Web Resource Provisioner Service).
	LatestVersionID    string `json:"latestVersionId"`    // The internal id of the latest version of the object (used by the Web Resource Provisioner Service).
	Checksum           string `json:"checksum"`           // The publish checksum, used to determine if the object has changed.
	ExtRef             string `json:"extRef"`             // The external reference to the object.
	Tags               []*Tag `json:"tags"`
	Descriptor         string `json:"descriptor"` // The contents of the Application’s or System Application’s {appName}.app.desc.json descriptor file (if it has one). This property is not used for a Case Folder.
}

// OrgFolder The definition of an Application, System Application, Case Folder or Org Folder. (This object is used for the definition of all four resources.)
type OrgFolder struct {
	ID                 string `json:"id"`                 // The internal id of the object (used by the Web Resource Provisioner Service).
	Name               string `json:"name"`               // The name (and unique identifier) of the object.
	Owner              string `json:"owner"`              // The id of the User who owns the object.
	OwnerSub           string `json:"ownerSub"`           // The id of the subscription to which the object belongs.
	OwnerSandbox       string `json:"ownerSandbox"`       // The id of the Sandbox to which the object belongs.
	CreationDate       string `json:"creationDate"`       // The date on which the object was created.
	LastModifiedDate   string `json:"lastModifiedDate"`   // The date on which the object was last modified.
	LastModifiedBy     string `json:"lastModifiedBy"`     // The id of the User who last modified the object.
	PublishedVersion   string `json:"publishedVersion"`   // The currently published Version of the object.
	LatestVersion      string `json:"latestVersion"`      // The latest Version of the object.
	PublishedVersionID string `json:"publishedVersionId"` // The internal id of the published version of the object (used by the Web Resource Provisioner Service).
	LatestVersionID    string `json:"latestVersionId"`    // The internal id of the latest version of the object (used by the Web Resource Provisioner Service).
	Checksum           string `json:"checksum"`           // The publish checksum, used to determine if the object has changed.
	ExtRef             string `json:"extRef"`             // The external reference to the object.
	Tags               []*Tag `json:"tags"`
	Descriptor         string `json:"descriptor"` // The contents of the Application’s or System Application’s {appName}.app.desc.json descriptor file (if it has one). This property is not used for a Case Folder.
}

// SystemApplication The definition of an Application, System Application, Case Folder or Org Folder. (This object is used for the definition of all four resources.)
type SystemApplication struct {
	ID                 string `json:"id"`                 // The internal id of the object (used by the Web Resource Provisioner Service).
	Name               string `json:"name"`               // The name (and unique identifier) of the object.
	Owner              string `json:"owner"`              // The id of the User who owns the object.
	OwnerSub           string `json:"ownerSub"`           // The id of the subscription to which the object belongs.
	OwnerSandbox       string `json:"ownerSandbox"`       // The id of the Sandbox to which the object belongs.
	CreationDate       string `json:"creationDate"`       // The date on which the object was created.
	LastModifiedDate   string `json:"lastModifiedDate"`   // The date on which the object was last modified.
	LastModifiedBy     string `json:"lastModifiedBy"`     // The id of the User who last modified the object.
	PublishedVersion   string `json:"publishedVersion"`   // The currently published Version of the object.
	LatestVersion      string `json:"latestVersion"`      // The latest Version of the object.
	PublishedVersionID string `json:"publishedVersionId"` // The internal id of the published version of the object (used by the Web Resource Provisioner Service).
	LatestVersionID    string `json:"latestVersionId"`    // The internal id of the latest version of the object (used by the Web Resource Provisioner Service).
	Checksum           string `json:"checksum"`           // The publish checksum, used to determine if the object has changed.
	ExtRef             string `json:"extRef"`             // The external reference to the object.
	Tags               []*Tag `json:"tags"`
	Descriptor         string `json:"descriptor"` // The contents of the Application’s or System Application’s {appName}.app.desc.json descriptor file (if it has one). This property is not used for a Case Folder.
}

// Tag The definition of a Tag, which can be used to categorize different types of customer-developed Application or System Application.
type Tag struct {
	ID       string `json:"id"`       // The internal id of the Tag (used by the Web Resource Provisioner Service).
	Name     string `json:"name"`     // The name (and unique identifier) of the Tag.
	OwnerSub string `json:"ownerSub"` // The id of the subscription to which the Tag belongs.
}

// Status The definition of a response message.
type Status struct {
	Message string `json:"message"` // Text giving details about the result of the REST call.
}

// ContextAttribute A name/value pair, used within Error, that provides contextual information about the Error.
type ContextAttribute struct {
	Name  string `json:"name"`  // The name of the context attribute.
	Value string `json:"value"` // The value of the context attribute.
}

// Error The definition of an error, providing a suitable message, error code and context information.
type Error struct {
	ErrorMsg          string              `json:"errorMsg"`         // The textual error message.
	ErrorCode         string              `json:"errorCode"`        // The following are the possible error codes in the Web Resource Provisioner Service (note that the description shown is not part of the error code):
	StackTrace        string              `json:"stactTrace"`       // Stack trace details (only provided in a debug environment).
	ContextAttributes []*ContextAttribute `json:"contextAttribute"` // A name/value pair that provides contextual information about the Error.
}
