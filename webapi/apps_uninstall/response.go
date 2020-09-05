// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    appsUninstall, err := UnmarshalAppsUninstall(bytes)
//    bytes, err = appsUninstall.Marshal()

package apps_uninstall

import "encoding/json"

func UnmarshalAppsUninstall(data []byte) (AppsUninstall, error) {
	var r AppsUninstall
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AppsUninstall) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AppsUninstall struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
