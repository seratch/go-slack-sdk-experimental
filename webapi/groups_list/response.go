// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    groupsList, err := UnmarshalGroupsList(bytes)
//    bytes, err = groupsList.Marshal()

package groups_list

import "encoding/json"

func UnmarshalGroupsList(data []byte) (GroupsList, error) {
	var r GroupsList
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GroupsList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type GroupsList struct {
	Ok               *bool             `json:"ok,omitempty"`               
	Groups           []Group           `json:"groups,omitempty"`           
	Warning          *string           `json:"warning,omitempty"`          
	ResponseMetadata *ResponseMetadata `json:"response_metadata,omitempty"`
	Error            *string           `json:"error,omitempty"`            
	Needed           *string           `json:"needed,omitempty"`           
	Provided         *string           `json:"provided,omitempty"`         
}

type Group struct {
	ID             *string  `json:"id,omitempty"`             
	Name           *string  `json:"name,omitempty"`           
	IsGroup        *bool    `json:"is_group,omitempty"`       
	Created        *int64   `json:"created,omitempty"`        
	Creator        *string  `json:"creator,omitempty"`        
	IsArchived     *bool    `json:"is_archived,omitempty"`    
	NameNormalized *string  `json:"name_normalized,omitempty"`
	IsMpim         *bool    `json:"is_mpim,omitempty"`        
	Members        []string `json:"members,omitempty"`        
	Topic          *Purpose `json:"topic,omitempty"`          
	Purpose        *Purpose `json:"purpose,omitempty"`        
	Priority       *int64   `json:"priority,omitempty"`       
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
