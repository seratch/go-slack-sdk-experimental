// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    chatUnfurl, err := UnmarshalChatUnfurl(bytes)
//    bytes, err = chatUnfurl.Marshal()

package chat_unfurl

import "encoding/json"

func UnmarshalChatUnfurl(data []byte) (ChatUnfurl, error) {
	var r ChatUnfurl
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ChatUnfurl) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ChatUnfurl struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
