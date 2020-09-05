// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminAppsApprove, err := UnmarshalAdminAppsApprove(bytes)
//    bytes, err = adminAppsApprove.Marshal()

package admin_apps_approve

import "encoding/json"

func UnmarshalAdminAppsApprove(data []byte) (AdminAppsApprove, error) {
	var r AdminAppsApprove
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminAppsApprove) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminAppsApprove struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
