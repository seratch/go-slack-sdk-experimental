// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    conversationsApproveSharedInvite, err := UnmarshalConversationsApproveSharedInvite(bytes)
//    bytes, err = conversationsApproveSharedInvite.Marshal()

package conversations_approveSharedInvite

import "encoding/json"

func UnmarshalConversationsApproveSharedInvite(data []byte) (ConversationsApproveSharedInvite, error) {
	var r ConversationsApproveSharedInvite
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationsApproveSharedInvite) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ConversationsApproveSharedInvite struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
