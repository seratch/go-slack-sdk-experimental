// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    appsPermissionsInfo, err := UnmarshalAppsPermissionsInfo(bytes)
//    bytes, err = appsPermissionsInfo.Marshal()

package apps_permissions_info

import "encoding/json"

func UnmarshalAppsPermissionsInfo(data []byte) (AppsPermissionsInfo, error) {
	var r AppsPermissionsInfo
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AppsPermissionsInfo) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AppsPermissionsInfo struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
