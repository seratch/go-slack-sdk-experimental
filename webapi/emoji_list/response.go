// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    emojiList, err := UnmarshalEmojiList(bytes)
//    bytes, err = emojiList.Marshal()

package emoji_list

import "encoding/json"

func UnmarshalEmojiList(data []byte) (EmojiList, error) {
	var r EmojiList
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *EmojiList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type EmojiList struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Emoji    *Emoji  `json:"emoji,omitempty"`   
	CacheTs  *string `json:"cache_ts,omitempty"`
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}

type Emoji struct {
	Emoji      *string `json:"emoji,omitempty"` 
	EmojiEmoji *string `json:"emoji_,omitempty"`
}
