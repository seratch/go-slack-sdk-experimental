// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    rtmConnect, err := UnmarshalRtmConnect(bytes)
//    bytes, err = rtmConnect.Marshal()

package rtm_connect

import "encoding/json"

func UnmarshalRtmConnect(data []byte) (RtmConnect, error) {
	var r RtmConnect
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RtmConnect) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type RtmConnect struct {
	Ok       *bool   `json:"ok,omitempty"`      
	URL      *string `json:"url,omitempty"`     
	Team     *Team   `json:"team,omitempty"`    
	Self     *Self   `json:"self,omitempty"`    
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}

type Self struct {
	ID   *string `json:"id,omitempty"`  
	Name *string `json:"name,omitempty"`
}

type Team struct {
	ID     *string `json:"id,omitempty"`    
	Name   *string `json:"name,omitempty"`  
	Domain *string `json:"domain,omitempty"`
}
