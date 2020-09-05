// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminInviteRequestsDeny, err := UnmarshalAdminInviteRequestsDeny(bytes)
//    bytes, err = adminInviteRequestsDeny.Marshal()

package admin_inviteRequests_deny

import "encoding/json"

func UnmarshalAdminInviteRequestsDeny(data []byte) (AdminInviteRequestsDeny, error) {
	var r AdminInviteRequestsDeny
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminInviteRequestsDeny) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminInviteRequestsDeny struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
