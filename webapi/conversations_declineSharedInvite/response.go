// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    conversationsDeclineSharedInvite, err := UnmarshalConversationsDeclineSharedInvite(bytes)
//    bytes, err = conversationsDeclineSharedInvite.Marshal()

package conversations_declineSharedInvite

import "encoding/json"

func UnmarshalConversationsDeclineSharedInvite(data []byte) (ConversationsDeclineSharedInvite, error) {
	var r ConversationsDeclineSharedInvite
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationsDeclineSharedInvite) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ConversationsDeclineSharedInvite struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
