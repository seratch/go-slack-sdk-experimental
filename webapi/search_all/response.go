// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    searchAll, err := UnmarshalSearchAll(bytes)
//    bytes, err = searchAll.Marshal()

package search_all

import "encoding/json"

func UnmarshalSearchAll(data []byte) (SearchAll, error) {
	var r SearchAll
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SearchAll) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SearchAll struct {
	Ok       *bool     `json:"ok,omitempty"`      
	Error    *string   `json:"error,omitempty"`   
	Query    *string   `json:"query,omitempty"`   
	Messages *Messages `json:"messages,omitempty"`
	Files    *Files    `json:"files,omitempty"`   
	Posts    *Posts    `json:"posts,omitempty"`   
	Needed   *string   `json:"needed,omitempty"`  
	Provided *string   `json:"provided,omitempty"`
}

type Files struct {
	Total      *int64       `json:"total,omitempty"`     
	Pagination *Pagination  `json:"pagination,omitempty"`
	Paging     *Paging      `json:"paging,omitempty"`    
	Matches    []FilesMatch `json:"matches,omitempty"`   
}

type FilesMatch struct {
	ID                 *string  `json:"id,omitempty"`                  
	Created            *int64   `json:"created,omitempty"`             
	Timestamp          *int64   `json:"timestamp,omitempty"`           
	Name               *string  `json:"name,omitempty"`                
	Title              *string  `json:"title,omitempty"`               
	Mimetype           *string  `json:"mimetype,omitempty"`            
	Filetype           *string  `json:"filetype,omitempty"`            
	PrettyType         *string  `json:"pretty_type,omitempty"`         
	User               *string  `json:"user,omitempty"`                
	Editable           *bool    `json:"editable,omitempty"`            
	Size               *int64   `json:"size,omitempty"`                
	Mode               *string  `json:"mode,omitempty"`                
	IsExternal         *bool    `json:"is_external,omitempty"`         
	ExternalType       *string  `json:"external_type,omitempty"`       
	IsPublic           *bool    `json:"is_public,omitempty"`           
	PublicURLShared    *bool    `json:"public_url_shared,omitempty"`   
	DisplayAsBot       *bool    `json:"display_as_bot,omitempty"`      
	Username           *string  `json:"username,omitempty"`            
	URLPrivate         *string  `json:"url_private,omitempty"`         
	Thumb64            *string  `json:"thumb_64,omitempty"`            
	Thumb80            *string  `json:"thumb_80,omitempty"`            
	Thumb360           *string  `json:"thumb_360,omitempty"`           
	Thumb360_W         *int64   `json:"thumb_360_w,omitempty"`         
	Thumb360_H         *int64   `json:"thumb_360_h,omitempty"`         
	Thumb480           *string  `json:"thumb_480,omitempty"`           
	Thumb480_W         *int64   `json:"thumb_480_w,omitempty"`         
	Thumb480_H         *int64   `json:"thumb_480_h,omitempty"`         
	Thumb160           *string  `json:"thumb_160,omitempty"`           
	Thumb720           *string  `json:"thumb_720,omitempty"`           
	Thumb720_W         *int64   `json:"thumb_720_w,omitempty"`         
	Thumb720_H         *int64   `json:"thumb_720_h,omitempty"`         
	Thumb800           *string  `json:"thumb_800,omitempty"`           
	Thumb800_W         *int64   `json:"thumb_800_w,omitempty"`         
	Thumb800_H         *int64   `json:"thumb_800_h,omitempty"`         
	Thumb960           *string  `json:"thumb_960,omitempty"`           
	Thumb960_W         *int64   `json:"thumb_960_w,omitempty"`         
	Thumb960_H         *int64   `json:"thumb_960_h,omitempty"`         
	Thumb1024          *string  `json:"thumb_1024,omitempty"`          
	Thumb1024_W        *int64   `json:"thumb_1024_w,omitempty"`        
	Thumb1024_H        *int64   `json:"thumb_1024_h,omitempty"`        
	OriginalW          *int64   `json:"original_w,omitempty"`          
	OriginalH          *int64   `json:"original_h,omitempty"`          
	ThumbTiny          *string  `json:"thumb_tiny,omitempty"`          
	Permalink          *string  `json:"permalink,omitempty"`           
	IsStarred          *bool    `json:"is_starred,omitempty"`          
	Shares             *Shares  `json:"shares,omitempty"`              
	Channels           []string `json:"channels,omitempty"`            
	Groups             []string `json:"groups,omitempty"`              
	Ims                []string `json:"ims,omitempty"`                 
	ExternalID         *string  `json:"external_id,omitempty"`         
	ExternalURL        *string  `json:"external_url,omitempty"`        
	HasRichPreview     *bool    `json:"has_rich_preview,omitempty"`    
	URLPrivateDownload *string  `json:"url_private_download,omitempty"`
	PermalinkPublic    *string  `json:"permalink_public,omitempty"`    
	EditLink           *string  `json:"edit_link,omitempty"`           
	Preview            *string  `json:"preview,omitempty"`             
	PreviewHighlight   *string  `json:"preview_highlight,omitempty"`   
	Lines              *int64   `json:"lines,omitempty"`               
	LinesMore          *int64   `json:"lines_more,omitempty"`          
	PreviewIsTruncated *bool    `json:"preview_is_truncated,omitempty"`
	ImageExifRotation  *int64   `json:"image_exif_rotation,omitempty"` 
}

type Shares struct {
	Public map[string][]Public `json:"public,omitempty"`
}

type Public struct {
	ReplyUsers      []string `json:"reply_users,omitempty"`      
	ReplyUsersCount *int64   `json:"reply_users_count,omitempty"`
	ReplyCount      *int64   `json:"reply_count,omitempty"`      
	Ts              *string  `json:"ts,omitempty"`               
	ChannelName     *string  `json:"channel_name,omitempty"`     
	TeamID          *string  `json:"team_id,omitempty"`          
	ShareUserID     *string  `json:"share_user_id,omitempty"`    
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

type Messages struct {
	Total      *int64          `json:"total,omitempty"`     
	Pagination *Pagination     `json:"pagination,omitempty"`
	Paging     *Paging         `json:"paging,omitempty"`    
	Matches    []MessagesMatch `json:"matches,omitempty"`   
}

type MessagesMatch struct {
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

type Posts struct {
	Total   *int64   `json:"total,omitempty"`  
	Matches []string `json:"matches,omitempty"`
}
