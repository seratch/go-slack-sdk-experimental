// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    starsAdd, err := UnmarshalStarsAdd(bytes)
//    bytes, err = starsAdd.Marshal()

package stars_add

import "encoding/json"

func UnmarshalStarsAdd(data []byte) (StarsAdd, error) {
	var r StarsAdd
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *StarsAdd) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type StarsAdd struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
