// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    filesList, err := UnmarshalFilesList(bytes)
//    bytes, err = filesList.Marshal()

package files_list

import "encoding/json"

func UnmarshalFilesList(data []byte) (FilesList, error) {
	var r FilesList
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FilesList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type FilesList struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Files    []File  `json:"files,omitempty"`   
	Paging   *Paging `json:"paging,omitempty"`  
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}

type File struct {
	ID                      *string         `json:"id,omitempty"`                       
	Created                 *int64          `json:"created,omitempty"`                  
	Timestamp               *int64          `json:"timestamp,omitempty"`                
	Name                    *string         `json:"name,omitempty"`                     
	Title                   *string         `json:"title,omitempty"`                    
	Mimetype                *string         `json:"mimetype,omitempty"`                 
	Filetype                *string         `json:"filetype,omitempty"`                 
	PrettyType              *string         `json:"pretty_type,omitempty"`              
	User                    *string         `json:"user,omitempty"`                     
	Editable                *bool           `json:"editable,omitempty"`                 
	Size                    *int64          `json:"size,omitempty"`                     
	Mode                    *string         `json:"mode,omitempty"`                     
	IsExternal              *bool           `json:"is_external,omitempty"`              
	ExternalType            *string         `json:"external_type,omitempty"`            
	IsPublic                *bool           `json:"is_public,omitempty"`                
	PublicURLShared         *bool           `json:"public_url_shared,omitempty"`        
	DisplayAsBot            *bool           `json:"display_as_bot,omitempty"`           
	Username                *string         `json:"username,omitempty"`                 
	URLPrivate              *string         `json:"url_private,omitempty"`              
	Permalink               *string         `json:"permalink,omitempty"`                
	Channels                []string        `json:"channels,omitempty"`                 
	Groups                  []string        `json:"groups,omitempty"`                   
	Ims                     []string        `json:"ims,omitempty"`                      
	CommentsCount           *int64          `json:"comments_count,omitempty"`           
	Thumb64                 *string         `json:"thumb_64,omitempty"`                 
	Thumb80                 *string         `json:"thumb_80,omitempty"`                 
	Thumb360                *string         `json:"thumb_360,omitempty"`                
	Thumb360_W              *int64          `json:"thumb_360_w,omitempty"`              
	Thumb360_H              *int64          `json:"thumb_360_h,omitempty"`              
	Thumb160                *string         `json:"thumb_160,omitempty"`                
	Thumb800                *string         `json:"thumb_800,omitempty"`                
	Thumb800_W              *int64          `json:"thumb_800_w,omitempty"`              
	Thumb800_H              *int64          `json:"thumb_800_h,omitempty"`              
	ImageExifRotation       *int64          `json:"image_exif_rotation,omitempty"`      
	OriginalW               *int64          `json:"original_w,omitempty"`               
	OriginalH               *int64          `json:"original_h,omitempty"`               
	ThumbTiny               *string         `json:"thumb_tiny,omitempty"`               
	URLPrivateDownload      *string         `json:"url_private_download,omitempty"`     
	Thumb480                *string         `json:"thumb_480,omitempty"`                
	Thumb480_W              *int64          `json:"thumb_480_w,omitempty"`              
	Thumb480_H              *int64          `json:"thumb_480_h,omitempty"`              
	Thumb720                *string         `json:"thumb_720,omitempty"`                
	Thumb720_W              *int64          `json:"thumb_720_w,omitempty"`              
	Thumb720_H              *int64          `json:"thumb_720_h,omitempty"`              
	PermalinkPublic         *string         `json:"permalink_public,omitempty"`         
	EditLink                *string         `json:"edit_link,omitempty"`                
	Preview                 *string         `json:"preview,omitempty"`                  
	PreviewHighlight        *string         `json:"preview_highlight,omitempty"`        
	Lines                   *int64          `json:"lines,omitempty"`                    
	LinesMore               *int64          `json:"lines_more,omitempty"`               
	PreviewIsTruncated      *bool           `json:"preview_is_truncated,omitempty"`     
	Thumb960                *string         `json:"thumb_960,omitempty"`                
	Thumb960_W              *int64          `json:"thumb_960_w,omitempty"`              
	Thumb960_H              *int64          `json:"thumb_960_h,omitempty"`              
	Thumb1024               *string         `json:"thumb_1024,omitempty"`               
	Thumb1024_W             *int64          `json:"thumb_1024_w,omitempty"`             
	Thumb1024_H             *int64          `json:"thumb_1024_h,omitempty"`             
	MediaDisplayType        *string         `json:"media_display_type,omitempty"`       
	NumStars                *int64          `json:"num_stars,omitempty"`                
	IsStarred               *bool           `json:"is_starred,omitempty"`               
	Thumb360_GIF            *string         `json:"thumb_360_gif,omitempty"`            
	Thumb480_GIF            *string         `json:"thumb_480_gif,omitempty"`            
	Deanimate               *string         `json:"deanimate,omitempty"`                
	DeanimateGIF            *string         `json:"deanimate_gif,omitempty"`            
	Subject                 *string         `json:"subject,omitempty"`                  
	To                      []Cc            `json:"to,omitempty"`                       
	From                    []Cc            `json:"from,omitempty"`                     
	Cc                      []Cc            `json:"cc,omitempty"`                       
	Attachments             []Attachment    `json:"attachments,omitempty"`              
	OriginalAttachmentCount *int64          `json:"original_attachment_count,omitempty"`
	PlainText               *string         `json:"plain_text,omitempty"`               
	PreviewPlainText        *string         `json:"preview_plain_text,omitempty"`       
	Headers                 *Headers        `json:"headers,omitempty"`                  
	AppID                   *string         `json:"app_id,omitempty"`                   
	AppName                 *string         `json:"app_name,omitempty"`                 
	HasMore                 *bool           `json:"has_more,omitempty"`                 
	SentToSelf              *bool           `json:"sent_to_self,omitempty"`             
	NonOwnerEditable        *bool           `json:"non_owner_editable,omitempty"`       
	Editor                  *string         `json:"editor,omitempty"`                   
	LastEditor              *string         `json:"last_editor,omitempty"`              
	Updated                 *int64          `json:"updated,omitempty"`                  
	ExternalID              *string         `json:"external_id,omitempty"`              
	ExternalURL             *string         `json:"external_url,omitempty"`             
	Thumb64_GIF             *string         `json:"thumb_64_gif,omitempty"`             
	Thumb64_W               *string         `json:"thumb_64_w,omitempty"`               
	Thumb64_H               *string         `json:"thumb_64_h,omitempty"`               
	Thumb80_GIF             *string         `json:"thumb_80_gif,omitempty"`             
	Thumb80_W               *string         `json:"thumb_80_w,omitempty"`               
	Thumb80_H               *string         `json:"thumb_80_h,omitempty"`               
	Thumb160_GIF            *string         `json:"thumb_160_gif,omitempty"`            
	Thumb160_W              *string         `json:"thumb_160_w,omitempty"`              
	Thumb160_H              *string         `json:"thumb_160_h,omitempty"`              
	Thumb720_GIF            *string         `json:"thumb_720_gif,omitempty"`            
	Thumb800_GIF            *string         `json:"thumb_800_gif,omitempty"`            
	Thumb960_GIF            *string         `json:"thumb_960_gif,omitempty"`            
	Thumb1024_GIF           *string         `json:"thumb_1024_gif,omitempty"`           
	ThumbVideo              *string         `json:"thumb_video,omitempty"`              
	ThumbGIF                *string         `json:"thumb_gif,omitempty"`                
	ThumbPDF                *string         `json:"thumb_pdf,omitempty"`                
	ThumbPDFW               *string         `json:"thumb_pdf_w,omitempty"`              
	ThumbPDFH               *string         `json:"thumb_pdf_h,omitempty"`              
	ConvertedPDF            *string         `json:"converted_pdf,omitempty"`            
	Pjpeg                   *string         `json:"pjpeg,omitempty"`                    
	HasRichPreview          *bool           `json:"has_rich_preview,omitempty"`         
	Shares                  *Shares         `json:"shares,omitempty"`                   
	ChannelActionsTs        *string         `json:"channel_actions_ts,omitempty"`       
	ChannelActionsCount     *int64          `json:"channel_actions_count,omitempty"`    
	SimplifiedHTML          *string         `json:"simplified_html,omitempty"`          
	BotID                   *string         `json:"bot_id,omitempty"`                   
	InitialComment          *InitialComment `json:"initial_comment,omitempty"`          
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
	VideoURL           *string   `json:"video_url,omitempty"`            
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
	ID              *string       `json:"id,omitempty"`              
	Name            *string       `json:"name,omitempty"`            
	Text            *string       `json:"text,omitempty"`            
	Style           *string       `json:"style,omitempty"`           
	Type            *string       `json:"type,omitempty"`            
	Value           *string       `json:"value,omitempty"`           
	Confirm         *Confirm      `json:"confirm,omitempty"`         
	Options         []Option      `json:"options,omitempty"`         
	SelectedOptions []Option      `json:"selected_options,omitempty"`
	DataSource      *string       `json:"data_source,omitempty"`     
	MinQueryLength  *int64        `json:"min_query_length,omitempty"`
	OptionGroups    []OptionGroup `json:"option_groups,omitempty"`   
	URL             *string       `json:"url,omitempty"`             
}

type Confirm struct {
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

type Cc struct {
	Address  *string `json:"address,omitempty"` 
	Name     *string `json:"name,omitempty"`    
	Original *string `json:"original,omitempty"`
}

type Headers struct {
	Date      *string `json:"date,omitempty"`       
	ReplyTo   *string `json:"reply_to,omitempty"`   
	MessageID *string `json:"message_id,omitempty"` 
	InReplyTo *string `json:"in_reply_to,omitempty"`
}

type InitialComment struct {
	ID        *string `json:"id,omitempty"`       
	Created   *int64  `json:"created,omitempty"`  
	Timestamp *int64  `json:"timestamp,omitempty"`
	User      *string `json:"user,omitempty"`     
	Comment   *string `json:"comment,omitempty"`  
	Channel   *string `json:"channel,omitempty"`  
	IsIntro   *bool   `json:"is_intro,omitempty"` 
}

type Shares struct {
}

type Paging struct {
	Count *int64 `json:"count,omitempty"`
	Total *int64 `json:"total,omitempty"`
	Page  *int64 `json:"page,omitempty"` 
	Pages *int64 `json:"pages,omitempty"`
}
