// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminUsersRemove, err := UnmarshalAdminUsersRemove(bytes)
//    bytes, err = adminUsersRemove.Marshal()

package admin_users_remove

import "encoding/json"

func UnmarshalAdminUsersRemove(data []byte) (AdminUsersRemove, error) {
	var r AdminUsersRemove
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminUsersRemove) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminUsersRemove struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
