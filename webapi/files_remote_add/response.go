// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    filesRemoteAdd, err := UnmarshalFilesRemoteAdd(bytes)
//    bytes, err = filesRemoteAdd.Marshal()

package files_remote_add

import "encoding/json"

func UnmarshalFilesRemoteAdd(data []byte) (FilesRemoteAdd, error) {
	var r FilesRemoteAdd
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FilesRemoteAdd) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type FilesRemoteAdd struct {
	Ok       *bool   `json:"ok,omitempty"`      
	File     *File   `json:"file,omitempty"`    
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
	Subject                 *string         `json:"subject,omitempty"`                  
	Mimetype                *string         `json:"mimetype,omitempty"`                 
	Filetype                *string         `json:"filetype,omitempty"`                 
	PrettyType              *string         `json:"pretty_type,omitempty"`              
	User                    *string         `json:"user,omitempty"`                     
	Mode                    *string         `json:"mode,omitempty"`                     
	Editable                *bool           `json:"editable,omitempty"`                 
	NonOwnerEditable        *bool           `json:"non_owner_editable,omitempty"`       
	Editor                  *string         `json:"editor,omitempty"`                   
	LastEditor              *string         `json:"last_editor,omitempty"`              
	Updated                 *int64          `json:"updated,omitempty"`                  
	OriginalAttachmentCount *int64          `json:"original_attachment_count,omitempty"`
	IsExternal              *bool           `json:"is_external,omitempty"`              
	ExternalType            *string         `json:"external_type,omitempty"`            
	ExternalID              *string         `json:"external_id,omitempty"`              
	ExternalURL             *string         `json:"external_url,omitempty"`             
	Username                *string         `json:"username,omitempty"`                 
	Size                    *int64          `json:"size,omitempty"`                     
	URLPrivate              *string         `json:"url_private,omitempty"`              
	URLPrivateDownload      *string         `json:"url_private_download,omitempty"`     
	AppID                   *string         `json:"app_id,omitempty"`                   
	AppName                 *string         `json:"app_name,omitempty"`                 
	Thumb64                 *string         `json:"thumb_64,omitempty"`                 
	Thumb64_GIF             *string         `json:"thumb_64_gif,omitempty"`             
	Thumb64_W               *string         `json:"thumb_64_w,omitempty"`               
	Thumb64_H               *string         `json:"thumb_64_h,omitempty"`               
	Thumb80                 *string         `json:"thumb_80,omitempty"`                 
	Thumb80_GIF             *string         `json:"thumb_80_gif,omitempty"`             
	Thumb80_W               *string         `json:"thumb_80_w,omitempty"`               
	Thumb80_H               *string         `json:"thumb_80_h,omitempty"`               
	Thumb160                *string         `json:"thumb_160,omitempty"`                
	Thumb160_GIF            *string         `json:"thumb_160_gif,omitempty"`            
	Thumb160_W              *string         `json:"thumb_160_w,omitempty"`              
	Thumb160_H              *string         `json:"thumb_160_h,omitempty"`              
	Thumb360                *string         `json:"thumb_360,omitempty"`                
	Thumb360_GIF            *string         `json:"thumb_360_gif,omitempty"`            
	Thumb360_W              *string         `json:"thumb_360_w,omitempty"`              
	Thumb360_H              *string         `json:"thumb_360_h,omitempty"`              
	Thumb480                *string         `json:"thumb_480,omitempty"`                
	Thumb480_GIF            *string         `json:"thumb_480_gif,omitempty"`            
	Thumb480_W              *string         `json:"thumb_480_w,omitempty"`              
	Thumb480_H              *string         `json:"thumb_480_h,omitempty"`              
	Thumb720                *string         `json:"thumb_720,omitempty"`                
	Thumb720_GIF            *string         `json:"thumb_720_gif,omitempty"`            
	Thumb720_W              *string         `json:"thumb_720_w,omitempty"`              
	Thumb720_H              *string         `json:"thumb_720_h,omitempty"`              
	Thumb800                *string         `json:"thumb_800,omitempty"`                
	Thumb800_GIF            *string         `json:"thumb_800_gif,omitempty"`            
	Thumb800_W              *string         `json:"thumb_800_w,omitempty"`              
	Thumb800_H              *string         `json:"thumb_800_h,omitempty"`              
	Thumb960                *string         `json:"thumb_960,omitempty"`                
	Thumb960_GIF            *string         `json:"thumb_960_gif,omitempty"`            
	Thumb960_W              *string         `json:"thumb_960_w,omitempty"`              
	Thumb960_H              *string         `json:"thumb_960_h,omitempty"`              
	Thumb1024               *string         `json:"thumb_1024,omitempty"`               
	Thumb1024_GIF           *string         `json:"thumb_1024_gif,omitempty"`           
	Thumb1024_W             *string         `json:"thumb_1024_w,omitempty"`             
	Thumb1024_H             *string         `json:"thumb_1024_h,omitempty"`             
	ThumbVideo              *string         `json:"thumb_video,omitempty"`              
	ThumbGIF                *string         `json:"thumb_gif,omitempty"`                
	ThumbPDF                *string         `json:"thumb_pdf,omitempty"`                
	ThumbPDFW               *string         `json:"thumb_pdf_w,omitempty"`              
	ThumbPDFH               *string         `json:"thumb_pdf_h,omitempty"`              
	ThumbTiny               *string         `json:"thumb_tiny,omitempty"`               
	ConvertedPDF            *string         `json:"converted_pdf,omitempty"`            
	ImageExifRotation       *int64          `json:"image_exif_rotation,omitempty"`      
	OriginalW               *string         `json:"original_w,omitempty"`               
	OriginalH               *string         `json:"original_h,omitempty"`               
	Deanimate               *string         `json:"deanimate,omitempty"`                
	DeanimateGIF            *string         `json:"deanimate_gif,omitempty"`            
	Pjpeg                   *string         `json:"pjpeg,omitempty"`                    
	Permalink               *string         `json:"permalink,omitempty"`                
	PermalinkPublic         *string         `json:"permalink_public,omitempty"`         
	EditLink                *string         `json:"edit_link,omitempty"`                
	HasRichPreview          *bool           `json:"has_rich_preview,omitempty"`         
	MediaDisplayType        *string         `json:"media_display_type,omitempty"`       
	PreviewIsTruncated      *bool           `json:"preview_is_truncated,omitempty"`     
	Preview                 *string         `json:"preview,omitempty"`                  
	PreviewHighlight        *string         `json:"preview_highlight,omitempty"`        
	PlainText               *string         `json:"plain_text,omitempty"`               
	PreviewPlainText        *string         `json:"preview_plain_text,omitempty"`       
	HasMore                 *bool           `json:"has_more,omitempty"`                 
	SentToSelf              *bool           `json:"sent_to_self,omitempty"`             
	Lines                   *int64          `json:"lines,omitempty"`                    
	LinesMore               *int64          `json:"lines_more,omitempty"`               
	IsPublic                *bool           `json:"is_public,omitempty"`                
	PublicURLShared         *bool           `json:"public_url_shared,omitempty"`        
	DisplayAsBot            *bool           `json:"display_as_bot,omitempty"`           
	Shares                  *Shares         `json:"shares,omitempty"`                   
	ChannelActionsTs        *string         `json:"channel_actions_ts,omitempty"`       
	ChannelActionsCount     *int64          `json:"channel_actions_count,omitempty"`    
	Headers                 *Headers        `json:"headers,omitempty"`                  
	SimplifiedHTML          *string         `json:"simplified_html,omitempty"`          
	BotID                   *string         `json:"bot_id,omitempty"`                   
	InitialComment          *InitialComment `json:"initial_comment,omitempty"`          
	NumStars                *int64          `json:"num_stars,omitempty"`                
	IsStarred               *bool           `json:"is_starred,omitempty"`               
	CommentsCount           *int64          `json:"comments_count,omitempty"`           
}

type Headers struct {
	Date      *string `json:"date,omitempty"`       
	InReplyTo *string `json:"in_reply_to,omitempty"`
	ReplyTo   *string `json:"reply_to,omitempty"`   
	MessageID *string `json:"message_id,omitempty"` 
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
