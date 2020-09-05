// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminTeamsSettingsSetDescription, err := UnmarshalAdminTeamsSettingsSetDescription(bytes)
//    bytes, err = adminTeamsSettingsSetDescription.Marshal()

package admin_teams_settings_setDescription

import "encoding/json"

func UnmarshalAdminTeamsSettingsSetDescription(data []byte) (AdminTeamsSettingsSetDescription, error) {
	var r AdminTeamsSettingsSetDescription
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminTeamsSettingsSetDescription) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminTeamsSettingsSetDescription struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
