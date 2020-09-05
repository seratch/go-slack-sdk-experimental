// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminConversationsWhitelistAdd, err := UnmarshalAdminConversationsWhitelistAdd(bytes)
//    bytes, err = adminConversationsWhitelistAdd.Marshal()

package admin_conversations_whitelist_add

import "encoding/json"

func UnmarshalAdminConversationsWhitelistAdd(data []byte) (AdminConversationsWhitelistAdd, error) {
	var r AdminConversationsWhitelistAdd
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminConversationsWhitelistAdd) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminConversationsWhitelistAdd struct {
	Ok               *bool             `json:"ok,omitempty"`               
	Error            *string           `json:"error,omitempty"`            
	ResponseMetadata *ResponseMetadata `json:"response_metadata,omitempty"`
	Warning          *string           `json:"warning,omitempty"`          
	Needed           *string           `json:"needed,omitempty"`           
	Provided         *string           `json:"provided,omitempty"`         
}

type ResponseMetadata struct {
	Messages []string `json:"messages,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
}
