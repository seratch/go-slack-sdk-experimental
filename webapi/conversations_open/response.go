// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    conversationsOpen, err := UnmarshalConversationsOpen(bytes)
//    bytes, err = conversationsOpen.Marshal()

package conversations_open

import "encoding/json"

func UnmarshalConversationsOpen(data []byte) (ConversationsOpen, error) {
	var r ConversationsOpen
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationsOpen) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ConversationsOpen struct {
	Ok          *bool    `json:"ok,omitempty"`          
	Channel     *Channel `json:"channel,omitempty"`     
	NoOp        *bool    `json:"no_op,omitempty"`       
	AlreadyOpen *bool    `json:"already_open,omitempty"`
	Error       *string  `json:"error,omitempty"`       
	Needed      *string  `json:"needed,omitempty"`      
	Provided    *string  `json:"provided,omitempty"`    
}

type Channel struct {
	ID                 *string `json:"id,omitempty"`                  
	Created            *int64  `json:"created,omitempty"`             
	IsArchived         *bool   `json:"is_archived,omitempty"`         
	IsIM               *bool   `json:"is_im,omitempty"`               
	IsOrgShared        *bool   `json:"is_org_shared,omitempty"`       
	User               *string `json:"user,omitempty"`                
	LastRead           *string `json:"last_read,omitempty"`           
	Latest             *Latest `json:"latest,omitempty"`              
	UnreadCount        *int64  `json:"unread_count,omitempty"`        
	UnreadCountDisplay *int64  `json:"unread_count_display,omitempty"`
	IsOpen             *bool   `json:"is_open,omitempty"`             
	Priority           *int64  `json:"priority,omitempty"`            
}

type Latest struct {
	Type    *string `json:"type,omitempty"`   
	Subtype *string `json:"subtype,omitempty"`
	Text    *string `json:"text,omitempty"`   
	Ts      *string `json:"ts,omitempty"`     
	BotID   *string `json:"bot_id,omitempty"` 
	Blocks  []Block `json:"blocks,omitempty"` 
}

type Block struct {
	Type        *string    `json:"type,omitempty"`        
	Elements    []Element  `json:"elements,omitempty"`    
	BlockID     *string    `json:"block_id,omitempty"`    
	Fallback    *string    `json:"fallback,omitempty"`    
	ImageURL    *string    `json:"image_url,omitempty"`   
	ImageWidth  *int64     `json:"image_width,omitempty"` 
	ImageHeight *int64     `json:"image_height,omitempty"`
	ImageBytes  *int64     `json:"image_bytes,omitempty"` 
	AltText     *string    `json:"alt_text,omitempty"`    
	Title       *Text      `json:"title,omitempty"`       
	Text        *Text      `json:"text,omitempty"`        
	Fields      []Text     `json:"fields,omitempty"`      
	Accessory   *Accessory `json:"accessory,omitempty"`   
}

type Accessory struct {
	Type        *string `json:"type,omitempty"`        
	ImageURL    *string `json:"image_url,omitempty"`   
	AltText     *string `json:"alt_text,omitempty"`    
	Fallback    *string `json:"fallback,omitempty"`    
	ImageWidth  *int64  `json:"image_width,omitempty"` 
	ImageHeight *int64  `json:"image_height,omitempty"`
	ImageBytes  *int64  `json:"image_bytes,omitempty"` 
}

type Element struct {
	Type                         *string        `json:"type,omitempty"`                           
	Text                         *Text          `json:"text,omitempty"`                           
	ActionID                     *string        `json:"action_id,omitempty"`                      
	URL                          *string        `json:"url,omitempty"`                            
	Value                        *string        `json:"value,omitempty"`                          
	Style                        *string        `json:"style,omitempty"`                          
	Confirm                      *Confirm       `json:"confirm,omitempty"`                        
	Placeholder                  *Text          `json:"placeholder,omitempty"`                    
	InitialChannel               *string        `json:"initial_channel,omitempty"`                
	ResponseURLEnabled           *bool          `json:"response_url_enabled,omitempty"`           
	InitialConversation          *string        `json:"initial_conversation,omitempty"`           
	DefaultToCurrentConversation *bool          `json:"default_to_current_conversation,omitempty"`
	Filter                       *Filter        `json:"filter,omitempty"`                         
	InitialDate                  *string        `json:"initial_date,omitempty"`                   
	InitialOption                *InitialOption `json:"initial_option,omitempty"`                 
	MinQueryLength               *int64         `json:"min_query_length,omitempty"`               
	ImageURL                     *string        `json:"image_url,omitempty"`                      
	AltText                      *string        `json:"alt_text,omitempty"`                       
	Fallback                     *string        `json:"fallback,omitempty"`                       
	ImageWidth                   *int64         `json:"image_width,omitempty"`                    
	ImageHeight                  *int64         `json:"image_height,omitempty"`                   
	ImageBytes                   *int64         `json:"image_bytes,omitempty"`                    
	InitialUser                  *string        `json:"initial_user,omitempty"`                   
}

type Confirm struct {
	Title   *Text   `json:"title,omitempty"`  
	Text    *Text   `json:"text,omitempty"`   
	Confirm *Text   `json:"confirm,omitempty"`
	Deny    *Text   `json:"deny,omitempty"`   
	Style   *string `json:"style,omitempty"`  
}

type Text struct {
	Type     *string `json:"type,omitempty"`    
	Text     *string `json:"text,omitempty"`    
	Emoji    *bool   `json:"emoji,omitempty"`   
	Verbatim *bool   `json:"verbatim,omitempty"`
}

type Filter struct {
	ExcludeExternalSharedChannels *bool `json:"exclude_external_shared_channels,omitempty"`
	ExcludeBotUsers               *bool `json:"exclude_bot_users,omitempty"`               
}

type InitialOption struct {
	Text        *Text   `json:"text,omitempty"`       
	Value       *string `json:"value,omitempty"`      
	Description *Text   `json:"description,omitempty"`
	URL         *string `json:"url,omitempty"`        
}
