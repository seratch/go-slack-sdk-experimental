// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    pinsAdd, err := UnmarshalPinsAdd(bytes)
//    bytes, err = pinsAdd.Marshal()

package pins_add

import "encoding/json"

func UnmarshalPinsAdd(data []byte) (PinsAdd, error) {
	var r PinsAdd
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PinsAdd) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type PinsAdd struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
