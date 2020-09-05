// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    chatScheduleMessage, err := UnmarshalChatScheduleMessage(bytes)
//    bytes, err = chatScheduleMessage.Marshal()

package chat_scheduleMessage

import "encoding/json"

func UnmarshalChatScheduleMessage(data []byte) (ChatScheduleMessage, error) {
	var r ChatScheduleMessage
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ChatScheduleMessage) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ChatScheduleMessage struct {
	Ok                 *bool    `json:"ok,omitempty"`                  
	ScheduledMessageID *string  `json:"scheduled_message_id,omitempty"`
	Channel            *string  `json:"channel,omitempty"`             
	PostAt             *int64   `json:"post_at,omitempty"`             
	Message            *Message `json:"message,omitempty"`             
	Error              *string  `json:"error,omitempty"`               
	Needed             *string  `json:"needed,omitempty"`              
	Provided           *string  `json:"provided,omitempty"`            
}

type Message struct {
	BotID      *string     `json:"bot_id,omitempty"`     
	Type       *string     `json:"type,omitempty"`       
	Text       *string     `json:"text,omitempty"`       
	User       *string     `json:"user,omitempty"`       
	Team       *string     `json:"team,omitempty"`       
	BotProfile *BotProfile `json:"bot_profile,omitempty"`
	Blocks     []Block     `json:"blocks,omitempty"`     
}

type Block struct {
	Type        *string    `json:"type,omitempty"`        
	BlockID     *string    `json:"block_id,omitempty"`    
	Text        *Text      `json:"text,omitempty"`        
	Elements    []Element  `json:"elements,omitempty"`    
	Fallback    *string    `json:"fallback,omitempty"`    
	ImageURL    *string    `json:"image_url,omitempty"`   
	ImageWidth  *int64     `json:"image_width,omitempty"` 
	ImageHeight *int64     `json:"image_height,omitempty"`
	ImageBytes  *int64     `json:"image_bytes,omitempty"` 
	AltText     *string    `json:"alt_text,omitempty"`    
	Title       *Text      `json:"title,omitempty"`       
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
