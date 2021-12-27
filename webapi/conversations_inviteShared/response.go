// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    conversationsInviteShared, err := UnmarshalConversationsInviteShared(bytes)
//    bytes, err = conversationsInviteShared.Marshal()

package conversations_inviteShared

import "encoding/json"

func UnmarshalConversationsInviteShared(data []byte) (ConversationsInviteShared, error) {
	var r ConversationsInviteShared
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationsInviteShared) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ConversationsInviteShared struct {
	Ok                    *bool   `json:"ok,omitempty"`                      
	Error                 *string `json:"error,omitempty"`                   
	URL                   *string `json:"url,omitempty"`                     
	InviteID              *string `json:"invite_id,omitempty"`               
	ConfCode              *string `json:"conf_code,omitempty"`               
	IsLegacySharedChannel *bool   `json:"is_legacy_shared_channel,omitempty"`
	Needed                *string `json:"needed,omitempty"`                  
	Provided              *string `json:"provided,omitempty"`                
}
