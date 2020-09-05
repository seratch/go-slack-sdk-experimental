// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    viewsOpen, err := UnmarshalViewsOpen(bytes)
//    bytes, err = viewsOpen.Marshal()

package views_open

import "encoding/json"

func UnmarshalViewsOpen(data []byte) (ViewsOpen, error) {
	var r ViewsOpen
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ViewsOpen) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ViewsOpen struct {
	Ok               *bool             `json:"ok,omitempty"`               
	Error            *string           `json:"error,omitempty"`            
	ResponseMetadata *ResponseMetadata `json:"response_metadata,omitempty"`
	Needed           *string           `json:"needed,omitempty"`           
	Provided         *string           `json:"provided,omitempty"`         
}

type ResponseMetadata struct {
	Messages []string `json:"messages,omitempty"`
}
