// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    filesCommentsEdit, err := UnmarshalFilesCommentsEdit(bytes)
//    bytes, err = filesCommentsEdit.Marshal()

package files_comments_edit

import "encoding/json"

func UnmarshalFilesCommentsEdit(data []byte) (FilesCommentsEdit, error) {
	var r FilesCommentsEdit
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FilesCommentsEdit) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type FilesCommentsEdit struct {
	Ok       *bool    `json:"ok,omitempty"`      
	Comment  *Comment `json:"comment,omitempty"` 
	Error    *string  `json:"error,omitempty"`   
	Needed   *string  `json:"needed,omitempty"`  
	Provided *string  `json:"provided,omitempty"`
}

type Comment struct {
	ID        *string `json:"id,omitempty"`       
	Created   *int64  `json:"created,omitempty"`  
	Timestamp *int64  `json:"timestamp,omitempty"`
	User      *string `json:"user,omitempty"`     
	IsIntro   *bool   `json:"is_intro,omitempty"` 
	Comment   *string `json:"comment,omitempty"`  
}
