// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    groupsKick, err := UnmarshalGroupsKick(bytes)
//    bytes, err = groupsKick.Marshal()

package groups_kick

import "encoding/json"

func UnmarshalGroupsKick(data []byte) (GroupsKick, error) {
	var r GroupsKick
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GroupsKick) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type GroupsKick struct {
	Ok               *bool             `json:"ok,omitempty"`               
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
