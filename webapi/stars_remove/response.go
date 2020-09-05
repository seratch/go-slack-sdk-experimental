// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    starsRemove, err := UnmarshalStarsRemove(bytes)
//    bytes, err = starsRemove.Marshal()

package stars_remove

import "encoding/json"

func UnmarshalStarsRemove(data []byte) (StarsRemove, error) {
	var r StarsRemove
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *StarsRemove) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type StarsRemove struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
