// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    conversationsMark, err := UnmarshalConversationsMark(bytes)
//    bytes, err = conversationsMark.Marshal()

package conversations_mark

import "encoding/json"

func UnmarshalConversationsMark(data []byte) (ConversationsMark, error) {
	var r ConversationsMark
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationsMark) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ConversationsMark struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
