// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    workflowsStepFailed, err := UnmarshalWorkflowsStepFailed(bytes)
//    bytes, err = workflowsStepFailed.Marshal()

package workflows_stepFailed

import "encoding/json"

func UnmarshalWorkflowsStepFailed(data []byte) (WorkflowsStepFailed, error) {
	var r WorkflowsStepFailed
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WorkflowsStepFailed) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type WorkflowsStepFailed struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
