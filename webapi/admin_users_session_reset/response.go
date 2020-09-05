// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminUsersSessionReset, err := UnmarshalAdminUsersSessionReset(bytes)
//    bytes, err = adminUsersSessionReset.Marshal()

package admin_users_session_reset

import "encoding/json"

func UnmarshalAdminUsersSessionReset(data []byte) (AdminUsersSessionReset, error) {
	var r AdminUsersSessionReset
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminUsersSessionReset) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminUsersSessionReset struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
