// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    chatDeleteScheduledMessage, err := UnmarshalChatDeleteScheduledMessage(bytes)
//    bytes, err = chatDeleteScheduledMessage.Marshal()

package chat_deleteScheduledMessage

import "encoding/json"

func UnmarshalChatDeleteScheduledMessage(data []byte) (ChatDeleteScheduledMessage, error) {
	var r ChatDeleteScheduledMessage
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ChatDeleteScheduledMessage) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ChatDeleteScheduledMessage struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
