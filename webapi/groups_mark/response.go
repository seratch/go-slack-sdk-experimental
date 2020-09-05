// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    groupsMark, err := UnmarshalGroupsMark(bytes)
//    bytes, err = groupsMark.Marshal()

package groups_mark

import "encoding/json"

func UnmarshalGroupsMark(data []byte) (GroupsMark, error) {
	var r GroupsMark
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GroupsMark) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type GroupsMark struct {
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
