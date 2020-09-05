// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminConversationsConvertToPrivate, err := UnmarshalAdminConversationsConvertToPrivate(bytes)
//    bytes, err = adminConversationsConvertToPrivate.Marshal()

package admin_conversations_convertToPrivate

import "encoding/json"

func UnmarshalAdminConversationsConvertToPrivate(data []byte) (AdminConversationsConvertToPrivate, error) {
	var r AdminConversationsConvertToPrivate
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminConversationsConvertToPrivate) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminConversationsConvertToPrivate struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
