// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    groupsSetTopic, err := UnmarshalGroupsSetTopic(bytes)
//    bytes, err = groupsSetTopic.Marshal()

package groups_setTopic

import "encoding/json"

func UnmarshalGroupsSetTopic(data []byte) (GroupsSetTopic, error) {
	var r GroupsSetTopic
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GroupsSetTopic) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type GroupsSetTopic struct {
	Topic            *string           `json:"topic,omitempty"`            
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
