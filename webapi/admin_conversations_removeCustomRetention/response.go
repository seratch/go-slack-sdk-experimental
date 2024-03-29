// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminConversationsRemoveCustomRetention, err := UnmarshalAdminConversationsRemoveCustomRetention(bytes)
//    bytes, err = adminConversationsRemoveCustomRetention.Marshal()

package admin_conversations_removeCustomRetention

import "encoding/json"

func UnmarshalAdminConversationsRemoveCustomRetention(data []byte) (AdminConversationsRemoveCustomRetention, error) {
	var r AdminConversationsRemoveCustomRetention
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminConversationsRemoveCustomRetention) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminConversationsRemoveCustomRetention struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
