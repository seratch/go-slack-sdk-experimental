// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    chatPostEphemeral, err := UnmarshalChatPostEphemeral(bytes)
//    bytes, err = chatPostEphemeral.Marshal()

package chat_postEphemeral

import "encoding/json"

func UnmarshalChatPostEphemeral(data []byte) (ChatPostEphemeral, error) {
	var r ChatPostEphemeral
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ChatPostEphemeral) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ChatPostEphemeral struct {
	Ok        *bool   `json:"ok,omitempty"`        
	MessageTs *string `json:"message_ts,omitempty"`
	Error     *string `json:"error,omitempty"`     
	Needed    *string `json:"needed,omitempty"`    
	Provided  *string `json:"provided,omitempty"`  
}
