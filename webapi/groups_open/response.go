// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    groupsOpen, err := UnmarshalGroupsOpen(bytes)
//    bytes, err = groupsOpen.Marshal()

package groups_open

import "encoding/json"

func UnmarshalGroupsOpen(data []byte) (GroupsOpen, error) {
	var r GroupsOpen
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GroupsOpen) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type GroupsOpen struct {
	Ok               *bool             `json:"ok,omitempty"`               
	NoOp             *bool             `json:"no_op,omitempty"`            
	AlreadyOpen      *bool             `json:"already_open,omitempty"`     
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
