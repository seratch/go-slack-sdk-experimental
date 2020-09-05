// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminTeamsOwnersList, err := UnmarshalAdminTeamsOwnersList(bytes)
//    bytes, err = adminTeamsOwnersList.Marshal()

package admin_teams_owners_list

import "encoding/json"

func UnmarshalAdminTeamsOwnersList(data []byte) (AdminTeamsOwnersList, error) {
	var r AdminTeamsOwnersList
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminTeamsOwnersList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminTeamsOwnersList struct {
	Ok               *bool             `json:"ok,omitempty"`               
	OwnerIDS         []string          `json:"owner_ids,omitempty"`        
	ResponseMetadata *ResponseMetadata `json:"response_metadata,omitempty"`
	Error            *string           `json:"error,omitempty"`            
	Needed           *string           `json:"needed,omitempty"`           
	Provided         *string           `json:"provided,omitempty"`         
}

type ResponseMetadata struct {
	NextCursor *string `json:"next_cursor,omitempty"`
}
