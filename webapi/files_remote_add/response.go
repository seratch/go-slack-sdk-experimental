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
}

type Shares struct {
}
