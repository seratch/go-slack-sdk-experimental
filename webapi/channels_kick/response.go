// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    channelsKick, err := UnmarshalChannelsKick(bytes)
//    bytes, err = channelsKick.Marshal()

package channels_kick

import "encoding/json"

func UnmarshalChannelsKick(data []byte) (ChannelsKick, error) {
	var r ChannelsKick
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ChannelsKick) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ChannelsKick struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
