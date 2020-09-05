// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    dNDTeamInfo, err := UnmarshalDNDTeamInfo(bytes)
//    bytes, err = dNDTeamInfo.Marshal()

package dnd_teamInfo

import "encoding/json"

func UnmarshalDNDTeamInfo(data []byte) (DNDTeamInfo, error) {
	var r DNDTeamInfo
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DNDTeamInfo) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DNDTeamInfo struct {
	Ok       *bool           `json:"ok,omitempty"`      
	Users    map[string]User `json:"users,omitempty"`   
	Error    *string         `json:"error,omitempty"`   
	Needed   *string         `json:"needed,omitempty"`  
	Provided *string         `json:"provided,omitempty"`
}

type User struct {
	DNDEnabled     *bool  `json:"dnd_enabled,omitempty"`      
	NextDNDStartTs *int64 `json:"next_dnd_start_ts,omitempty"`
	NextDNDEndTs   *int64 `json:"next_dnd_end_ts,omitempty"`  
}
