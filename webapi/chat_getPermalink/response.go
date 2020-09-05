// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    chatGetPermalink, err := UnmarshalChatGetPermalink(bytes)
//    bytes, err = chatGetPermalink.Marshal()

package chat_getPermalink

import "encoding/json"

func UnmarshalChatGetPermalink(data []byte) (ChatGetPermalink, error) {
	var r ChatGetPermalink
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ChatGetPermalink) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ChatGetPermalink struct {
	Ok        *bool   `json:"ok,omitempty"`       
	Permalink *string `json:"permalink,omitempty"`
	Channel   *string `json:"channel,omitempty"`  
	Error     *string `json:"error,omitempty"`    
	Needed    *string `json:"needed,omitempty"`   
	Provided  *string `json:"provided,omitempty"` 
}
