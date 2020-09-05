// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminInviteRequestsApprove, err := UnmarshalAdminInviteRequestsApprove(bytes)
//    bytes, err = adminInviteRequestsApprove.Marshal()

package admin_inviteRequests_approve

import "encoding/json"

func UnmarshalAdminInviteRequestsApprove(data []byte) (AdminInviteRequestsApprove, error) {
	var r AdminInviteRequestsApprove
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminInviteRequestsApprove) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminInviteRequestsApprove struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
