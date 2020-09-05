// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminInviteRequestsList, err := UnmarshalAdminInviteRequestsList(bytes)
//    bytes, err = adminInviteRequestsList.Marshal()

package admin_inviteRequests_list

import "encoding/json"

func UnmarshalAdminInviteRequestsList(data []byte) (AdminInviteRequestsList, error) {
	var r AdminInviteRequestsList
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminInviteRequestsList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminInviteRequestsList struct {
	Ok             *bool    `json:"ok,omitempty"`             
	InviteRequests []string `json:"invite_requests,omitempty"`
	Error          *string  `json:"error,omitempty"`          
	Needed         *string  `json:"needed,omitempty"`         
	Provided       *string  `json:"provided,omitempty"`       
}
