// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminAppsUninstall, err := UnmarshalAdminAppsUninstall(bytes)
//    bytes, err = adminAppsUninstall.Marshal()

package admin_apps_uninstall

import "encoding/json"

func UnmarshalAdminAppsUninstall(data []byte) (AdminAppsUninstall, error) {
	var r AdminAppsUninstall
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminAppsUninstall) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminAppsUninstall struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Warning  *string `json:"warning,omitempty"` 
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
