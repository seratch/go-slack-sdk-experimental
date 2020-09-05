// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    pinsRemove, err := UnmarshalPinsRemove(bytes)
//    bytes, err = pinsRemove.Marshal()

package pins_remove

import "encoding/json"

func UnmarshalPinsRemove(data []byte) (PinsRemove, error) {
	var r PinsRemove
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PinsRemove) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type PinsRemove struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
