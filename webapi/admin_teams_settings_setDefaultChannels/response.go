// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminTeamsSettingsSetDefaultChannels, err := UnmarshalAdminTeamsSettingsSetDefaultChannels(bytes)
//    bytes, err = adminTeamsSettingsSetDefaultChannels.Marshal()

package admin_teams_settings_setDefaultChannels

import "encoding/json"

func UnmarshalAdminTeamsSettingsSetDefaultChannels(data []byte) (AdminTeamsSettingsSetDefaultChannels, error) {
	var r AdminTeamsSettingsSetDefaultChannels
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminTeamsSettingsSetDefaultChannels) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminTeamsSettingsSetDefaultChannels struct {
	Ok               *bool             `json:"ok,omitempty"`               
	Error            *string           `json:"error,omitempty"`            
	ResponseMetadata *ResponseMetadata `json:"response_metadata,omitempty"`
	Needed           *string           `json:"needed,omitempty"`           
	Provided         *string           `json:"provided,omitempty"`         
}

type ResponseMetadata struct {
	Messages []string `json:"messages,omitempty"`
}
