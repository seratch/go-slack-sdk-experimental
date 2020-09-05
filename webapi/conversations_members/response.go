// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    conversationsMembers, err := UnmarshalConversationsMembers(bytes)
//    bytes, err = conversationsMembers.Marshal()

package conversations_members

import "encoding/json"

func UnmarshalConversationsMembers(data []byte) (ConversationsMembers, error) {
	var r ConversationsMembers
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationsMembers) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ConversationsMembers struct {
	Ok               *bool             `json:"ok,omitempty"`               
	Members          []string          `json:"members,omitempty"`          
	ResponseMetadata *ResponseMetadata `json:"response_metadata,omitempty"`
	Error            *string           `json:"error,omitempty"`            
	Needed           *string           `json:"needed,omitempty"`           
	Provided         *string           `json:"provided,omitempty"`         
}

type ResponseMetadata struct {
	NextCursor *string `json:"next_cursor,omitempty"`
}
