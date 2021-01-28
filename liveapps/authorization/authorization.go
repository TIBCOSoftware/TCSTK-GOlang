/*
* BSD 3-Clause License
* Copyright © 2020. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
*/
package authorization

const (
	AuthorizationV1 = "/organisation/v1"
	GROUPS          = "/groups"
	MAPPINGS        = "/userGroupMappings"
)

// UserDetails The details for a User, excluding their id
type UserDetails struct {
	ExternalID  string `json:"externalId"`  // The external identifier for the User. Note: In general, use the unique identifier (‘id’) when specifying a User. Do not use the externalId value unless specifically requested to do so.
	FirstName   string `json:"firstName"`   // First name of the User
	LastName    string `json:"lastName"`    // Last name of the User
	Username    string `json:"username"`    // Username of the User
	Email       string `json:"email"`       // email address of the User
	ReportsTo   string `json:"reportsTo"`   // The display name of the person this user reports to
	JobTitle    string `json:"jobTitle"`    // The job title of the user
	Description string `json:"description"` // A text description of the user
	Department  string `json:"Department"`  // The name of the department the user is in
	Type        string `json:"type"`        // Type of the user. (Test users cannot be created.)
}

// User The complete details for a User, including their unique id
type User struct {
	ExternalID  string `json:"externalId"`  // The external identifier for the User. Note: In general, use the unique identifier (‘id’) when specifying a User. Do not use the externalId value unless specifically requested to do so.
	FirstName   string `json:"firstName"`   // First name of the User
	LastName    string `json:"lastName"`    // Last name of the User
	Username    string `json:"username"`    // Username of the User
	Email       string `json:"email"`       // email address of the User
	ReportsTo   string `json:"reportsTo"`   // The display name of the person this user reports to
	JobTitle    string `json:"jobTitle"`    // The job title of the user
	Description string `json:"description"` // A text description of the user
	Department  string `json:"department"`  // The name of the department the user is in
	Type        string `json:"type"`        // Type of the user. (Test users cannot be created.)
	ID          string `json:"id"`          // The unique identifier for a User
}

// ClaimsGroup The Group details needed by the Claims
type ClaimsGroup struct {
	ID   string `json:"id"`
	Type string `json:"type"` // The type of Group: AllUsers, Administrator, ApplicationDeveloper, UIDeveloper, SubscriptionDefined
}

// ClaimsSandbox The Sandbox details needed by the Claims
type ClaimsSandbox struct {
	ID     string        `json:"id"`
	Type   string        `json:"type"`   // The type of Sandbox: Production, Developer
	Groups []ClaimsGroup `json:"groups"` // The Groups to which the user with this Sandbox has access
}

// ClaimsValues Details supplied for the Claims
type ClaimsValues struct {
	SubscriptionID string          `json:"subscriptionId"` // The subscription to which this User belongs
	FirstName      string          `json:"firstName"`      // First name of the User
	LastName       string          `json:"lastName"`       // Last name of the User
	Username       string          `json:"username"`       // Username of the User
	Email          string          `json:"email"`          // email address of the User
	Status         []string        `json:"status"`         // An array of status values applicable to this User. Valid values are 'VerifyEmailOutstanding’.
	Sandboxes      []ClaimsSandbox `json:"sandboxes"`      // The Sandboxes to which the User has access
}

// Claims Details supplied for the Claims
type Claims struct {
	SubscriptionID string          `json:"subscriptionId"` // The subscription to which this User belongs
	FirstName      string          `json:"firstName"`      // First name of the User
	LastName       string          `json:"lastName"`       // Last name of the User
	Username       string          `json:"username"`       // Username of the User
	Email          string          `json:"email"`          // email address of the User
	Status         []string        `json:"status"`         // An array of status values applicable to this User. Valid values are 'VerifyEmailOutstanding’.
	Sandboxes      []ClaimsSandbox `json:"sandboxes"`      // The Sandboxes to which the User has access
	ID             string          `json:"id"`             // The unique identifier for a User
}

// GroupDetails The required details for a given Group
type GroupDetails struct {
	Name        string `json:"name"`        // Name of the Group
	Description string `json:"description"` // Description of the Group
	Type        string `json:"type"`        // The type of Group (SubscriptionDefined if not set): AllUsers, Administrator, ApplicationDeveloper, UIDeveloper, SubscriptionDefined
}

// Group Complete details for a Group, including the unique id
type Group struct {
	Name        string `json:"name"`        // Name of the Group
	Description string `json:"description"` // Description of the Group
	Type        string `json:"type"`        // The type of Group (SubscriptionDefined if not set): AllUsers, Administrator, ApplicationDeveloper, UIDeveloper, SubscriptionDefined
	ID          string `json:"id"`          // The unique identifier for a Group.
}

// UserGroupMappingContent The required details for a UserGroupMapping
type UserGroupMappingContent struct {
	SandboxID string `json:"sandboxId"` // The id of the Sandbox which the mapping is for
	GroupID   string `json:"groupId"`   // The id of the Group
	UserID    string `json:"userId"`    // The id of the User
}

// UserGroupMapping Complete details for a UserGroupMapping, including its unique id
type UserGroupMapping struct {
	SandboxID string `json:"sandboxId"` // The id of the Sandbox which the mapping is for
	GroupID   string `json:"groupId"`   // The id of the Group
	UserID    string `json:"userId"`    // The id of the User
	ID        string `json:"id"`        // The unique identifier for a UserGroupMapping
}

// SandboxContent The required details for a Sandbox
type SandboxContent struct {
	Name           string `json:"name"`           // Name of the Sandbox
	Type           string `json:"type"`           // The type of Sandbox: Production, Developer
	SubscriptionID string `json:"subscriptionId"` // The subscription to which this Sandbox belongs
	OwnerID        string `json:"ownerId"`        // The id of the User that owns the Sandbox
}

// Sandbox Complete details for a Sandbox, including its unique ID
type Sandbox struct {
	Name           string `json:"name"`           // Name of the Sandbox
	Type           string `json:"type"`           // The type of Sandbox: Production, Developer
	SubscriptionID string `json:"subscriptionId"` // The subscription to which this Sandbox belongs
	OwnerID        string `json:"ownerId"`        // The id of the User that owns the Sandbox
	ID             string `json:"id"`             // The unique identifier for a Sandbox
}

// Parameter A Parameter for a subscription, as a name-value pair
type Parameter struct {
	Name  string `json:"name"`  // The name of the Parameter (case insensitive)
	Value string `json:"value"` // The value of the Parameter
}

// ErrorAttribute A given error attribute
type ErrorAttribute struct {
	Name  string `json:"name"`  // Name of the error attrubute
	Value string `json:"value"` // Value of the error attribute
}

// Error Error
type Error struct {
	ErrorMsg         string           `json:"errorMsg"`          // Verbose error message
	ErrorCode        string           `json:"errorCode"`         // Possible error codes in the Authorization Engine Service
	StackTrace       string           `json:"stackTrace"`        // Added if available
	ContextAttribute []ErrorAttribute `json:"contextAttributes"` // Error Attributes
}
