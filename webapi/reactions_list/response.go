// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    reactionsList, err := UnmarshalReactionsList(bytes)
//    bytes, err = reactionsList.Marshal()

package reactions_list

import "encoding/json"

func UnmarshalReactionsList(data []byte) (ReactionsList, error) {
	var r ReactionsList
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ReactionsList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ReactionsList struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Items    []Item  `json:"items,omitempty"`   
	Paging   *Paging `json:"paging,omitempty"`  
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}

type Item struct {
	Type    *string  `json:"type,omitempty"`   
	Channel *string  `json:"channel,omitempty"`
	Message *Message `json:"message,omitempty"`
}

type Message struct {
	Type            *string     `json:"type,omitempty"`             
	Text            *string     `json:"text,omitempty"`             
	Files           []File      `json:"files,omitempty"`            
	Upload          *bool       `json:"upload,omitempty"`           
	User            *string     `json:"user,omitempty"`             
	DisplayAsBot    *bool       `json:"display_as_bot,omitempty"`   
	Ts              *string     `json:"ts,omitempty"`               
	Reactions       []Reaction  `json:"reactions,omitempty"`        
	Permalink       *string     `json:"permalink,omitempty"`        
	ClientMsgID     *string     `json:"client_msg_id,omitempty"`    
	Team            *string     `json:"team,omitempty"`             
	Blocks          []Block     `json:"blocks,omitempty"`           
	BotID           *string     `json:"bot_id,omitempty"`           
	BotProfile      *BotProfile `json:"bot_profile,omitempty"`      
	ThreadTs        *string     `json:"thread_ts,omitempty"`        
	ReplyCount      *int64      `json:"reply_count,omitempty"`      
	ReplyUsersCount *int64      `json:"reply_users_count,omitempty"`
	LatestReply     *string     `json:"latest_reply,omitempty"`     
	ReplyUsers      []string    `json:"reply_users,omitempty"`      
	Subscribed      *bool       `json:"subscribed,omitempty"`       
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
	Preview            *string `json:"preview,omitempty"`             
	LastEditor         *string `json:"last_editor,omitempty"`         
	NonOwnerEditable   *bool   `json:"non_owner_editable,omitempty"`  
	Updated            *int64  `json:"updated,omitempty"`             
	IsStarred          *bool   `json:"is_starred,omitempty"`          
	HasRichPreview     *bool   `json:"has_rich_preview,omitempty"`    
}

type Reaction struct {
	Name  *string  `json:"name,omitempty"` 
	Users []string `json:"users,omitempty"`
	Count *int64   `json:"count,omitempty"`
}

type Paging struct {
	Count *int64 `json:"count,omitempty"`
	Total *int64 `json:"total,omitempty"`
	Page  *int64 `json:"page,omitempty"` 
	Pages *int64 `json:"pages,omitempty"`
}
