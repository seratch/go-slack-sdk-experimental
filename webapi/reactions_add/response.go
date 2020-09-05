// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    reactionsAdd, err := UnmarshalReactionsAdd(bytes)
//    bytes, err = reactionsAdd.Marshal()

package reactions_add

import "encoding/json"

func UnmarshalReactionsAdd(data []byte) (ReactionsAdd, error) {
	var r ReactionsAdd
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ReactionsAdd) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ReactionsAdd struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
