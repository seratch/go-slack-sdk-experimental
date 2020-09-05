// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    searchMessages, err := UnmarshalSearchMessages(bytes)
//    bytes, err = searchMessages.Marshal()

package search_messages

import "encoding/json"

func UnmarshalSearchMessages(data []byte) (SearchMessages, error) {
	var r SearchMessages
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SearchMessages) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SearchMessages struct {
	Ok       *bool     `json:"ok,omitempty"`      
	Query    *string   `json:"query,omitempty"`   
	Messages *Messages `json:"messages,omitempty"`
	Error    *string   `json:"error,omitempty"`   
	Needed   *string   `json:"needed,omitempty"`  
	Provided *string   `json:"provided,omitempty"`
}

type Messages struct {
	Total      *int64      `json:"total,omitempty"`     
	Pagination *Pagination `json:"pagination,omitempty"`
	Paging     *Paging     `json:"paging,omitempty"`    
	Matches    []Match     `json:"matches,omitempty"`   
}

type Match struct {
	Iid         *string   `json:"iid,omitempty"`         
	Team        *string   `json:"team,omitempty"`        
	Channel     *Channel  `json:"channel,omitempty"`     
	Type        *string   `json:"type,omitempty"`        
	User        *string   `json:"user,omitempty"`        
	Username    *string   `json:"username,omitempty"`    
	Ts          *string   `json:"ts,omitempty"`          
	Text        *string   `json:"text,omitempty"`        
	Permalink   *string   `json:"permalink,omitempty"`   
	NoReactions *bool     `json:"no_reactions,omitempty"`
	Previous    *Previous `json:"previous,omitempty"`    
	Previous2   *Previous `json:"previous_2,omitempty"`  
	Blocks      []Block   `json:"blocks,omitempty"`      
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

type Channel struct {
	ID                 *string  `json:"id,omitempty"`                   
	IsChannel          *bool    `json:"is_channel,omitempty"`           
	IsGroup            *bool    `json:"is_group,omitempty"`             
	IsIM               *bool    `json:"is_im,omitempty"`                
	Name               *string  `json:"name,omitempty"`                 
	IsShared           *bool    `json:"is_shared,omitempty"`            
	IsOrgShared        *bool    `json:"is_org_shared,omitempty"`        
	IsEXTShared        *bool    `json:"is_ext_shared,omitempty"`        
	IsPrivate          *bool    `json:"is_private,omitempty"`           
	IsMpim             *bool    `json:"is_mpim,omitempty"`              
	PendingShared      []string `json:"pending_shared,omitempty"`       
	IsPendingEXTShared *bool    `json:"is_pending_ext_shared,omitempty"`
	User               *string  `json:"user,omitempty"`                 
}

type Previous struct {
	Type      *string `json:"type,omitempty"`     
	User      *string `json:"user,omitempty"`     
	Username  *string `json:"username,omitempty"` 
	Ts        *string `json:"ts,omitempty"`       
	Text      *string `json:"text,omitempty"`     
	Iid       *string `json:"iid,omitempty"`      
	Permalink *string `json:"permalink,omitempty"`
}

type Pagination struct {
	TotalCount *int64 `json:"total_count,omitempty"`
	Page       *int64 `json:"page,omitempty"`       
	PerPage    *int64 `json:"per_page,omitempty"`   
	PageCount  *int64 `json:"page_count,omitempty"` 
	First      *int64 `json:"first,omitempty"`      
	Last       *int64 `json:"last,omitempty"`       
}

type Paging struct {
	Count *int64 `json:"count,omitempty"`
	Total *int64 `json:"total,omitempty"`
	Page  *int64 `json:"page,omitempty"` 
	Pages *int64 `json:"pages,omitempty"`
}
