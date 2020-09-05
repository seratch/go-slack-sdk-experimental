// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminTeamsSettingsSetDiscoverability, err := UnmarshalAdminTeamsSettingsSetDiscoverability(bytes)
//    bytes, err = adminTeamsSettingsSetDiscoverability.Marshal()

package admin_teams_settings_setDiscoverability

import "encoding/json"

func UnmarshalAdminTeamsSettingsSetDiscoverability(data []byte) (AdminTeamsSettingsSetDiscoverability, error) {
	var r AdminTeamsSettingsSetDiscoverability
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminTeamsSettingsSetDiscoverability) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminTeamsSettingsSetDiscoverability struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
