// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    groupsInfo, err := UnmarshalGroupsInfo(bytes)
//    bytes, err = groupsInfo.Marshal()

package groups_info

import "encoding/json"

func UnmarshalGroupsInfo(data []byte) (GroupsInfo, error) {
	var r GroupsInfo
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GroupsInfo) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type GroupsInfo struct {
	Ok               *bool             `json:"ok,omitempty"`               
	Group            *Group            `json:"group,omitempty"`            
	Warning          *string           `json:"warning,omitempty"`          
	ResponseMetadata *ResponseMetadata `json:"response_metadata,omitempty"`
	Error            *string           `json:"error,omitempty"`            
	Needed           *string           `json:"needed,omitempty"`           
	Provided         *string           `json:"provided,omitempty"`         
}

type Group struct {
	ID                 *string  `json:"id,omitempty"`                  
	Name               *string  `json:"name,omitempty"`                
	IsGroup            *bool    `json:"is_group,omitempty"`            
	Created            *int64   `json:"created,omitempty"`             
	Creator            *string  `json:"creator,omitempty"`             
	IsArchived         *bool    `json:"is_archived,omitempty"`         
	NameNormalized     *string  `json:"name_normalized,omitempty"`     
	IsMpim             *bool    `json:"is_mpim,omitempty"`             
	IsOpen             *bool    `json:"is_open,omitempty"`             
	LastRead           *string  `json:"last_read,omitempty"`           
	Latest             *Latest  `json:"latest,omitempty"`              
	UnreadCount        *int64   `json:"unread_count,omitempty"`        
	UnreadCountDisplay *int64   `json:"unread_count_display,omitempty"`
	Members            []string `json:"members,omitempty"`             
	Topic              *Purpose `json:"topic,omitempty"`               
	Purpose            *Purpose `json:"purpose,omitempty"`             
}

type Latest struct {
	Type    *string `json:"type,omitempty"`   
	Subtype *string `json:"subtype,omitempty"`
	Ts      *string `json:"ts,omitempty"`     
	User    *string `json:"user,omitempty"`   
	Text    *string `json:"text,omitempty"`   
}

type Purpose struct {
	Value   *string `json:"value,omitempty"`   
	Creator *string `json:"creator,omitempty"` 
	LastSet *int64  `json:"last_set,omitempty"`
}

type ResponseMetadata struct {
	Messages []string `json:"messages,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
}
