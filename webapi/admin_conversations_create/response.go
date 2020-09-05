// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminConversationsCreate, err := UnmarshalAdminConversationsCreate(bytes)
//    bytes, err = adminConversationsCreate.Marshal()

package admin_conversations_create

import "encoding/json"

func UnmarshalAdminConversationsCreate(data []byte) (AdminConversationsCreate, error) {
	var r AdminConversationsCreate
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminConversationsCreate) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminConversationsCreate struct {
	Ok        *bool   `json:"ok,omitempty"`        
	ChannelID *string `json:"channel_id,omitempty"`
	Error     *string `json:"error,omitempty"`     
	Needed    *string `json:"needed,omitempty"`    
	Provided  *string `json:"provided,omitempty"`  
}
