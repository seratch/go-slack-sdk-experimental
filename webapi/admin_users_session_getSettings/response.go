// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminUsersSessionGetSettings, err := UnmarshalAdminUsersSessionGetSettings(bytes)
//    bytes, err = adminUsersSessionGetSettings.Marshal()

package admin_users_session_getSettings

import "encoding/json"

func UnmarshalAdminUsersSessionGetSettings(data []byte) (AdminUsersSessionGetSettings, error) {
	var r AdminUsersSessionGetSettings
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminUsersSessionGetSettings) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminUsersSessionGetSettings struct {
	Ok                *bool            `json:"ok,omitempty"`                 
	SessionSettings   []SessionSetting `json:"session_settings,omitempty"`   
	NoSettingsApplied []string         `json:"no_settings_applied,omitempty"`
	Error             *string          `json:"error,omitempty"`              
	Needed            *string          `json:"needed,omitempty"`             
	Provided          *string          `json:"provided,omitempty"`           
}

type SessionSetting struct {
	UserID                *string `json:"user_id,omitempty"`                 
	DesktopAppBrowserQuit *bool   `json:"desktop_app_browser_quit,omitempty"`
	Duration              *int64  `json:"duration,omitempty"`                
}
