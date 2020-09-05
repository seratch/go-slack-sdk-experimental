// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    teamAccessLogs, err := UnmarshalTeamAccessLogs(bytes)
//    bytes, err = teamAccessLogs.Marshal()

package team_accessLogs

import "encoding/json"

func UnmarshalTeamAccessLogs(data []byte) (TeamAccessLogs, error) {
	var r TeamAccessLogs
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TeamAccessLogs) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TeamAccessLogs struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Logins   []Login `json:"logins,omitempty"`  
	Paging   *Paging `json:"paging,omitempty"`  
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}

type Login struct {
	UserID    *string `json:"user_id,omitempty"`   
	Username  *string `json:"username,omitempty"`  
	DateFirst *int64  `json:"date_first,omitempty"`
	DateLast  *int64  `json:"date_last,omitempty"` 
	Count     *int64  `json:"count,omitempty"`     
	IP        *string `json:"ip,omitempty"`        
	UserAgent *string `json:"user_agent,omitempty"`
	ISP       *string `json:"isp,omitempty"`       
	Country   *string `json:"country,omitempty"`   
	Region    *string `json:"region,omitempty"`    
}

type Paging struct {
	Count *int64 `json:"count,omitempty"`
	Total *int64 `json:"total,omitempty"`
	Page  *int64 `json:"page,omitempty"` 
	Pages *int64 `json:"pages,omitempty"`
}
