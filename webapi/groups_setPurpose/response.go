// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    groupsSetPurpose, err := UnmarshalGroupsSetPurpose(bytes)
//    bytes, err = groupsSetPurpose.Marshal()

package groups_setPurpose

import "encoding/json"

func UnmarshalGroupsSetPurpose(data []byte) (GroupsSetPurpose, error) {
	var r GroupsSetPurpose
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GroupsSetPurpose) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type GroupsSetPurpose struct {
	Ok               *bool             `json:"ok,omitempty"`               
	Purpose          *string           `json:"purpose,omitempty"`          
	Warning          *string           `json:"warning,omitempty"`          
	ResponseMetadata *ResponseMetadata `json:"response_metadata,omitempty"`
	Error            *string           `json:"error,omitempty"`            
	Needed           *string           `json:"needed,omitempty"`           
	Provided         *string           `json:"provided,omitempty"`         
}

type ResponseMetadata struct {
	Messages []string `json:"messages,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
}
