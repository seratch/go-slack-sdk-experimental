// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminConversationsDelete, err := UnmarshalAdminConversationsDelete(bytes)
//    bytes, err = adminConversationsDelete.Marshal()

package admin_conversations_delete

import "encoding/json"

func UnmarshalAdminConversationsDelete(data []byte) (AdminConversationsDelete, error) {
	var r AdminConversationsDelete
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminConversationsDelete) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminConversationsDelete struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
