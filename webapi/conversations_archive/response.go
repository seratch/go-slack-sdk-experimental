// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    conversationsArchive, err := UnmarshalConversationsArchive(bytes)
//    bytes, err = conversationsArchive.Marshal()

package conversations_archive

import "encoding/json"

func UnmarshalConversationsArchive(data []byte) (ConversationsArchive, error) {
	var r ConversationsArchive
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationsArchive) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ConversationsArchive struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
