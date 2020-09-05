// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    appsPermissionsResourcesList, err := UnmarshalAppsPermissionsResourcesList(bytes)
//    bytes, err = appsPermissionsResourcesList.Marshal()

package apps_permissions_resources_list

import "encoding/json"

func UnmarshalAppsPermissionsResourcesList(data []byte) (AppsPermissionsResourcesList, error) {
	var r AppsPermissionsResourcesList
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AppsPermissionsResourcesList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AppsPermissionsResourcesList struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
