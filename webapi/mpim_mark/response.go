// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    mpimMark, err := UnmarshalMpimMark(bytes)
//    bytes, err = mpimMark.Marshal()

package mpim_mark

import "encoding/json"

func UnmarshalMpimMark(data []byte) (MpimMark, error) {
	var r MpimMark
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MpimMark) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MpimMark struct {
	Ok               *bool             `json:"ok,omitempty"`               
	Error            *string           `json:"error,omitempty"`            
	Warning          *string           `json:"warning,omitempty"`          
	ResponseMetadata *ResponseMetadata `json:"response_metadata,omitempty"`
	Needed           *string           `json:"needed,omitempty"`           
	Provided         *string           `json:"provided,omitempty"`         
}

type ResponseMetadata struct {
	Messages []string `json:"messages,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
}
