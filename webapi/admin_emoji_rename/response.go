// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminEmojiRename, err := UnmarshalAdminEmojiRename(bytes)
//    bytes, err = adminEmojiRename.Marshal()

package admin_emoji_rename

import "encoding/json"

func UnmarshalAdminEmojiRename(data []byte) (AdminEmojiRename, error) {
	var r AdminEmojiRename
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminEmojiRename) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminEmojiRename struct {
	Ok               *bool             `json:"ok,omitempty"`               
	Error            *string           `json:"error,omitempty"`            
	ResponseMetadata *ResponseMetadata `json:"response_metadata,omitempty"`
	Needed           *string           `json:"needed,omitempty"`           
	Provided         *string           `json:"provided,omitempty"`         
}

type ResponseMetadata struct {
	Messages []string `json:"messages,omitempty"`
}
