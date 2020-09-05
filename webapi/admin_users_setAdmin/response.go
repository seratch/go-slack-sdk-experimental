// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminUsersSetAdmin, err := UnmarshalAdminUsersSetAdmin(bytes)
//    bytes, err = adminUsersSetAdmin.Marshal()

package admin_users_setAdmin

import "encoding/json"

func UnmarshalAdminUsersSetAdmin(data []byte) (AdminUsersSetAdmin, error) {
	var r AdminUsersSetAdmin
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminUsersSetAdmin) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminUsersSetAdmin struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
