// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminInviteRequestsDeniedList, err := UnmarshalAdminInviteRequestsDeniedList(bytes)
//    bytes, err = adminInviteRequestsDeniedList.Marshal()

package admin_inviteRequests_denied_list

import "encoding/json"

func UnmarshalAdminInviteRequestsDeniedList(data []byte) (AdminInviteRequestsDeniedList, error) {
	var r AdminInviteRequestsDeniedList
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminInviteRequestsDeniedList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminInviteRequestsDeniedList struct {
	Ok             *bool    `json:"ok,omitempty"`             
	DeniedRequests []string `json:"denied_requests,omitempty"`
	Error          *string  `json:"error,omitempty"`          
	Needed         *string  `json:"needed,omitempty"`         
	Provided       *string  `json:"provided,omitempty"`       
}
