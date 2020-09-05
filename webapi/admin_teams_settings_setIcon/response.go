// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminTeamsSettingsSetIcon, err := UnmarshalAdminTeamsSettingsSetIcon(bytes)
//    bytes, err = adminTeamsSettingsSetIcon.Marshal()

package admin_teams_settings_setIcon

import "encoding/json"

func UnmarshalAdminTeamsSettingsSetIcon(data []byte) (AdminTeamsSettingsSetIcon, error) {
	var r AdminTeamsSettingsSetIcon
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminTeamsSettingsSetIcon) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminTeamsSettingsSetIcon struct {
	Ok               *bool             `json:"ok,omitempty"`               
	Error            *string           `json:"error,omitempty"`            
	ResponseMetadata *ResponseMetadata `json:"response_metadata,omitempty"`
	Needed           *string           `json:"needed,omitempty"`           
	Provided         *string           `json:"provided,omitempty"`         
}

type ResponseMetadata struct {
	Messages []string `json:"messages,omitempty"`
}
