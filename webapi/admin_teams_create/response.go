// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminTeamsCreate, err := UnmarshalAdminTeamsCreate(bytes)
//    bytes, err = adminTeamsCreate.Marshal()

package admin_teams_create

import "encoding/json"

func UnmarshalAdminTeamsCreate(data []byte) (AdminTeamsCreate, error) {
	var r AdminTeamsCreate
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminTeamsCreate) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminTeamsCreate struct {
	Ok               *bool             `json:"ok,omitempty"`               
	Error            *string           `json:"error,omitempty"`            
	ResponseMetadata *ResponseMetadata `json:"response_metadata,omitempty"`
	Team             *string           `json:"team,omitempty"`             
	Needed           *string           `json:"needed,omitempty"`           
	Provided         *string           `json:"provided,omitempty"`         
}

type ResponseMetadata struct {
	Messages []string `json:"messages,omitempty"`
}
