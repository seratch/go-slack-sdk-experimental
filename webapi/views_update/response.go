// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    viewsUpdate, err := UnmarshalViewsUpdate(bytes)
//    bytes, err = viewsUpdate.Marshal()

package views_update

import "encoding/json"

func UnmarshalViewsUpdate(data []byte) (ViewsUpdate, error) {
	var r ViewsUpdate
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ViewsUpdate) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ViewsUpdate struct {
	Ok               *bool             `json:"ok,omitempty"`               
	Error            *string           `json:"error,omitempty"`            
	ResponseMetadata *ResponseMetadata `json:"response_metadata,omitempty"`
	Needed           *string           `json:"needed,omitempty"`           
	Provided         *string           `json:"provided,omitempty"`         
}

type ResponseMetadata struct {
	Messages []string `json:"messages,omitempty"`
}
