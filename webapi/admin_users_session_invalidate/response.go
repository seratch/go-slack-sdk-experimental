// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminUsersSessionInvalidate, err := UnmarshalAdminUsersSessionInvalidate(bytes)
//    bytes, err = adminUsersSessionInvalidate.Marshal()

package admin_users_session_invalidate

import "encoding/json"

func UnmarshalAdminUsersSessionInvalidate(data []byte) (AdminUsersSessionInvalidate, error) {
	var r AdminUsersSessionInvalidate
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminUsersSessionInvalidate) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminUsersSessionInvalidate struct {
	Ok               *bool             `json:"ok,omitempty"`               
	Error            *string           `json:"error,omitempty"`            
	ResponseMetadata *ResponseMetadata `json:"response_metadata,omitempty"`
	Needed           *string           `json:"needed,omitempty"`           
	Provided         *string           `json:"provided,omitempty"`         
}

type ResponseMetadata struct {
	Messages []string `json:"messages,omitempty"`
}
