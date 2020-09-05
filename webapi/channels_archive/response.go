// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    channelsArchive, err := UnmarshalChannelsArchive(bytes)
//    bytes, err = channelsArchive.Marshal()

package channels_archive

import "encoding/json"

func UnmarshalChannelsArchive(data []byte) (ChannelsArchive, error) {
	var r ChannelsArchive
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ChannelsArchive) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ChannelsArchive struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
