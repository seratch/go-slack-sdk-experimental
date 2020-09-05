// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminEmojiRemove, err := UnmarshalAdminEmojiRemove(bytes)
//    bytes, err = adminEmojiRemove.Marshal()

package admin_emoji_remove

import "encoding/json"

func UnmarshalAdminEmojiRemove(data []byte) (AdminEmojiRemove, error) {
	var r AdminEmojiRemove
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminEmojiRemove) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminEmojiRemove struct {
	Ok               *bool             `json:"ok,omitempty"`               
	Error            *string           `json:"error,omitempty"`            
	ResponseMetadata *ResponseMetadata `json:"response_metadata,omitempty"`
	Needed           *string           `json:"needed,omitempty"`           
	Provided         *string           `json:"provided,omitempty"`         
}

type ResponseMetadata struct {
	Messages []string `json:"messages,omitempty"`
}
