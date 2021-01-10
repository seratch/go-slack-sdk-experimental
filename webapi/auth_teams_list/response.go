// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    authTeamsList, err := UnmarshalAuthTeamsList(bytes)
//    bytes, err = authTeamsList.Marshal()

package auth_teams_list

import "encoding/json"

func UnmarshalAuthTeamsList(data []byte) (AuthTeamsList, error) {
	var r AuthTeamsList
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AuthTeamsList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AuthTeamsList struct {
	Ok               *bool             `json:"ok,omitempty"`               
	Teams            []Team            `json:"teams,omitempty"`            
	Error            *string           `json:"error,omitempty"`            
	ResponseMetadata *ResponseMetadata `json:"response_metadata,omitempty"`
	Needed           *string           `json:"needed,omitempty"`           
	Provided         *string           `json:"provided,omitempty"`         
}

type ResponseMetadata struct {
	NextCursor *string `json:"next_cursor,omitempty"`
}

type Team struct {
	ID   *string `json:"id,omitempty"`  
	Name *string `json:"name,omitempty"`
}
