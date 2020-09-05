// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    filesRemoteRemove, err := UnmarshalFilesRemoteRemove(bytes)
//    bytes, err = filesRemoteRemove.Marshal()

package files_remote_remove

import "encoding/json"

func UnmarshalFilesRemoteRemove(data []byte) (FilesRemoteRemove, error) {
	var r FilesRemoteRemove
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FilesRemoteRemove) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type FilesRemoteRemove struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
