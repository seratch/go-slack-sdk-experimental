// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    conversationsKick, err := UnmarshalConversationsKick(bytes)
//    bytes, err = conversationsKick.Marshal()

package conversations_kick

import "encoding/json"

func UnmarshalConversationsKick(data []byte) (ConversationsKick, error) {
	var r ConversationsKick
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationsKick) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ConversationsKick struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
