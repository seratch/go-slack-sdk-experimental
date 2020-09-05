// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    groupsCreate, err := UnmarshalGroupsCreate(bytes)
//    bytes, err = groupsCreate.Marshal()

package groups_create

import "encoding/json"

func UnmarshalGroupsCreate(data []byte) (GroupsCreate, error) {
	var r GroupsCreate
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GroupsCreate) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type GroupsCreate struct {
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
	UnreadCount        *int64   `json:"unread_count,omitempty"`        
	UnreadCountDisplay *int64   `json:"unread_count_display,omitempty"`
	Topic              *Purpose `json:"topic,omitempty"`               
	Purpose            *Purpose `json:"purpose,omitempty"`             
	Priority           *int64   `json:"priority,omitempty"`            
	Members            []string `json:"members,omitempty"`             
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
