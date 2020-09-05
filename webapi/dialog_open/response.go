// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    dialogOpen, err := UnmarshalDialogOpen(bytes)
//    bytes, err = dialogOpen.Marshal()

package dialog_open

import "encoding/json"

func UnmarshalDialogOpen(data []byte) (DialogOpen, error) {
	var r DialogOpen
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DialogOpen) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DialogOpen struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
