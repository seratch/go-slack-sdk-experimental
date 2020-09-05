// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    starsList, err := UnmarshalStarsList(bytes)
//    bytes, err = starsList.Marshal()

package stars_list

import "encoding/json"

func UnmarshalStarsList(data []byte) (StarsList, error) {
	var r StarsList
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *StarsList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type StarsList struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Items    []Item  `json:"items,omitempty"`   
	Paging   *Paging `json:"paging,omitempty"`  
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}

type Item struct {
	Type       *string  `json:"type,omitempty"`       
	Channel    *string  `json:"channel,omitempty"`    
	DateCreate *int64   `json:"date_create,omitempty"`
	Message    *Message `json:"message,omitempty"`    
	File       *File    `json:"file,omitempty"`       
	Comment    *Comment `json:"comment,omitempty"`    
}

type Comment struct {
	ID        *string `json:"id,omitempty"`        
	Created   *int64  `json:"created,omitempty"`   
	Timestamp *int64  `json:"timestamp,omitempty"` 
	User      *string `json:"user,omitempty"`      
	IsIntro   *bool   `json:"is_intro,omitempty"`  
	Comment   *string `json:"comment,omitempty"`   
	NumStars  *int64  `json:"num_stars,omitempty"` 
	IsStarred *bool   `json:"is_starred,omitempty"`
}

type File struct {
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
	ImageExifRotation  *int64   `json:"image_exif_rotation,omitempty"` 
	OriginalW          *int64   `json:"original_w,omitempty"`          
	OriginalH          *int64   `json:"original_h,omitempty"`          
	Permalink          *string  `json:"permalink,omitempty"`           
	CommentsCount      *int64   `json:"comments_count,omitempty"`      
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
	InitialComment     *Comment `json:"initial_comment,omitempty"`     
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

type Message struct {
	BotID           *string      `json:"bot_id,omitempty"`           
	Type            *string      `json:"type,omitempty"`             
	Text            *string      `json:"text,omitempty"`             
	User            *string      `json:"user,omitempty"`             
	Ts              *string      `json:"ts,omitempty"`               
	Team            *string      `json:"team,omitempty"`             
	Attachments     []Attachment `json:"attachments,omitempty"`      
	IsStarred       *bool        `json:"is_starred,omitempty"`       
	Permalink       *string      `json:"permalink,omitempty"`        
	Subtype         *string      `json:"subtype,omitempty"`          
	Username        *string      `json:"username,omitempty"`         
	Blocks          []Block      `json:"blocks,omitempty"`           
	ClientMsgID     *string      `json:"client_msg_id,omitempty"`    
	ThreadTs        *string      `json:"thread_ts,omitempty"`        
	ReplyCount      *int64       `json:"reply_count,omitempty"`      
	ReplyUsersCount *int64       `json:"reply_users_count,omitempty"`
	LatestReply     *string      `json:"latest_reply,omitempty"`     
	ReplyUsers      []string     `json:"reply_users,omitempty"`      
	Subscribed      *bool        `json:"subscribed,omitempty"`       
	LastRead        *string      `json:"last_read,omitempty"`        
	Reactions       []Reaction   `json:"reactions,omitempty"`        
}

type Attachment struct {
	MsgSubtype         *string   `json:"msg_subtype,omitempty"`          
	Fallback           *string   `json:"fallback,omitempty"`             
	CallbackID         *string   `json:"callback_id,omitempty"`          
	Color              *string   `json:"color,omitempty"`                
	Pretext            *string   `json:"pretext,omitempty"`              
	ServiceURL         *string   `json:"service_url,omitempty"`          
	ServiceName        *string   `json:"service_name,omitempty"`         
	ServiceIcon        *string   `json:"service_icon,omitempty"`         
	AuthorID           *string   `json:"author_id,omitempty"`            
	AuthorName         *string   `json:"author_name,omitempty"`          
	AuthorLink         *string   `json:"author_link,omitempty"`          
	AuthorIcon         *string   `json:"author_icon,omitempty"`          
	FromURL            *string   `json:"from_url,omitempty"`             
	OriginalURL        *string   `json:"original_url,omitempty"`         
	AuthorSubname      *string   `json:"author_subname,omitempty"`       
	ChannelID          *string   `json:"channel_id,omitempty"`           
	ChannelName        *string   `json:"channel_name,omitempty"`         
	ID                 *int64    `json:"id,omitempty"`                   
	BotID              *string   `json:"bot_id,omitempty"`               
	Indent             *bool     `json:"indent,omitempty"`               
	IsMsgUnfurl        *bool     `json:"is_msg_unfurl,omitempty"`        
	IsReplyUnfurl      *bool     `json:"is_reply_unfurl,omitempty"`      
	IsThreadRootUnfurl *bool     `json:"is_thread_root_unfurl,omitempty"`
	IsAppUnfurl        *bool     `json:"is_app_unfurl,omitempty"`        
	AppUnfurlURL       *string   `json:"app_unfurl_url,omitempty"`       
	Title              *string   `json:"title,omitempty"`                
	TitleLink          *string   `json:"title_link,omitempty"`           
	Text               *string   `json:"text,omitempty"`                 
	Fields             []Field   `json:"fields,omitempty"`               
	ImageURL           *string   `json:"image_url,omitempty"`            
	ImageWidth         *int64    `json:"image_width,omitempty"`          
	ImageHeight        *int64    `json:"image_height,omitempty"`         
	ImageBytes         *int64    `json:"image_bytes,omitempty"`          
	ThumbURL           *string   `json:"thumb_url,omitempty"`            
	ThumbWidth         *int64    `json:"thumb_width,omitempty"`          
	ThumbHeight        *int64    `json:"thumb_height,omitempty"`         
	VideoHTML          *string   `json:"video_html,omitempty"`           
	VideoHTMLWidth     *int64    `json:"video_html_width,omitempty"`     
	VideoHTMLHeight    *int64    `json:"video_html_height,omitempty"`    
	Footer             *string   `json:"footer,omitempty"`               
	FooterIcon         *string   `json:"footer_icon,omitempty"`          
	Ts                 *string   `json:"ts,omitempty"`                   
	MrkdwnIn           []string  `json:"mrkdwn_in,omitempty"`            
	Actions            []Action  `json:"actions,omitempty"`              
	Filename           *string   `json:"filename,omitempty"`             
	Size               *int64    `json:"size,omitempty"`                 
	Mimetype           *string   `json:"mimetype,omitempty"`             
	URL                *string   `json:"url,omitempty"`                  
	Metadata           *Metadata `json:"metadata,omitempty"`             
}

type Action struct {
	ID              *string        `json:"id,omitempty"`              
	Name            *string        `json:"name,omitempty"`            
	Text            *string        `json:"text,omitempty"`            
	Style           *string        `json:"style,omitempty"`           
	Type            *string        `json:"type,omitempty"`            
	Value           *string        `json:"value,omitempty"`           
	Confirm         *ActionConfirm `json:"confirm,omitempty"`         
	Options         []Option       `json:"options,omitempty"`         
	SelectedOptions []Option       `json:"selected_options,omitempty"`
	DataSource      *string        `json:"data_source,omitempty"`     
	MinQueryLength  *int64         `json:"min_query_length,omitempty"`
	OptionGroups    []OptionGroup  `json:"option_groups,omitempty"`   
	URL             *string        `json:"url,omitempty"`             
}

type ActionConfirm struct {
	Title       *string `json:"title,omitempty"`       
	Text        *string `json:"text,omitempty"`        
	OkText      *string `json:"ok_text,omitempty"`     
	DismissText *string `json:"dismiss_text,omitempty"`
}

type OptionGroup struct {
	Text *string `json:"text,omitempty"`
}

type Option struct {
	Text  *string `json:"text,omitempty"` 
	Value *string `json:"value,omitempty"`
}

type Field struct {
	Title *string `json:"title,omitempty"`
	Value *string `json:"value,omitempty"`
	Short *bool   `json:"short,omitempty"`
}

type Metadata struct {
	Thumb64    *bool   `json:"thumb_64,omitempty"`   
	Thumb80    *bool   `json:"thumb_80,omitempty"`   
	Thumb160   *bool   `json:"thumb_160,omitempty"`  
	OriginalW  *int64  `json:"original_w,omitempty"` 
	OriginalH  *int64  `json:"original_h,omitempty"` 
	Thumb360_W *int64  `json:"thumb_360_w,omitempty"`
	Thumb360_H *int64  `json:"thumb_360_h,omitempty"`
	Format     *string `json:"format,omitempty"`     
	Extension  *string `json:"extension,omitempty"`  
	Rotation   *int64  `json:"rotation,omitempty"`   
	ThumbTiny  *string `json:"thumb_tiny,omitempty"` 
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
	Type                         *string         `json:"type,omitempty"`                           
	Text                         *Text           `json:"text,omitempty"`                           
	ActionID                     *string         `json:"action_id,omitempty"`                      
	URL                          *string         `json:"url,omitempty"`                            
	Value                        *string         `json:"value,omitempty"`                          
	Style                        *string         `json:"style,omitempty"`                          
	Confirm                      *ElementConfirm `json:"confirm,omitempty"`                        
	Placeholder                  *Text           `json:"placeholder,omitempty"`                    
	InitialChannel               *string         `json:"initial_channel,omitempty"`                
	ResponseURLEnabled           *bool           `json:"response_url_enabled,omitempty"`           
	InitialConversation          *string         `json:"initial_conversation,omitempty"`           
	DefaultToCurrentConversation *bool           `json:"default_to_current_conversation,omitempty"`
	Filter                       *Filter         `json:"filter,omitempty"`                         
	InitialDate                  *string         `json:"initial_date,omitempty"`                   
	InitialOption                *InitialOption  `json:"initial_option,omitempty"`                 
	MinQueryLength               *int64          `json:"min_query_length,omitempty"`               
	ImageURL                     *string         `json:"image_url,omitempty"`                      
	AltText                      *string         `json:"alt_text,omitempty"`                       
	Fallback                     *string         `json:"fallback,omitempty"`                       
	ImageWidth                   *int64          `json:"image_width,omitempty"`                    
	ImageHeight                  *int64          `json:"image_height,omitempty"`                   
	ImageBytes                   *int64          `json:"image_bytes,omitempty"`                    
	InitialUser                  *string         `json:"initial_user,omitempty"`                   
}

type ElementConfirm struct {
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

type Reaction struct {
	Name  *string  `json:"name,omitempty"` 
	Users []string `json:"users,omitempty"`
	Count *int64   `json:"count,omitempty"`
}

type Paging struct {
	PerPage *int64 `json:"per_page,omitempty"`
	Spill   *int64 `json:"spill,omitempty"`   
	Page    *int64 `json:"page,omitempty"`    
	Total   *int64 `json:"total,omitempty"`   
	Pages   *int64 `json:"pages,omitempty"`   
}
