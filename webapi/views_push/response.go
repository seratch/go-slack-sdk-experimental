// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    viewsPush, err := UnmarshalViewsPush(bytes)
//    bytes, err = viewsPush.Marshal()

package views_push

import "encoding/json"

func UnmarshalViewsPush(data []byte) (ViewsPush, error) {
	var r ViewsPush
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ViewsPush) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ViewsPush struct {
	Ok               *bool             `json:"ok,omitempty"`               
	Error            *string           `json:"error,omitempty"`            
	ResponseMetadata *ResponseMetadata `json:"response_metadata,omitempty"`
	Needed           *string           `json:"needed,omitempty"`           
	Provided         *string           `json:"provided,omitempty"`         
}

type ResponseMetadata struct {
	Messages []string `json:"messages,omitempty"`
}
