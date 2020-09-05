// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminConversationsRestrictAccessRemoveGroup, err := UnmarshalAdminConversationsRestrictAccessRemoveGroup(bytes)
//    bytes, err = adminConversationsRestrictAccessRemoveGroup.Marshal()

package admin_conversations_restrictAccess_removeGroup

import "encoding/json"

func UnmarshalAdminConversationsRestrictAccessRemoveGroup(data []byte) (AdminConversationsRestrictAccessRemoveGroup, error) {
	var r AdminConversationsRestrictAccessRemoveGroup
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminConversationsRestrictAccessRemoveGroup) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminConversationsRestrictAccessRemoveGroup struct {
	Ok               *bool             `json:"ok,omitempty"`               
	Error            *string           `json:"error,omitempty"`            
	ResponseMetadata *ResponseMetadata `json:"response_metadata,omitempty"`
	Needed           *string           `json:"needed,omitempty"`           
	Provided         *string           `json:"provided,omitempty"`         
}

type ResponseMetadata struct {
	Messages []string `json:"messages,omitempty"`
}
