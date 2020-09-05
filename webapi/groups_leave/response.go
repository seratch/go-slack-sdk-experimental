// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    groupsLeave, err := UnmarshalGroupsLeave(bytes)
//    bytes, err = groupsLeave.Marshal()

package groups_leave

import "encoding/json"

func UnmarshalGroupsLeave(data []byte) (GroupsLeave, error) {
	var r GroupsLeave
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GroupsLeave) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type GroupsLeave struct {
	Ok               *bool             `json:"ok,omitempty"`               
	Error            *string           `json:"error,omitempty"`            
	Warning          *string           `json:"warning,omitempty"`          
	ResponseMetadata *ResponseMetadata `json:"response_metadata,omitempty"`
	Needed           *string           `json:"needed,omitempty"`           
	Provided         *string           `json:"provided,omitempty"`         
}

type ResponseMetadata struct {
	Messages []string `json:"messages,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
}
