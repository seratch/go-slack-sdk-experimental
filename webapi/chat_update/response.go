// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    chatUpdate, err := UnmarshalChatUpdate(bytes)
//    bytes, err = chatUpdate.Marshal()

package chat_update

import "encoding/json"

func UnmarshalChatUpdate(data []byte) (ChatUpdate, error) {
	var r ChatUpdate
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ChatUpdate) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ChatUpdate struct {
	Ok       *bool    `json:"ok,omitempty"`      
	Channel  *string  `json:"channel,omitempty"` 
	Ts       *string  `json:"ts,omitempty"`      
	Text     *string  `json:"text,omitempty"`    
	Message  *Message `json:"message,omitempty"` 
	Error    *string  `json:"error,omitempty"`   
	Needed   *string  `json:"needed,omitempty"`  
	Provided *string  `json:"provided,omitempty"`
}

type Message struct {
	BotID        *string     `json:"bot_id,omitempty"`        
	Type         *string     `json:"type,omitempty"`          
	Text         *string     `json:"text,omitempty"`          
	User         *string     `json:"user,omitempty"`          
	Team         *string     `json:"team,omitempty"`          
	BotProfile   *BotProfile `json:"bot_profile,omitempty"`   
	Blocks       []Block     `json:"blocks,omitempty"`        
	Edited       *Edited     `json:"edited,omitempty"`        
	Files        []File      `json:"files,omitempty"`         
	Upload       *bool       `json:"upload,omitempty"`        
	DisplayAsBot *bool       `json:"display_as_bot,omitempty"`
}

type Block struct {
	Type        *string    `json:"type,omitempty"`        
	BlockID     *string    `json:"block_id,omitempty"`    
	Text        *Text      `json:"text,omitempty"`        
	Accessory   *Accessory `json:"accessory,omitempty"`   
	Elements    []Element  `json:"elements,omitempty"`    
	Fallback    *string    `json:"fallback,omitempty"`    
	ImageURL    *string    `json:"image_url,omitempty"`   
	ImageWidth  *int64     `json:"image_width,omitempty"` 
	ImageHeight *int64     `json:"image_height,omitempty"`
	ImageBytes  *int64     `json:"image_bytes,omitempty"` 
	AltText     *string    `json:"alt_text,omitempty"`    
	Title       *Text      `json:"title,omitempty"`       
	Fields      []Text     `json:"fields,omitempty"`      
}

type Accessory struct {
	Fallback    *string `json:"fallback,omitempty"`    
	ImageURL    *string `json:"image_url,omitempty"`   
	ImageWidth  *int64  `json:"image_width,omitempty"` 
	ImageHeight *int64  `json:"image_height,omitempty"`
	ImageBytes  *int64  `json:"image_bytes,omitempty"` 
	Type        *string `json:"type,omitempty"`        
	AltText     *string `json:"alt_text,omitempty"`    
}

type Element struct {
	Type                         *string        `json:"type,omitempty"`                           
	ActionID                     *string        `json:"action_id,omitempty"`                      
	Text                         *Text          `json:"text,omitempty"`                           
	Value                        *string        `json:"value,omitempty"`                          
	URL                          *string        `json:"url,omitempty"`                            
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

type Edited struct {
	User *string `json:"user,omitempty"`
	Ts   *string `json:"ts,omitempty"`  
}

type File struct {
	ID                 *string `json:"id,omitempty"`                  
	Created            *int64  `json:"created,omitempty"`             
	Timestamp          *int64  `json:"timestamp,omitempty"`           
	Name               *string `json:"name,omitempty"`                
	Title              *string `json:"title,omitempty"`               
	Mimetype           *string `json:"mimetype,omitempty"`            
	Filetype           *string `json:"filetype,omitempty"`            
	PrettyType         *string `json:"pretty_type,omitempty"`         
	User               *string `json:"user,omitempty"`                
	Editable           *bool   `json:"editable,omitempty"`            
	Size               *int64  `json:"size,omitempty"`                
	Mode               *string `json:"mode,omitempty"`                
	IsExternal         *bool   `json:"is_external,omitempty"`         
	ExternalType       *string `json:"external_type,omitempty"`       
	IsPublic           *bool   `json:"is_public,omitempty"`           
	PublicURLShared    *bool   `json:"public_url_shared,omitempty"`   
	DisplayAsBot       *bool   `json:"display_as_bot,omitempty"`      
	Username           *string `json:"username,omitempty"`            
	URLPrivate         *string `json:"url_private,omitempty"`         
	URLPrivateDownload *string `json:"url_private_download,omitempty"`
	Permalink          *string `json:"permalink,omitempty"`           
	PermalinkPublic    *string `json:"permalink_public,omitempty"`    
	EditLink           *string `json:"edit_link,omitempty"`           
	Preview            *string `json:"preview,omitempty"`             
	PreviewHighlight   *string `json:"preview_highlight,omitempty"`   
	Lines              *int64  `json:"lines,omitempty"`               
	LinesMore          *int64  `json:"lines_more,omitempty"`          
	PreviewIsTruncated *bool   `json:"preview_is_truncated,omitempty"`
	IsStarred          *bool   `json:"is_starred,omitempty"`          
	HasRichPreview     *bool   `json:"has_rich_preview,omitempty"`    
}
