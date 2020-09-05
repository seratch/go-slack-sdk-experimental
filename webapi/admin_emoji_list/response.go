// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminEmojiList, err := UnmarshalAdminEmojiList(bytes)
//    bytes, err = adminEmojiList.Marshal()

package admin_emoji_list

import "encoding/json"

func UnmarshalAdminEmojiList(data []byte) (AdminEmojiList, error) {
	var r AdminEmojiList
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminEmojiList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminEmojiList struct {
	Ok               *bool                `json:"ok,omitempty"`               
	Emoji            *AdminEmojiListEmoji `json:"emoji,omitempty"`            
	ResponseMetadata *ResponseMetadata    `json:"response_metadata,omitempty"`
	Error            *string              `json:"error,omitempty"`            
	Needed           *string              `json:"needed,omitempty"`           
	Provided         *string              `json:"provided,omitempty"`         
}

type AdminEmojiListEmoji struct {
	Emoji      *EmojiClass `json:"emoji,omitempty"` 
	EmojiEmoji *EmojiClass `json:"emoji_,omitempty"`
}

type EmojiClass struct {
	Emoji      *string `json:"emoji,omitempty"` 
	EmojiEmoji *string `json:"emoji_,omitempty"`
}

type ResponseMetadata struct {
	NextCursor *string `json:"next_cursor,omitempty"`
}
