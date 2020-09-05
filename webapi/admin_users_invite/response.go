// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminUsersInvite, err := UnmarshalAdminUsersInvite(bytes)
//    bytes, err = adminUsersInvite.Marshal()

package admin_users_invite

import "encoding/json"

func UnmarshalAdminUsersInvite(data []byte) (AdminUsersInvite, error) {
	var r AdminUsersInvite
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminUsersInvite) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminUsersInvite struct {
	Ok               *bool             `json:"ok,omitempty"`               
	Error            *string           `json:"error,omitempty"`            
	ResponseMetadata *ResponseMetadata `json:"response_metadata,omitempty"`
	Needed           *string           `json:"needed,omitempty"`           
	Provided         *string           `json:"provided,omitempty"`         
}

type ResponseMetadata struct {
	Messages []string `json:"messages,omitempty"`
}
