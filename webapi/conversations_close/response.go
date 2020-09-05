// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    conversationsClose, err := UnmarshalConversationsClose(bytes)
//    bytes, err = conversationsClose.Marshal()

package conversations_close

import "encoding/json"

func UnmarshalConversationsClose(data []byte) (ConversationsClose, error) {
	var r ConversationsClose
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationsClose) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ConversationsClose struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
