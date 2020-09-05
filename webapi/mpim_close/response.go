// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    mpimClose, err := UnmarshalMpimClose(bytes)
//    bytes, err = mpimClose.Marshal()

package mpim_close

import "encoding/json"

func UnmarshalMpimClose(data []byte) (MpimClose, error) {
	var r MpimClose
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MpimClose) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MpimClose struct {
	Ok               *bool             `json:"ok,omitempty"`               
	Warning          *string           `json:"warning,omitempty"`          
	ResponseMetadata *ResponseMetadata `json:"response_metadata,omitempty"`
	Error            *string           `json:"error,omitempty"`            
	Needed           *string           `json:"needed,omitempty"`           
	Provided         *string           `json:"provided,omitempty"`         
}

type ResponseMetadata struct {
	Messages []string `json:"messages,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
}
