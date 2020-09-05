// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminConversationsRestrictAccessListGroups, err := UnmarshalAdminConversationsRestrictAccessListGroups(bytes)
//    bytes, err = adminConversationsRestrictAccessListGroups.Marshal()

package admin_conversations_restrictAccess_listGroups

import "encoding/json"

func UnmarshalAdminConversationsRestrictAccessListGroups(data []byte) (AdminConversationsRestrictAccessListGroups, error) {
	var r AdminConversationsRestrictAccessListGroups
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminConversationsRestrictAccessListGroups) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminConversationsRestrictAccessListGroups struct {
	Ok               *bool             `json:"ok,omitempty"`               
	GroupIDS         []string          `json:"group_ids,omitempty"`        
	Error            *string           `json:"error,omitempty"`            
	ResponseMetadata *ResponseMetadata `json:"response_metadata,omitempty"`
	Needed           *string           `json:"needed,omitempty"`           
	Provided         *string           `json:"provided,omitempty"`         
}

type ResponseMetadata struct {
	Messages []string `json:"messages,omitempty"`
}
