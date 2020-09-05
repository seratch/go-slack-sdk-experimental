// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminUsersAssign, err := UnmarshalAdminUsersAssign(bytes)
//    bytes, err = adminUsersAssign.Marshal()

package admin_users_assign

import "encoding/json"

func UnmarshalAdminUsersAssign(data []byte) (AdminUsersAssign, error) {
	var r AdminUsersAssign
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminUsersAssign) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminUsersAssign struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
