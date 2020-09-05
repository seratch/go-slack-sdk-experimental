// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    reactionsRemove, err := UnmarshalReactionsRemove(bytes)
//    bytes, err = reactionsRemove.Marshal()

package reactions_remove

import "encoding/json"

func UnmarshalReactionsRemove(data []byte) (ReactionsRemove, error) {
	var r ReactionsRemove
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ReactionsRemove) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ReactionsRemove struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
