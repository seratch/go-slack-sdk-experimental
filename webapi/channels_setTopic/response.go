// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    channelsSetTopic, err := UnmarshalChannelsSetTopic(bytes)
//    bytes, err = channelsSetTopic.Marshal()

package channels_setTopic

import "encoding/json"

func UnmarshalChannelsSetTopic(data []byte) (ChannelsSetTopic, error) {
	var r ChannelsSetTopic
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ChannelsSetTopic) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ChannelsSetTopic struct {
	Topic    *string `json:"topic,omitempty"`   
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
