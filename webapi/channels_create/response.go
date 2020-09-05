// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    channelsCreate, err := UnmarshalChannelsCreate(bytes)
//    bytes, err = channelsCreate.Marshal()

package channels_create

import "encoding/json"

func UnmarshalChannelsCreate(data []byte) (ChannelsCreate, error) {
	var r ChannelsCreate
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ChannelsCreate) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ChannelsCreate struct {
	Ok       *bool    `json:"ok,omitempty"`      
	Channel  *Channel `json:"channel,omitempty"` 
	Error    *string  `json:"error,omitempty"`   
	Needed   *string  `json:"needed,omitempty"`  
	Provided *string  `json:"provided,omitempty"`
}

type Channel struct {
	ID                 *string  `json:"id,omitempty"`                  
	Name               *string  `json:"name,omitempty"`                
	IsChannel          *bool    `json:"is_channel,omitempty"`          
	Created            *int64   `json:"created,omitempty"`             
	IsArchived         *bool    `json:"is_archived,omitempty"`         
	IsGeneral          *bool    `json:"is_general,omitempty"`          
	Unlinked           *int64   `json:"unlinked,omitempty"`            
	Creator            *string  `json:"creator,omitempty"`             
	NameNormalized     *string  `json:"name_normalized,omitempty"`     
	IsShared           *bool    `json:"is_shared,omitempty"`           
	IsOrgShared        *bool    `json:"is_org_shared,omitempty"`       
	IsMember           *bool    `json:"is_member,omitempty"`           
	IsPrivate          *bool    `json:"is_private,omitempty"`          
	IsMpim             *bool    `json:"is_mpim,omitempty"`             
	LastRead           *string  `json:"last_read,omitempty"`           
	UnreadCount        *int64   `json:"unread_count,omitempty"`        
	UnreadCountDisplay *int64   `json:"unread_count_display,omitempty"`
	Members            []string `json:"members,omitempty"`             
	Topic              *Purpose `json:"topic,omitempty"`               
	Purpose            *Purpose `json:"purpose,omitempty"`             
	PreviousNames      []string `json:"previous_names,omitempty"`      
	Priority           *int64   `json:"priority,omitempty"`            
}

type Purpose struct {
	Value   *string `json:"value,omitempty"`   
	Creator *string `json:"creator,omitempty"` 
	LastSet *int64  `json:"last_set,omitempty"`
}
