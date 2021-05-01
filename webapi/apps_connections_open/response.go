// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    appsConnectionsOpen, err := UnmarshalAppsConnectionsOpen(bytes)
//    bytes, err = appsConnectionsOpen.Marshal()

package apps_connections_open

import "encoding/json"

func UnmarshalAppsConnectionsOpen(data []byte) (AppsConnectionsOpen, error) {
	var r AppsConnectionsOpen
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AppsConnectionsOpen) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AppsConnectionsOpen struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
	URL      *string `json:"url,omitempty"`     
}
