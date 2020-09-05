// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminEmojiAddAlias, err := UnmarshalAdminEmojiAddAlias(bytes)
//    bytes, err = adminEmojiAddAlias.Marshal()

package admin_emoji_addAlias

import "encoding/json"

func UnmarshalAdminEmojiAddAlias(data []byte) (AdminEmojiAddAlias, error) {
	var r AdminEmojiAddAlias
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminEmojiAddAlias) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminEmojiAddAlias struct {
	Ok               *bool             `json:"ok,omitempty"`               
	Error            *string           `json:"error,omitempty"`            
	ResponseMetadata *ResponseMetadata `json:"response_metadata,omitempty"`
	Needed           *string           `json:"needed,omitempty"`           
	Provided         *string           `json:"provided,omitempty"`         
}

type ResponseMetadata struct {
	Messages []string `json:"messages,omitempty"`
}
