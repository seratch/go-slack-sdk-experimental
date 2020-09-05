// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminUsersSetRegular, err := UnmarshalAdminUsersSetRegular(bytes)
//    bytes, err = adminUsersSetRegular.Marshal()

package admin_users_setRegular

import "encoding/json"

func UnmarshalAdminUsersSetRegular(data []byte) (AdminUsersSetRegular, error) {
	var r AdminUsersSetRegular
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminUsersSetRegular) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminUsersSetRegular struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
