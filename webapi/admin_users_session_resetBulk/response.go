// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminUsersSessionResetBulk, err := UnmarshalAdminUsersSessionResetBulk(bytes)
//    bytes, err = adminUsersSessionResetBulk.Marshal()

package admin_users_session_resetBulk

import "encoding/json"

func UnmarshalAdminUsersSessionResetBulk(data []byte) (AdminUsersSessionResetBulk, error) {
	var r AdminUsersSessionResetBulk
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminUsersSessionResetBulk) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminUsersSessionResetBulk struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
