// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    iMHistory, err := UnmarshalIMHistory(bytes)
//    bytes, err = iMHistory.Marshal()

package im_history

import "encoding/json"

func UnmarshalIMHistory(data []byte) (IMHistory, error) {
	var r IMHistory
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *IMHistory) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type IMHistory struct {
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
	BotID           *string     `json:"bot_id,omitempty"`           
	Type            *string     `json:"type,omitempty"`             
	Text            *string     `json:"text,omitempty"`             
	User            *string     `json:"user,omitempty"`             
	Ts              *string     `json:"ts,omitempty"`               
	Team            *string     `json:"team,omitempty"`             
	BotProfile      *BotProfile `json:"bot_profile,omitempty"`      
	ThreadTs        *string     `json:"thread_ts,omitempty"`        
	ParentUserID    *string     `json:"parent_user_id,omitempty"`   
	ReplyCount      *int64      `json:"reply_count,omitempty"`      
	ReplyUsersCount *int64      `json:"reply_users_count,omitempty"`
	LatestReply     *string     `json:"latest_reply,omitempty"`     
	ReplyUsers      []string    `json:"reply_users,omitempty"`      
	Subscribed      *bool       `json:"subscribed,omitempty"`       
	LastRead        *string     `json:"last_read,omitempty"`        
	ClientMsgID     *string     `json:"client_msg_id,omitempty"`    
	Blocks          []Block     `json:"blocks,omitempty"`           
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

type ResponseMetadata struct {
	Messages []string `json:"messages,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
}
