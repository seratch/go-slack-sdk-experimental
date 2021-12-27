// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    conversationsAcceptSharedInvite, err := UnmarshalConversationsAcceptSharedInvite(bytes)
//    bytes, err = conversationsAcceptSharedInvite.Marshal()

package conversations_acceptSharedInvite

import "encoding/json"

func UnmarshalConversationsAcceptSharedInvite(data []byte) (ConversationsAcceptSharedInvite, error) {
	var r ConversationsAcceptSharedInvite
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationsAcceptSharedInvite) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ConversationsAcceptSharedInvite struct {
	Ok               *bool   `json:"ok,omitempty"`               
	Error            *string `json:"error,omitempty"`            
	ImplicitApproval *bool   `json:"implicit_approval,omitempty"`
	ChannelID        *string `json:"channel_id,omitempty"`       
	InviteID         *string `json:"invite_id,omitempty"`        
	CanOpenScdm      *bool   `json:"can_open_scdm,omitempty"`    
	Needed           *string `json:"needed,omitempty"`           
	Provided         *string `json:"provided,omitempty"`         
}
