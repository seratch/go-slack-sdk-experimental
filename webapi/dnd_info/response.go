// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    dNDInfo, err := UnmarshalDNDInfo(bytes)
//    bytes, err = dNDInfo.Marshal()

package dnd_info

import "encoding/json"

func UnmarshalDNDInfo(data []byte) (DNDInfo, error) {
	var r DNDInfo
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DNDInfo) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DNDInfo struct {
	Ok             *bool   `json:"ok,omitempty"`               
	DNDEnabled     *bool   `json:"dnd_enabled,omitempty"`      
	NextDNDStartTs *int64  `json:"next_dnd_start_ts,omitempty"`
	NextDNDEndTs   *int64  `json:"next_dnd_end_ts,omitempty"`  
	Error          *string `json:"error,omitempty"`            
	Needed         *string `json:"needed,omitempty"`           
	Provided       *string `json:"provided,omitempty"`         
}
