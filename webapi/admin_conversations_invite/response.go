// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminConversationsInvite, err := UnmarshalAdminConversationsInvite(bytes)
//    bytes, err = adminConversationsInvite.Marshal()

package admin_conversations_invite

import "encoding/json"

func UnmarshalAdminConversationsInvite(data []byte) (AdminConversationsInvite, error) {
	var r AdminConversationsInvite
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminConversationsInvite) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminConversationsInvite struct {
	Ok            *bool          `json:"ok,omitempty"`             
	Error         *string        `json:"error,omitempty"`          
	FailedUserIDS *FailedUserIDS `json:"failed_user_ids,omitempty"`
	Needed        *string        `json:"needed,omitempty"`         
	Provided      *string        `json:"provided,omitempty"`       
}

type FailedUserIDS struct {
	U00000000 *string `json:"U00000000,omitempty"`
	U00000001 *string `json:"U00000001,omitempty"`
}
