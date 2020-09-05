// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    channelsSetPurpose, err := UnmarshalChannelsSetPurpose(bytes)
//    bytes, err = channelsSetPurpose.Marshal()

package channels_setPurpose

import "encoding/json"

func UnmarshalChannelsSetPurpose(data []byte) (ChannelsSetPurpose, error) {
	var r ChannelsSetPurpose
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ChannelsSetPurpose) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ChannelsSetPurpose struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Purpose  *string `json:"purpose,omitempty"` 
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
