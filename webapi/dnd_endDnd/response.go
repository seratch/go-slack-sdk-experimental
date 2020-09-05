// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    dNDEndDND, err := UnmarshalDNDEndDND(bytes)
//    bytes, err = dNDEndDND.Marshal()

package dnd_endDnd

import "encoding/json"

func UnmarshalDNDEndDND(data []byte) (DNDEndDND, error) {
	var r DNDEndDND
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DNDEndDND) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DNDEndDND struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
