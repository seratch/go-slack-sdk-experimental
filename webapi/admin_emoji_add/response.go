// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminEmojiAdd, err := UnmarshalAdminEmojiAdd(bytes)
//    bytes, err = adminEmojiAdd.Marshal()

package admin_emoji_add

import "encoding/json"

func UnmarshalAdminEmojiAdd(data []byte) (AdminEmojiAdd, error) {
	var r AdminEmojiAdd
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminEmojiAdd) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminEmojiAdd struct {
	Ok               *bool             `json:"ok,omitempty"`               
	Error            *string           `json:"error,omitempty"`            
	ResponseMetadata *ResponseMetadata `json:"response_metadata,omitempty"`
	Needed           *string           `json:"needed,omitempty"`           
	Provided         *string           `json:"provided,omitempty"`         
}

type ResponseMetadata struct {
	Messages []string `json:"messages,omitempty"`
}
