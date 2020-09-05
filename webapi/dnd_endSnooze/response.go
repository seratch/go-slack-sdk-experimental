// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    dNDEndSnooze, err := UnmarshalDNDEndSnooze(bytes)
//    bytes, err = dNDEndSnooze.Marshal()

package dnd_endSnooze

import "encoding/json"

func UnmarshalDNDEndSnooze(data []byte) (DNDEndSnooze, error) {
	var r DNDEndSnooze
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DNDEndSnooze) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DNDEndSnooze struct {
	Ok             *bool   `json:"ok,omitempty"`               
	Error          *string `json:"error,omitempty"`            
	DNDEnabled     *bool   `json:"dnd_enabled,omitempty"`      
	NextDNDStartTs *int64  `json:"next_dnd_start_ts,omitempty"`
	NextDNDEndTs   *int64  `json:"next_dnd_end_ts,omitempty"`  
	SnoozeEnabled  *bool   `json:"snooze_enabled,omitempty"`   
	Needed         *string `json:"needed,omitempty"`           
	Provided       *string `json:"provided,omitempty"`         
}
