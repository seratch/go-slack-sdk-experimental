// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminUsersSessionSetSettings, err := UnmarshalAdminUsersSessionSetSettings(bytes)
//    bytes, err = adminUsersSessionSetSettings.Marshal()

package admin_users_session_setSettings

import "encoding/json"

func UnmarshalAdminUsersSessionSetSettings(data []byte) (AdminUsersSessionSetSettings, error) {
	var r AdminUsersSessionSetSettings
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminUsersSessionSetSettings) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminUsersSessionSetSettings struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
