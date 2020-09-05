// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    chatDelete, err := UnmarshalChatDelete(bytes)
//    bytes, err = chatDelete.Marshal()

package chat_delete

import "encoding/json"

func UnmarshalChatDelete(data []byte) (ChatDelete, error) {
	var r ChatDelete
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ChatDelete) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ChatDelete struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Channel  *string `json:"channel,omitempty"` 
	Ts       *string `json:"ts,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
