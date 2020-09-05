// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    filesRemoteInfo, err := UnmarshalFilesRemoteInfo(bytes)
//    bytes, err = filesRemoteInfo.Marshal()

package files_remote_info

import "encoding/json"

func UnmarshalFilesRemoteInfo(data []byte) (FilesRemoteInfo, error) {
	var r FilesRemoteInfo
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FilesRemoteInfo) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type FilesRemoteInfo struct {
	Ok       *bool   `json:"ok,omitempty"`      
	File     *File   `json:"file,omitempty"`    
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}

type File struct {
	ID              *string  `json:"id,omitempty"`               
	Created         *int64   `json:"created,omitempty"`          
	Timestamp       *int64   `json:"timestamp,omitempty"`        
	Name            *string  `json:"name,omitempty"`             
	Title           *string  `json:"title,omitempty"`            
	Mimetype        *string  `json:"mimetype,omitempty"`         
	Filetype        *string  `json:"filetype,omitempty"`         
	PrettyType      *string  `json:"pretty_type,omitempty"`      
	User            *string  `json:"user,omitempty"`             
	Editable        *bool    `json:"editable,omitempty"`         
	Size            *int64   `json:"size,omitempty"`             
	Mode            *string  `json:"mode,omitempty"`             
	IsExternal      *bool    `json:"is_external,omitempty"`      
	ExternalType    *string  `json:"external_type,omitempty"`    
	IsPublic        *bool    `json:"is_public,omitempty"`        
	PublicURLShared *bool    `json:"public_url_shared,omitempty"`
	DisplayAsBot    *bool    `json:"display_as_bot,omitempty"`   
	Username        *string  `json:"username,omitempty"`         
	URLPrivate      *string  `json:"url_private,omitempty"`      
	Permalink       *string  `json:"permalink,omitempty"`        
	CommentsCount   *int64   `json:"comments_count,omitempty"`   
	IsStarred       *bool    `json:"is_starred,omitempty"`       
	Shares          *Shares  `json:"shares,omitempty"`           
	Channels        []string `json:"channels,omitempty"`         
	Groups          []string `json:"groups,omitempty"`           
	Ims             []string `json:"ims,omitempty"`              
	ExternalID      *string  `json:"external_id,omitempty"`      
	ExternalURL     *string  `json:"external_url,omitempty"`     
	HasRichPreview  *bool    `json:"has_rich_preview,omitempty"` 
	Thumb64         *string  `json:"thumb_64,omitempty"`         
	Thumb80         *string  `json:"thumb_80,omitempty"`         
	Thumb360        *string  `json:"thumb_360,omitempty"`        
	Thumb360_W      *int64   `json:"thumb_360_w,omitempty"`      
	Thumb360_H      *int64   `json:"thumb_360_h,omitempty"`      
	Thumb480        *string  `json:"thumb_480,omitempty"`        
	Thumb480_W      *int64   `json:"thumb_480_w,omitempty"`      
	Thumb480_H      *int64   `json:"thumb_480_h,omitempty"`      
	Thumb160        *string  `json:"thumb_160,omitempty"`        
	Thumb720        *string  `json:"thumb_720,omitempty"`        
	Thumb720_W      *int64   `json:"thumb_720_w,omitempty"`      
	Thumb720_H      *int64   `json:"thumb_720_h,omitempty"`      
	Thumb800        *string  `json:"thumb_800,omitempty"`        
	Thumb800_W      *int64   `json:"thumb_800_w,omitempty"`      
	Thumb800_H      *int64   `json:"thumb_800_h,omitempty"`      
	Thumb960        *string  `json:"thumb_960,omitempty"`        
	Thumb960_W      *int64   `json:"thumb_960_w,omitempty"`      
	Thumb960_H      *int64   `json:"thumb_960_h,omitempty"`      
	Thumb1024       *string  `json:"thumb_1024,omitempty"`       
	Thumb1024_W     *int64   `json:"thumb_1024_w,omitempty"`     
	Thumb1024_H     *int64   `json:"thumb_1024_h,omitempty"`     
	OriginalW       *int64   `json:"original_w,omitempty"`       
	OriginalH       *int64   `json:"original_h,omitempty"`       
	ThumbTiny       *string  `json:"thumb_tiny,omitempty"`       
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
