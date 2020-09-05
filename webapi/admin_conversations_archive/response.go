// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminConversationsArchive, err := UnmarshalAdminConversationsArchive(bytes)
//    bytes, err = adminConversationsArchive.Marshal()

package admin_conversations_archive

import "encoding/json"

func UnmarshalAdminConversationsArchive(data []byte) (AdminConversationsArchive, error) {
	var r AdminConversationsArchive
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminConversationsArchive) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminConversationsArchive struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
