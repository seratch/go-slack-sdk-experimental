// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    workflowsUpdateStep, err := UnmarshalWorkflowsUpdateStep(bytes)
//    bytes, err = workflowsUpdateStep.Marshal()

package workflows_updateStep

import "encoding/json"

func UnmarshalWorkflowsUpdateStep(data []byte) (WorkflowsUpdateStep, error) {
	var r WorkflowsUpdateStep
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WorkflowsUpdateStep) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type WorkflowsUpdateStep struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
