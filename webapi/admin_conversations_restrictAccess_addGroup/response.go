// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminConversationsRestrictAccessAddGroup, err := UnmarshalAdminConversationsRestrictAccessAddGroup(bytes)
//    bytes, err = adminConversationsRestrictAccessAddGroup.Marshal()

package admin_conversations_restrictAccess_addGroup

import "encoding/json"

func UnmarshalAdminConversationsRestrictAccessAddGroup(data []byte) (AdminConversationsRestrictAccessAddGroup, error) {
	var r AdminConversationsRestrictAccessAddGroup
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminConversationsRestrictAccessAddGroup) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminConversationsRestrictAccessAddGroup struct {
	Ok               *bool             `json:"ok,omitempty"`               
	Error            *string           `json:"error,omitempty"`            
	ResponseMetadata *ResponseMetadata `json:"response_metadata,omitempty"`
	Needed           *string           `json:"needed,omitempty"`           
	Provided         *string           `json:"provided,omitempty"`         
}

type ResponseMetadata struct {
	Messages []string `json:"messages,omitempty"`
}
