// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminTeamsAdminsList, err := UnmarshalAdminTeamsAdminsList(bytes)
//    bytes, err = adminTeamsAdminsList.Marshal()

package admin_teams_admins_list

import "encoding/json"

func UnmarshalAdminTeamsAdminsList(data []byte) (AdminTeamsAdminsList, error) {
	var r AdminTeamsAdminsList
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminTeamsAdminsList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminTeamsAdminsList struct {
	Ok               *bool             `json:"ok,omitempty"`               
	AdminIDS         []string          `json:"admin_ids,omitempty"`        
	ResponseMetadata *ResponseMetadata `json:"response_metadata,omitempty"`
	Error            *string           `json:"error,omitempty"`            
	Needed           *string           `json:"needed,omitempty"`           
	Provided         *string           `json:"provided,omitempty"`         
}

type ResponseMetadata struct {
	NextCursor *string `json:"next_cursor,omitempty"`
}
