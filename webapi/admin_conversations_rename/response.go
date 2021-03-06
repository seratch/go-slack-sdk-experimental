// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminConversationsRename, err := UnmarshalAdminConversationsRename(bytes)
//    bytes, err = adminConversationsRename.Marshal()

package admin_conversations_rename

import "encoding/json"

func UnmarshalAdminConversationsRename(data []byte) (AdminConversationsRename, error) {
	var r AdminConversationsRename
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminConversationsRename) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminConversationsRename struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
