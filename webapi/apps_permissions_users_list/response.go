// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    appsPermissionsUsersList, err := UnmarshalAppsPermissionsUsersList(bytes)
//    bytes, err = appsPermissionsUsersList.Marshal()

package apps_permissions_users_list

import "encoding/json"

func UnmarshalAppsPermissionsUsersList(data []byte) (AppsPermissionsUsersList, error) {
	var r AppsPermissionsUsersList
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AppsPermissionsUsersList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AppsPermissionsUsersList struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
