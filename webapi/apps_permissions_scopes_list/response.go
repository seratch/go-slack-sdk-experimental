// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    appsPermissionsScopesList, err := UnmarshalAppsPermissionsScopesList(bytes)
//    bytes, err = appsPermissionsScopesList.Marshal()

package apps_permissions_scopes_list

import "encoding/json"

func UnmarshalAppsPermissionsScopesList(data []byte) (AppsPermissionsScopesList, error) {
	var r AppsPermissionsScopesList
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AppsPermissionsScopesList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AppsPermissionsScopesList struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
