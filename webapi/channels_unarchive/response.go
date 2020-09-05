// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    channelsUnarchive, err := UnmarshalChannelsUnarchive(bytes)
//    bytes, err = channelsUnarchive.Marshal()

package channels_unarchive

import "encoding/json"

func UnmarshalChannelsUnarchive(data []byte) (ChannelsUnarchive, error) {
	var r ChannelsUnarchive
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ChannelsUnarchive) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ChannelsUnarchive struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
