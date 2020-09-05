// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    conversationsLeave, err := UnmarshalConversationsLeave(bytes)
//    bytes, err = conversationsLeave.Marshal()

package conversations_leave

import "encoding/json"

func UnmarshalConversationsLeave(data []byte) (ConversationsLeave, error) {
	var r ConversationsLeave
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationsLeave) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ConversationsLeave struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
