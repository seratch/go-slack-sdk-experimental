// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    channelsMark, err := UnmarshalChannelsMark(bytes)
//    bytes, err = channelsMark.Marshal()

package channels_mark

import "encoding/json"

func UnmarshalChannelsMark(data []byte) (ChannelsMark, error) {
	var r ChannelsMark
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ChannelsMark) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ChannelsMark struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
