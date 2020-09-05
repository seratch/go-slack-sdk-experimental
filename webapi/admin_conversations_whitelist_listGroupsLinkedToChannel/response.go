// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminConversationsWhitelistListGroupsLinkedToChannel, err := UnmarshalAdminConversationsWhitelistListGroupsLinkedToChannel(bytes)
//    bytes, err = adminConversationsWhitelistListGroupsLinkedToChannel.Marshal()

package admin_conversations_whitelist_listGroupsLinkedToChannel

import "encoding/json"

func UnmarshalAdminConversationsWhitelistListGroupsLinkedToChannel(data []byte) (AdminConversationsWhitelistListGroupsLinkedToChannel, error) {
	var r AdminConversationsWhitelistListGroupsLinkedToChannel
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminConversationsWhitelistListGroupsLinkedToChannel) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminConversationsWhitelistListGroupsLinkedToChannel struct {
	Ok               *bool             `json:"ok,omitempty"`               
	Error            *string           `json:"error,omitempty"`            
	GroupIDS         []string          `json:"group_ids,omitempty"`        
	ResponseMetadata *ResponseMetadata `json:"response_metadata,omitempty"`
	Warning          *string           `json:"warning,omitempty"`          
	Needed           *string           `json:"needed,omitempty"`           
	Provided         *string           `json:"provided,omitempty"`         
}

type ResponseMetadata struct {
	Messages []string `json:"messages,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
}
