// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminAppsClearResolution, err := UnmarshalAdminAppsClearResolution(bytes)
//    bytes, err = adminAppsClearResolution.Marshal()

package admin_apps_clearResolution

import "encoding/json"

func UnmarshalAdminAppsClearResolution(data []byte) (AdminAppsClearResolution, error) {
	var r AdminAppsClearResolution
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminAppsClearResolution) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminAppsClearResolution struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
