// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    appsPermissionsRequest, err := UnmarshalAppsPermissionsRequest(bytes)
//    bytes, err = appsPermissionsRequest.Marshal()

package apps_permissions_request

import "encoding/json"

func UnmarshalAppsPermissionsRequest(data []byte) (AppsPermissionsRequest, error) {
	var r AppsPermissionsRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AppsPermissionsRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AppsPermissionsRequest struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
