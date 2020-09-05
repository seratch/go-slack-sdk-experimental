// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminConversationsSetConversationPrefs, err := UnmarshalAdminConversationsSetConversationPrefs(bytes)
//    bytes, err = adminConversationsSetConversationPrefs.Marshal()

package admin_conversations_setConversationPrefs

import "encoding/json"

func UnmarshalAdminConversationsSetConversationPrefs(data []byte) (AdminConversationsSetConversationPrefs, error) {
	var r AdminConversationsSetConversationPrefs
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminConversationsSetConversationPrefs) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminConversationsSetConversationPrefs struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
