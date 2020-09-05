// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminTeamsSettingsSetName, err := UnmarshalAdminTeamsSettingsSetName(bytes)
//    bytes, err = adminTeamsSettingsSetName.Marshal()

package admin_teams_settings_setName

import "encoding/json"

func UnmarshalAdminTeamsSettingsSetName(data []byte) (AdminTeamsSettingsSetName, error) {
	var r AdminTeamsSettingsSetName
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminTeamsSettingsSetName) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminTeamsSettingsSetName struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
