// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminInviteRequestsApprovedList, err := UnmarshalAdminInviteRequestsApprovedList(bytes)
//    bytes, err = adminInviteRequestsApprovedList.Marshal()

package admin_inviteRequests_approved_list

import "encoding/json"

func UnmarshalAdminInviteRequestsApprovedList(data []byte) (AdminInviteRequestsApprovedList, error) {
	var r AdminInviteRequestsApprovedList
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminInviteRequestsApprovedList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminInviteRequestsApprovedList struct {
	Ok               *bool    `json:"ok,omitempty"`               
	ApprovedRequests []string `json:"approved_requests,omitempty"`
	Error            *string  `json:"error,omitempty"`            
	Needed           *string  `json:"needed,omitempty"`           
	Provided         *string  `json:"provided,omitempty"`         
}
