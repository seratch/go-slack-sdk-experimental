// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    filesCommentsAdd, err := UnmarshalFilesCommentsAdd(bytes)
//    bytes, err = filesCommentsAdd.Marshal()

package files_comments_add

import "encoding/json"

func UnmarshalFilesCommentsAdd(data []byte) (FilesCommentsAdd, error) {
	var r FilesCommentsAdd
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FilesCommentsAdd) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type FilesCommentsAdd struct {
	Ok        *bool   `json:"ok,omitempty"`        
	Error     *string `json:"error,omitempty"`     
	ReqMethod *string `json:"req_method,omitempty"`
	Needed    *string `json:"needed,omitempty"`    
	Provided  *string `json:"provided,omitempty"`  
}
