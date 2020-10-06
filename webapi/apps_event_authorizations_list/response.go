// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    appsEventAuthorizationsList, err := UnmarshalAppsEventAuthorizationsList(bytes)
//    bytes, err = appsEventAuthorizationsList.Marshal()

package apps_event_authorizations_list

import "encoding/json"

func UnmarshalAppsEventAuthorizationsList(data []byte) (AppsEventAuthorizationsList, error) {
	var r AppsEventAuthorizationsList
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AppsEventAuthorizationsList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AppsEventAuthorizationsList struct {
	Ok             *bool           `json:"ok,omitempty"`            
	Error          *string         `json:"error,omitempty"`         
	Authorizations []Authorization `json:"authorizations,omitempty"`
	Needed         *string         `json:"needed,omitempty"`        
	Provided       *string         `json:"provided,omitempty"`      
}

type Authorization struct {
	EnterpriseID        *string `json:"enterprise_id,omitempty"`        
	TeamID              *string `json:"team_id,omitempty"`              
	UserID              *string `json:"user_id,omitempty"`              
	IsBot               *bool   `json:"is_bot,omitempty"`               
	IsEnterpriseInstall *bool   `json:"is_enterprise_install,omitempty"`
}
