// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminConversationsSetCustomRetention, err := UnmarshalAdminConversationsSetCustomRetention(bytes)
//    bytes, err = adminConversationsSetCustomRetention.Marshal()

package admin_conversations_setCustomRetention

import "encoding/json"

func UnmarshalAdminConversationsSetCustomRetention(data []byte) (AdminConversationsSetCustomRetention, error) {
	var r AdminConversationsSetCustomRetention
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminConversationsSetCustomRetention) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminConversationsSetCustomRetention struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
