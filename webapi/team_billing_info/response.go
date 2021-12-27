// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    teamBillingInfo, err := UnmarshalTeamBillingInfo(bytes)
//    bytes, err = teamBillingInfo.Marshal()

package team_billing_info

import "encoding/json"

func UnmarshalTeamBillingInfo(data []byte) (TeamBillingInfo, error) {
	var r TeamBillingInfo
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TeamBillingInfo) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TeamBillingInfo struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
	Plan     *string `json:"plan,omitempty"`    
}
