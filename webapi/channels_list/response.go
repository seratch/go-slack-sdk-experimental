// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    channelsList, err := UnmarshalChannelsList(bytes)
//    bytes, err = channelsList.Marshal()

package channels_list

import "encoding/json"

func UnmarshalChannelsList(data []byte) (ChannelsList, error) {
	var r ChannelsList
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ChannelsList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ChannelsList struct {
	Ok               *bool             `json:"ok,omitempty"`               
	Channels         []Channel         `json:"channels,omitempty"`         
	Warning          *string           `json:"warning,omitempty"`          
	ResponseMetadata *ResponseMetadata `json:"response_metadata,omitempty"`
	Error            *string           `json:"error,omitempty"`            
	Needed           *string           `json:"needed,omitempty"`           
	Provided         *string           `json:"provided,omitempty"`         
}

type Channel struct {
	ID             *string  `json:"id,omitempty"`             
	Name           *string  `json:"name,omitempty"`           
	IsChannel      *bool    `json:"is_channel,omitempty"`     
	Created        *int64   `json:"created,omitempty"`        
	IsArchived     *bool    `json:"is_archived,omitempty"`    
	IsGeneral      *bool    `json:"is_general,omitempty"`     
	Unlinked       *int64   `json:"unlinked,omitempty"`       
	Creator        *string  `json:"creator,omitempty"`        
	NameNormalized *string  `json:"name_normalized,omitempty"`
	IsShared       *bool    `json:"is_shared,omitempty"`      
	IsOrgShared    *bool    `json:"is_org_shared,omitempty"`  
	IsMember       *bool    `json:"is_member,omitempty"`      
	IsPrivate      *bool    `json:"is_private,omitempty"`     
	IsMpim         *bool    `json:"is_mpim,omitempty"`        
	Members        []string `json:"members,omitempty"`        
	Topic          *Purpose `json:"topic,omitempty"`          
	Purpose        *Purpose `json:"purpose,omitempty"`        
	PreviousNames  []string `json:"previous_names,omitempty"` 
	NumMembers     *int64   `json:"num_members,omitempty"`    
}

type Purpose struct {
	Value   *string `json:"value,omitempty"`   
	Creator *string `json:"creator,omitempty"` 
	LastSet *int64  `json:"last_set,omitempty"`
}

type ResponseMetadata struct {
	NextCursor *string  `json:"next_cursor,omitempty"`
	Messages   []string `json:"messages,omitempty"`   
	Warnings   []string `json:"warnings,omitempty"`   
}
