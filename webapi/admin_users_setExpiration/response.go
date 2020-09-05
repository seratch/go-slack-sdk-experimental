// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminUsersSetExpiration, err := UnmarshalAdminUsersSetExpiration(bytes)
//    bytes, err = adminUsersSetExpiration.Marshal()

package admin_users_setExpiration

import "encoding/json"

func UnmarshalAdminUsersSetExpiration(data []byte) (AdminUsersSetExpiration, error) {
	var r AdminUsersSetExpiration
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminUsersSetExpiration) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminUsersSetExpiration struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
