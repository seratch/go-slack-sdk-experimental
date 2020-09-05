// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    viewsPublish, err := UnmarshalViewsPublish(bytes)
//    bytes, err = viewsPublish.Marshal()

package views_publish

import "encoding/json"

func UnmarshalViewsPublish(data []byte) (ViewsPublish, error) {
	var r ViewsPublish
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ViewsPublish) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ViewsPublish struct {
	Ok               *bool             `json:"ok,omitempty"`               
	Error            *string           `json:"error,omitempty"`            
	ResponseMetadata *ResponseMetadata `json:"response_metadata,omitempty"`
	Needed           *string           `json:"needed,omitempty"`           
	Provided         *string           `json:"provided,omitempty"`         
}

type ResponseMetadata struct {
	Messages []string `json:"messages,omitempty"`
}
