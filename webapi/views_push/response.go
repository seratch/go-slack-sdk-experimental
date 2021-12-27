// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    viewsPush, err := UnmarshalViewsPush(bytes)
//    bytes, err = viewsPush.Marshal()

package views_push

import "encoding/json"

func UnmarshalViewsPush(data []byte) (ViewsPush, error) {
	var r ViewsPush
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ViewsPush) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ViewsPush struct {
	Ok               *bool             `json:"ok,omitempty"`               
	Warning          *string           `json:"warning,omitempty"`          
	Error            *string           `json:"error,omitempty"`            
	Needed           *string           `json:"needed,omitempty"`           
	Provided         *string           `json:"provided,omitempty"`         
	View             *View             `json:"view,omitempty"`             
	ResponseMetadata *ResponseMetadata `json:"response_metadata,omitempty"`
}

type ResponseMetadata struct {
	Messages []string `json:"messages,omitempty"`
}

type View struct {
	ID                 *string `json:"id,omitempty"`                   
	TeamID             *string `json:"team_id,omitempty"`              
	Type               *string `json:"type,omitempty"`                 
	Title              *Close  `json:"title,omitempty"`                
	Submit             *Close  `json:"submit,omitempty"`               
	Close              *Close  `json:"close,omitempty"`                
	Blocks             []Block `json:"blocks,omitempty"`               
	PrivateMetadata    *string `json:"private_metadata,omitempty"`     
	CallbackID         *string `json:"callback_id,omitempty"`          
	ExternalID         *string `json:"external_id,omitempty"`          
	State              *State  `json:"state,omitempty"`                
	Hash               *string `json:"hash,omitempty"`                 
	ClearOnClose       *bool   `json:"clear_on_close,omitempty"`       
	NotifyOnClose      *bool   `json:"notify_on_close,omitempty"`      
	SubmitDisabled     *bool   `json:"submit_disabled,omitempty"`      
	RootViewID         *string `json:"root_view_id,omitempty"`         
	PreviousViewID     *string `json:"previous_view_id,omitempty"`     
	AppID              *string `json:"app_id,omitempty"`               
	AppInstalledTeamID *string `json:"app_installed_team_id,omitempty"`
	BotID              *string `json:"bot_id,omitempty"`               
}

type Block struct {
	Type           *string    `json:"type,omitempty"`           
	BlockID        *string    `json:"block_id,omitempty"`       
	Label          *Close     `json:"label,omitempty"`          
	Element        *Element   `json:"element,omitempty"`        
	DispatchAction *bool      `json:"dispatch_action,omitempty"`
	Hint           *Close     `json:"hint,omitempty"`           
	Optional       *bool      `json:"optional,omitempty"`       
	Elements       []Element  `json:"elements,omitempty"`       
	Fallback       *string    `json:"fallback,omitempty"`       
	ImageURL       *string    `json:"image_url,omitempty"`      
	ImageWidth     *int64     `json:"image_width,omitempty"`    
	ImageHeight    *int64     `json:"image_height,omitempty"`   
	ImageBytes     *int64     `json:"image_bytes,omitempty"`    
	AltText        *string    `json:"alt_text,omitempty"`       
	Title          *Close     `json:"title,omitempty"`          
	Text           *Close     `json:"text,omitempty"`           
	Fields         []Close    `json:"fields,omitempty"`         
	Accessory      *Accessory `json:"accessory,omitempty"`      
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
	Type                         *string               `json:"type,omitempty"`                           
	ActionID                     *string               `json:"action_id,omitempty"`                      
	Placeholder                  *Close                `json:"placeholder,omitempty"`                    
	InitialValue                 *string               `json:"initial_value,omitempty"`                  
	Multiline                    *bool                 `json:"multiline,omitempty"`                      
	MinLength                    *int64                `json:"min_length,omitempty"`                     
	MaxLength                    *int64                `json:"max_length,omitempty"`                     
	DispatchActionConfig         *DispatchActionConfig `json:"dispatch_action_config,omitempty"`         
	FocusOnLoad                  *bool                 `json:"focus_on_load,omitempty"`                  
	Options                      []Option              `json:"options,omitempty"`                        
	InitialOption                *Option               `json:"initial_option,omitempty"`                 
	Confirm                      *Confirm              `json:"confirm,omitempty"`                        
	Text                         *Close                `json:"text,omitempty"`                           
	URL                          *string               `json:"url,omitempty"`                            
	Value                        *string               `json:"value,omitempty"`                          
	Style                        *string               `json:"style,omitempty"`                          
	InitialChannel               *string               `json:"initial_channel,omitempty"`                
	ResponseURLEnabled           *bool                 `json:"response_url_enabled,omitempty"`           
	InitialConversation          *string               `json:"initial_conversation,omitempty"`           
	DefaultToCurrentConversation *bool                 `json:"default_to_current_conversation,omitempty"`
	Filter                       *Filter               `json:"filter,omitempty"`                         
	InitialDate                  *string               `json:"initial_date,omitempty"`                   
	InitialTime                  *string               `json:"initial_time,omitempty"`                   
	MinQueryLength               *int64                `json:"min_query_length,omitempty"`               
	ImageURL                     *string               `json:"image_url,omitempty"`                      
	AltText                      *string               `json:"alt_text,omitempty"`                       
	Fallback                     *string               `json:"fallback,omitempty"`                       
	ImageWidth                   *int64                `json:"image_width,omitempty"`                    
	ImageHeight                  *int64                `json:"image_height,omitempty"`                   
	ImageBytes                   *int64                `json:"image_bytes,omitempty"`                    
	InitialUser                  *string               `json:"initial_user,omitempty"`                   
}

type Confirm struct {
	Title   *Close  `json:"title,omitempty"`  
	Text    *Close  `json:"text,omitempty"`   
	Confirm *Close  `json:"confirm,omitempty"`
	Deny    *Close  `json:"deny,omitempty"`   
	Style   *string `json:"style,omitempty"`  
}

type Close struct {
	Type     *Type   `json:"type,omitempty"`    
	Text     *string `json:"text,omitempty"`    
	Emoji    *bool   `json:"emoji,omitempty"`   
	Verbatim *bool   `json:"verbatim,omitempty"`
}

type DispatchActionConfig struct {
	TriggerActionsOn []string `json:"trigger_actions_on,omitempty"`
}

type Filter struct {
	ExcludeExternalSharedChannels *bool `json:"exclude_external_shared_channels,omitempty"`
	ExcludeBotUsers               *bool `json:"exclude_bot_users,omitempty"`               
}

type Option struct {
	Text        *Close  `json:"text,omitempty"`       
	Value       *string `json:"value,omitempty"`      
	Description *Close  `json:"description,omitempty"`
	URL         *string `json:"url,omitempty"`        
}

type State struct {
}

type Type string
const (
	Empty Type = ""
	Mrkdwn Type = "mrkdwn"
	PlainText Type = "plain_text"
)
