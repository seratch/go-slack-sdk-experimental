// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    dNDSetSnooze, err := UnmarshalDNDSetSnooze(bytes)
//    bytes, err = dNDSetSnooze.Marshal()

package dnd_setSnooze

import "encoding/json"

func UnmarshalDNDSetSnooze(data []byte) (DNDSetSnooze, error) {
	var r DNDSetSnooze
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DNDSetSnooze) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DNDSetSnooze struct {
	Ok                 *bool   `json:"ok,omitempty"`                  
	Error              *string `json:"error,omitempty"`               
	SnoozeEnabled      *bool   `json:"snooze_enabled,omitempty"`      
	SnoozeEndtime      *int64  `json:"snooze_endtime,omitempty"`      
	SnoozeRemaining    *int64  `json:"snooze_remaining,omitempty"`    
	SnoozeIsIndefinite *bool   `json:"snooze_is_indefinite,omitempty"`
	Needed             *string `json:"needed,omitempty"`              
	Provided           *string `json:"provided,omitempty"`            
}
