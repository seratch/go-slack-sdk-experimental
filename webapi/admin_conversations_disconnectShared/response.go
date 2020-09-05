// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminConversationsDisconnectShared, err := UnmarshalAdminConversationsDisconnectShared(bytes)
//    bytes, err = adminConversationsDisconnectShared.Marshal()

package admin_conversations_disconnectShared

import "encoding/json"

func UnmarshalAdminConversationsDisconnectShared(data []byte) (AdminConversationsDisconnectShared, error) {
	var r AdminConversationsDisconnectShared
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminConversationsDisconnectShared) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminConversationsDisconnectShared struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
