// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    channelsLeave, err := UnmarshalChannelsLeave(bytes)
//    bytes, err = channelsLeave.Marshal()

package channels_leave

import "encoding/json"

func UnmarshalChannelsLeave(data []byte) (ChannelsLeave, error) {
	var r ChannelsLeave
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ChannelsLeave) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ChannelsLeave struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
