// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    appsPermissionsUsersRequest, err := UnmarshalAppsPermissionsUsersRequest(bytes)
//    bytes, err = appsPermissionsUsersRequest.Marshal()

package apps_permissions_users_request

import "encoding/json"

func UnmarshalAppsPermissionsUsersRequest(data []byte) (AppsPermissionsUsersRequest, error) {
	var r AppsPermissionsUsersRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AppsPermissionsUsersRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AppsPermissionsUsersRequest struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
