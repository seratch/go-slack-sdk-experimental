// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminConversationsGetConversationPrefs, err := UnmarshalAdminConversationsGetConversationPrefs(bytes)
//    bytes, err = adminConversationsGetConversationPrefs.Marshal()

package admin_conversations_getConversationPrefs

import "encoding/json"

func UnmarshalAdminConversationsGetConversationPrefs(data []byte) (AdminConversationsGetConversationPrefs, error) {
	var r AdminConversationsGetConversationPrefs
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminConversationsGetConversationPrefs) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminConversationsGetConversationPrefs struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Prefs    *Prefs  `json:"prefs,omitempty"`   
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}

type Prefs struct {
	WhoCanPost *CanThread `json:"who_can_post,omitempty"`
	CanThread  *CanThread `json:"can_thread,omitempty"`  
}

type CanThread struct {
	Type []string `json:"type,omitempty"`
	User []string `json:"user,omitempty"`
}
