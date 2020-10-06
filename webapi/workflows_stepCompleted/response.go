// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    workflowsStepCompleted, err := UnmarshalWorkflowsStepCompleted(bytes)
//    bytes, err = workflowsStepCompleted.Marshal()

package workflows_stepCompleted

import "encoding/json"

func UnmarshalWorkflowsStepCompleted(data []byte) (WorkflowsStepCompleted, error) {
	var r WorkflowsStepCompleted
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WorkflowsStepCompleted) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type WorkflowsStepCompleted struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
