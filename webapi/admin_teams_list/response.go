// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminTeamsList, err := UnmarshalAdminTeamsList(bytes)
//    bytes, err = adminTeamsList.Marshal()

package admin_teams_list

import "encoding/json"

func UnmarshalAdminTeamsList(data []byte) (AdminTeamsList, error) {
	var r AdminTeamsList
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminTeamsList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminTeamsList struct {
	Ok               *bool             `json:"ok,omitempty"`               
	Teams            []Team            `json:"teams,omitempty"`            
	ResponseMetadata *ResponseMetadata `json:"response_metadata,omitempty"`
	Error            *string           `json:"error,omitempty"`            
	Needed           *string           `json:"needed,omitempty"`           
	Provided         *string           `json:"provided,omitempty"`         
}

type ResponseMetadata struct {
	NextCursor *string `json:"next_cursor,omitempty"`
}

type Team struct {
	ID              *string       `json:"id,omitempty"`             
	Name            *string       `json:"name,omitempty"`           
	Discoverability *string       `json:"discoverability,omitempty"`
	PrimaryOwner    *PrimaryOwner `json:"primary_owner,omitempty"`  
	TeamURL         *string       `json:"team_url,omitempty"`       
}

type PrimaryOwner struct {
	UserID *string `json:"user_id,omitempty"`
	Email  *string `json:"email,omitempty"`  
}
