/*
* BSD 3-Clause License
* Copyright © 2020. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
*/
package processmanagement

const (
	// ProcessManagementV1 Base URL for ProcessManagement V1 API
	ProcessManagementV1 = "/process/v1"
)

// InstanceState The current State and data for a Process Instance.
type InstanceState struct {
	SandboxID            string      `json:"sandboxId"`                      // The id of the Sandbox containing the Process Instance.
	State                string      `json:"state"`                          // The current State of the Process Instance. Values can be: STARTING, ACTIVE, COMPLETED, CANCELLED, HALTED, UNKNOWN, DELAYED
	SystemAtributes      []Attribute `json:"systemAttributes,omitempty"`     // The internal, system-defined data fields and their current values for the Process Instance.
	DefinitionAttributes []Attribute `json:"definitionAttributes,omitempty"` // The process-defined data fields and their current values for the Process Instance.
}

// Instance Summary information about a Process Instance (of a Case Creator or Case Action process).
type Instance struct {
	ID                      string `json:"id"`                      // The unique id of the Process Instance.
	CaseReference           string `json:"caseReference"`           // The Case Reference of the Case that is associated with the Process Instance.
	ApplicationID           string `json:"applicationId"`           // The id of the application in which the Case Creator or Case Action process is defined.
	ApplicationName         string `json:"applicationName"`         // The name of the application in which the Case Creator or Case Action process is defined.
	ApplicationInternalName string `json:"applicationInternalName"` // The internal name of the application in which the Case Creator or Case Action process is defined.
	ProcessID               string `json:"processId"`               // The id of the Case Creator or Case Action process from which the Process Instance was started.
	ProcessName             string `json:"processName"`             // The name of the Case Creator or Case Action process from which the Process Instance was started.
	ProcessInternalName     string `json:"processInternalName"`     // The internal name of the Case Creator or Case Action process from which the Process Instance was started.
	Version                 string `json:"version"`                 // The version of the application in which the Case Creator or Case Action process is defined.
	State                   string `json:"state"`                   // The current State of the Process Instance. Value can be: STARTING, ACTIVE, COMPLETED, CANCELLED, HALTED, UNKNOWN, DELAYED
	TimeStarted             string `json:"timeStarted"`             // The date/time at which the Process Instance started.
	TimeCompleted           string `json:"timeCompleted"`           // The date/time at which the Process Instance completed (if it has completed).
}

// ProcessDetails Details of process to start an instance of
type ProcessDetails struct {
	ID             string `json:"id"`                       // The id of the Case Creator process or Case Action process used to start the Process Instance.
	ApplicationID  string `json:"applicationId"`            // The id of the application that contains the Case Creator process or Case Action process used to start the Process Instance.
	SandboxID      string `json:"sandboxId"`                // The id of the Sandbox containing the specified applicationId.
	CaseReference  string `json:"caseReference,omitempty"`  // The Case reference of the newly created or updated case.
	CaseIdentifier string `json:"caseIdentifier,omitempty"` // The Case Reference of the Case that is associated with the Process Instance.
	Data           string `json:"data"`                     // The JSON data, serialized to a string, used to start the Process Instance.
}

// Attribute The definition of a system-defined or process-defined data field.
type Attribute struct {
	Name  string `json:"name"`  // Name of the attribute.
	Type  string `json:"type"`  // Type of the attribute.
	Value string `json:"value"` // Value of the attribute.
}

// Error Error
type Error struct {
	ErrorMsg    string `json:"errorMsg"`    // Verbose error message.
	ErrorCode   string `json:"errorCode"`   // The following are the possible error codes in the Business Process Management Service (note that the description shown is not part of the error code)
	StackTrace  string `json:"stackTrace"`  // Stack trace details (only provided in a debug environment).
	ContextInfo string `json:"contextInfo"` // Contextual information about the Error.
}
