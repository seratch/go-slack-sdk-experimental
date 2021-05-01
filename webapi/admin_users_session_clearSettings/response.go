// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminUsersSessionClearSettings, err := UnmarshalAdminUsersSessionClearSettings(bytes)
//    bytes, err = adminUsersSessionClearSettings.Marshal()

package admin_users_session_clearSettings

import "encoding/json"

func UnmarshalAdminUsersSessionClearSettings(data []byte) (AdminUsersSessionClearSettings, error) {
	var r AdminUsersSessionClearSettings
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminUsersSessionClearSettings) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminUsersSessionClearSettings struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
