// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    mpimOpen, err := UnmarshalMpimOpen(bytes)
//    bytes, err = mpimOpen.Marshal()

package mpim_open

import "encoding/json"

func UnmarshalMpimOpen(data []byte) (MpimOpen, error) {
	var r MpimOpen
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MpimOpen) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MpimOpen struct {
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
	Priority           *int64   `json:"priority,omitempty"`            
}

type Latest struct {
	BotID        *string     `json:"bot_id,omitempty"`        
	Type         *string     `json:"type,omitempty"`          
	Text         *string     `json:"text,omitempty"`          
	User         *string     `json:"user,omitempty"`          
	Ts           *string     `json:"ts,omitempty"`            
	Team         *string     `json:"team,omitempty"`          
	BotProfile   *BotProfile `json:"bot_profile,omitempty"`   
	ThreadTs     *string     `json:"thread_ts,omitempty"`     
	ParentUserID *string     `json:"parent_user_id,omitempty"`
}

type BotProfile struct {
	ID      *string `json:"id,omitempty"`     
	Deleted *bool   `json:"deleted,omitempty"`
	Name    *string `json:"name,omitempty"`   
	Updated *int64  `json:"updated,omitempty"`
	AppID   *string `json:"app_id,omitempty"` 
	Icons   *Icons  `json:"icons,omitempty"`  
	TeamID  *string `json:"team_id,omitempty"`
}

type Icons struct {
	Image36 *string `json:"image_36,omitempty"`
	Image48 *string `json:"image_48,omitempty"`
	Image72 *string `json:"image_72,omitempty"`
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
