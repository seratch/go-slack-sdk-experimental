// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminUsersSetOwner, err := UnmarshalAdminUsersSetOwner(bytes)
//    bytes, err = adminUsersSetOwner.Marshal()

package admin_users_setOwner

import "encoding/json"

func UnmarshalAdminUsersSetOwner(data []byte) (AdminUsersSetOwner, error) {
	var r AdminUsersSetOwner
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminUsersSetOwner) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminUsersSetOwner struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
