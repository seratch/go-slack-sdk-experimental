// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminAppsRestrict, err := UnmarshalAdminAppsRestrict(bytes)
//    bytes, err = adminAppsRestrict.Marshal()

package admin_apps_restrict

import "encoding/json"

func UnmarshalAdminAppsRestrict(data []byte) (AdminAppsRestrict, error) {
	var r AdminAppsRestrict
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminAppsRestrict) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminAppsRestrict struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
