// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    iMOpen, err := UnmarshalIMOpen(bytes)
//    bytes, err = iMOpen.Marshal()

package im_open

import "encoding/json"

func UnmarshalIMOpen(data []byte) (IMOpen, error) {
	var r IMOpen
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *IMOpen) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type IMOpen struct {
	Ok               *bool             `json:"ok,omitempty"`               
	Channel          *Channel          `json:"channel,omitempty"`          
	Warning          *string           `json:"warning,omitempty"`          
	ResponseMetadata *ResponseMetadata `json:"response_metadata,omitempty"`
	NoOp             *bool             `json:"no_op,omitempty"`            
	AlreadyOpen      *bool             `json:"already_open,omitempty"`     
	Error            *string           `json:"error,omitempty"`            
	Needed           *string           `json:"needed,omitempty"`           
	Provided         *string           `json:"provided,omitempty"`         
}

type Channel struct {
	ID *string `json:"id,omitempty"`
}

type ResponseMetadata struct {
	Messages []string `json:"messages,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
}
