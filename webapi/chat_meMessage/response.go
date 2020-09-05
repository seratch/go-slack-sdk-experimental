// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    chatMeMessage, err := UnmarshalChatMeMessage(bytes)
//    bytes, err = chatMeMessage.Marshal()

package chat_meMessage

import "encoding/json"

func UnmarshalChatMeMessage(data []byte) (ChatMeMessage, error) {
	var r ChatMeMessage
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ChatMeMessage) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ChatMeMessage struct {
	Channel  *string `json:"channel,omitempty"` 
	Ts       *string `json:"ts,omitempty"`      
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
