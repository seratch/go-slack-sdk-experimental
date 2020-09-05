// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    teamBillableInfo, err := UnmarshalTeamBillableInfo(bytes)
//    bytes, err = teamBillableInfo.Marshal()

package team_billableInfo

import "encoding/json"

func UnmarshalTeamBillableInfo(data []byte) (TeamBillableInfo, error) {
	var r TeamBillableInfo
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TeamBillableInfo) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TeamBillableInfo struct {
	Ok           *bool                   `json:"ok,omitempty"`           
	BillableInfo map[string]BillableInfo `json:"billable_info,omitempty"`
	Error        *string                 `json:"error,omitempty"`        
	Needed       *string                 `json:"needed,omitempty"`       
	Provided     *string                 `json:"provided,omitempty"`     
}

type BillableInfo struct {
	BillingActive *bool `json:"billing_active,omitempty"`
}
