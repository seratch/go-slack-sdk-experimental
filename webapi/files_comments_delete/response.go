// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    filesCommentsDelete, err := UnmarshalFilesCommentsDelete(bytes)
//    bytes, err = filesCommentsDelete.Marshal()

package files_comments_delete

import "encoding/json"

func UnmarshalFilesCommentsDelete(data []byte) (FilesCommentsDelete, error) {
	var r FilesCommentsDelete
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FilesCommentsDelete) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type FilesCommentsDelete struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
