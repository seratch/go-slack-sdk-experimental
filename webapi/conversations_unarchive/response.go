// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    conversationsUnarchive, err := UnmarshalConversationsUnarchive(bytes)
//    bytes, err = conversationsUnarchive.Marshal()

package conversations_unarchive

import "encoding/json"

func UnmarshalConversationsUnarchive(data []byte) (ConversationsUnarchive, error) {
	var r ConversationsUnarchive
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationsUnarchive) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ConversationsUnarchive struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
