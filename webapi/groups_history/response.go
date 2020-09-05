// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    groupsHistory, err := UnmarshalGroupsHistory(bytes)
//    bytes, err = groupsHistory.Marshal()

package groups_history

import "encoding/json"

func UnmarshalGroupsHistory(data []byte) (GroupsHistory, error) {
	var r GroupsHistory
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GroupsHistory) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type GroupsHistory struct {
	Ok                  *bool             `json:"ok,omitempty"`                   
	Messages            []Message         `json:"messages,omitempty"`             
	HasMore             *bool             `json:"has_more,omitempty"`             
	ChannelActionsCount *int64            `json:"channel_actions_count,omitempty"`
	Warning             *string           `json:"warning,omitempty"`              
	ResponseMetadata    *ResponseMetadata `json:"response_metadata,omitempty"`    
	Error               *string           `json:"error,omitempty"`                
	Needed              *string           `json:"needed,omitempty"`               
	Provided            *string           `json:"provided,omitempty"`             
}

type Message struct {
	Type    *string `json:"type,omitempty"`   
	Subtype *string `json:"subtype,omitempty"`
	Ts      *string `json:"ts,omitempty"`     
	User    *string `json:"user,omitempty"`   
	Text    *string `json:"text,omitempty"`   
}

type ResponseMetadata struct {
	Messages []string `json:"messages,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
}
