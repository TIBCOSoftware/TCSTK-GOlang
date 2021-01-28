/*
* BSD 3-Clause License
* Copyright © 2020. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
*/
package casemanagement

const (
	// CaseManagementV1 Base URL
	CaseManagementV1 = "/case/v1"
)

// DeleteCasesResponse The response type for the DELETE /cases API
type DeleteCasesResponse struct {
	MfpCount int `json:"mfpCount"` // The number of Cases marked for deletion/purging.
}

// GetCaseResponseItem The response type for a GET /types API. The API returns an array of these.
type GetCaseResponseItem struct {
	CaseReference    string                      `json:"caseReference"`    // The Case reference for the Case
	CaseData         string                      `json:"caseData"`         // The JSON Case data, serialized to a string.
	UntaggedCasedata string                      `json:"untaggedCasedata"` // The JSON Case data, sserialized to a string, with system identifiers (_id and _value properties) removed. This matches the JSON schema that is returned in the ‘jsonSchema’ property of the response item for the GET /types method
	Summary          string                      `json:"summary"`          // The JSON Case summary, serialized to a string.
	Metadata         GetCaseResponseItemMetadata `json:"metadata"`
}

// GetCaseResponseItemMetadata Describes the metadata for a Case. Used within GetCaseResponseItem in a GET /cases response and also within PutCasesRequest when unlocking a Case.
type GetCaseResponseItemMetadata struct {
	CreatedBy             string `json:"createdBy"`             // The unique identifier of the User that created the Case (if known).
	CreationTimestamp     string `json:"creationTimestamp"`     // The date/time at which the Case was created.
	ModifiedBy            string `json:"modifiedBy"`            // The unique identifier of the User that last modified the Case (if known).
	ModificationTimestamp string `json:"modificationTimestamp"` // The date/time at which the Case was last modified.
	Lock                  bool   `json:"lock"`                  // true if the Case is considered locked (i.e. msLockExpiry > msSystemTime)
	LockType              string `json:"lockType"`              // Always "1"
	LockedBy              string `json:"lockedBy"`              // The user ID of the user that last locked the Case when the Case is considered locked
	MsLockExpiry          string `json:"msLockExpiry"`          // The time (milliseconds since epoch) that the Case’s lock is/was due to expire.
	MsSystemTime          string `json:"msSystemTime"`          // The current time (in milliseconds since epoch).
	MarkedForPurge        bool   `json:"markedForPurge"`        // true if the Case has been marked for purge, but has yet to be purged.
	ApplicationID         string `json:"applicationId"`         // The unique identifier of the application containing the Case’s Type.
	TypeID                string `json:"typeId"`                // The unique identifier of the Case’s Type.
}

// PutCasesRequest The request type for a PUT /cases request. As the metadata’s lock property is the only thing that can be updated, everything else is ignored.
type PutCasesRequest struct {
	Metadata GetCaseResponseItemMetadata `json:"metdata"` // Describes the metadata for a Case. Used within GetCaseResponseItem in a GET /cases response and also within PutCasesRequest when unlocking a Case.
}

// GetTypeResponseItem The response type for a GET /types API. The API returns an array of these.
type GetTypeResponseItem struct {
	ID                      string                         `json:"id"`                      // The unique identifier of the Type.
	Name                    string                         `json:"name"`                    // The name of the Type.
	Label                   string                         `json:"label"`                   // The label of the Type.
	IsCase                  bool                           `json:"isCase"`                  // Either true for a Case Type, or false for a Structured Type.
	ApplicationID           string                         `json:"applicationId"`           // The unique identifier of the application in which the Type is defined.
	ApplicationName         string                         `json:"applicationName"`         // The name of the application in which the Type is defined.
	ApplicationInternalName string                         `json:"applicationInternalName"` // The internal name of the application in which the Type is defined.
	ApplicationVersion      int                            `json:"applicationVersion"`      // The version of the application in which the Type is defined.
	Attributes              []GetTypeResponseItemAttribute `json:"attributes"`              // The definitions of all attributes defined in the Type.
	SummaryAttributes       []GetTypeResponseItemAttribute `json:"summaryAttributes"`       // The definitions of all attributes that are included in the Case summary for the Type.
	States                  []GetTypeResponseItemState     `json:"states"`                  // The definitions of all states defined in the Type.
	JSONSchema              interface{}                    `json:"jsonSchema"`              // A JSON schema describing the format of Case data that is returned in the ‘untaggedCasedata’ response item property.
	Creators                []GetTypeResponseItemCreator   `json:"creators"`                // Details of Case Creator processes that are defined in the application containing this Case Type. (Case Creator processes are not defined for a Structured Type. Case Creator processes that contain participant fields are also excluded.)
	Actions                 []GetTypeResponseItemAction    `json:"actions"`                 // Details of Case Action processes that are defined in the application containing this Case Type. (Case Action processes are not defined for a Structured Type. Case Action processes that contain participant fields are also excluded.)
}

// GetTypeResponseItemAttribute Describes the attributes of a Case or Structured type. Used within GetTypeResponseItem.
type GetTypeResponseItemAttribute struct {
	Name               string                      `json:"name"`                         // The name of the attribute, which is automatically derived from the label.
	Label              string                      `json:"label"`                        // The name of the attribute as defined by the application designer in Live Apps Designer.
	Type               string                      `json:"type"`                         // If the attribute refers to another Type, the ID of that Type. If the attribute is of a Case Type, then either "Text", "Number", "Date", “Time","User","Group","Email” or "WebLink".
	IsStructuredType   bool                        `json:"isStructuredType,omitempty"`   // If true, this indicates that the value for ‘type’ is the ID of a Structured Type.
	IsIdentifier       bool                        `json:"isIdentifier,omitempty"`       // true if the attribute is the Type’s designated ‘identifier’ attribute.
	IsState            bool                        `json:"isState,omitempty"`            // true if the attribute is the Type’s designated ‘state’ attribute.
	IsArray            bool                        `json:"isArray,omitempty"`            // true if the attribute is an array.
	IsMandatory        bool                        `json:"isMandatory,omitempty"`        // true if the attribute must be populated
	Constraints        AttributeConstraints        `json:"constraints,omitempty"`        // Constraints on permitted values for the attribute
	DisplayPreferences AttributeDisplayPreferences `json:"displayPreferences,omitempty"` // Indicates how clients should display an attribute
}

// AttributeConstraints Constraints on permitted values for the attribute
type AttributeConstraints struct {
	DecimalPlaces int `json:"decimalPlaces"` // The maximum number of decimal places permitted (only applicable to attributes of the Number type)
}

// AttributeDisplayPreferences Indicates how clients should display an attribute
type AttributeDisplayPreferences struct {
	Multiline bool   `json:"multiline"` // true if clients should present the attribute’s value in a multi-line control as opposed to a single line control
	Prefix    string `json:"prefix"`    // a string clients should present as a prefix to attribute values. For example, an attribute representing currency may have a prefix ‘$’ representing dollars.
	Suffix    string `json:"suffix"`    // a string clients should present as a suffix to attribute values. For example, an attribute representing a person’s age may have a suffix 'year(s)'.
}

// GetTypeResponseItemState Describes a Case type’s states. Used within GetTypeResponseItem
type GetTypeResponseItemState struct {
	ID         string `json:"id"`                   // The ID of the state.
	Label      string `json:"label"`                // The name of the state defined by the application designer in Live Apps Designer.
	Value      string `json:"value"`                // The name of the state defined by the application designer in Live Apps Designer.
	IsTerminal bool   `json:"isTerminal,omitempty"` // true if the state is a ‘terminal’ state (one in which Cases are no longer considered active).
}

// GetTypeResponseItemCreator The definition of a Case Creator process (in the application containing this Case Type).
type GetTypeResponseItemCreator struct {
	ID         string      `json:"id"`         // The id of the Case Creator process.
	Name       string      `json:"name"`       // The name of the Case Creator process.
	JSONSchema interface{} `json:"jsonSchema"` // A JSON schema describing the data expected as input to this Case Creator process.
}

// GetTypeResponseItemAction The definition of a Case Action process (in the application containing this Case Type).
type GetTypeResponseItemAction struct {
	ID         string      `json:"id"`         // The id of the Case Action process.
	Name       string      `json:"name"`       // The name of the Case Action process.
	JSONSchema interface{} `json:"jsonSchema"` // A JSON schema describing the data expected as input to this Case Action process.
}

// ContextAttribute A name/value pair providing contextual information when an error occur. Used within Error.
type ContextAttribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Error Provides details when an error occurs
type Error struct {
	ErrorMsg         string             `json:"errorMsg"`
	ErrorCode        string             `json:"errorCode"`
	StackTrace       string             `json:"stackTrace"`
	ContextAttribute []ContextAttribute `json:"contextAttributes"`
}
