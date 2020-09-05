// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminConversationsUnarchive, err := UnmarshalAdminConversationsUnarchive(bytes)
//    bytes, err = adminConversationsUnarchive.Marshal()

package admin_conversations_unarchive

import "encoding/json"

func UnmarshalAdminConversationsUnarchive(data []byte) (AdminConversationsUnarchive, error) {
	var r AdminConversationsUnarchive
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminConversationsUnarchive) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminConversationsUnarchive struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
