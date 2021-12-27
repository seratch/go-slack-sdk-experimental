// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    teamInfo, err := UnmarshalTeamInfo(bytes)
//    bytes, err = teamInfo.Marshal()

package team_info

import "encoding/json"

func UnmarshalTeamInfo(data []byte) (TeamInfo, error) {
	var r TeamInfo
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TeamInfo) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TeamInfo struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Team     *Team   `json:"team,omitempty"`    
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}

type Team struct {
	ID               *string `json:"id,omitempty"`               
	Name             *string `json:"name,omitempty"`             
	Domain           *string `json:"domain,omitempty"`           
	EmailDomain      *string `json:"email_domain,omitempty"`     
	Icon             *Icon   `json:"icon,omitempty"`             
	IsVerified       *bool   `json:"is_verified,omitempty"`      
	URL              *string `json:"url,omitempty"`              
	EnterpriseID     *string `json:"enterprise_id,omitempty"`    
	EnterpriseName   *string `json:"enterprise_name,omitempty"`  
	EnterpriseDomain *string `json:"enterprise_domain,omitempty"`
	Discoverable     *string `json:"discoverable,omitempty"`     
}

type Icon struct {
	Image34       *string `json:"image_34,omitempty"`      
	Image44       *string `json:"image_44,omitempty"`      
	Image68       *string `json:"image_68,omitempty"`      
	Image88       *string `json:"image_88,omitempty"`      
	Image102      *string `json:"image_102,omitempty"`     
	Image132      *string `json:"image_132,omitempty"`     
	Image230      *string `json:"image_230,omitempty"`     
	ImageOriginal *string `json:"image_original,omitempty"`
}
