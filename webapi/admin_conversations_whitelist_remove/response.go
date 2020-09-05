// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminConversationsWhitelistRemove, err := UnmarshalAdminConversationsWhitelistRemove(bytes)
//    bytes, err = adminConversationsWhitelistRemove.Marshal()

package admin_conversations_whitelist_remove

import "encoding/json"

func UnmarshalAdminConversationsWhitelistRemove(data []byte) (AdminConversationsWhitelistRemove, error) {
	var r AdminConversationsWhitelistRemove
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminConversationsWhitelistRemove) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminConversationsWhitelistRemove struct {
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
