// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    teamIntegrationLogs, err := UnmarshalTeamIntegrationLogs(bytes)
//    bytes, err = teamIntegrationLogs.Marshal()

package team_integrationLogs

import "encoding/json"

func UnmarshalTeamIntegrationLogs(data []byte) (TeamIntegrationLogs, error) {
	var r TeamIntegrationLogs
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TeamIntegrationLogs) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TeamIntegrationLogs struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Logs     []Log   `json:"logs,omitempty"`    
	Paging   *Paging `json:"paging,omitempty"`  
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}

type Log struct {
	UserID            *string `json:"user_id,omitempty"`             
	UserName          *string `json:"user_name,omitempty"`           
	Date              *string `json:"date,omitempty"`                
	ChangeType        *string `json:"change_type,omitempty"`         
	AppType           *string `json:"app_type,omitempty"`            
	AppID             *string `json:"app_id,omitempty"`              
	Scope             *string `json:"scope,omitempty"`               
	RSSFeed           *bool   `json:"rss_feed,omitempty"`            
	RSSFeedChangeType *string `json:"rss_feed_change_type,omitempty"`
	RSSFeedTitle      *string `json:"rss_feed_title,omitempty"`      
	RSSFeedURL        *string `json:"rss_feed_url,omitempty"`        
	ServiceID         *int64  `json:"service_id,omitempty"`          
	ServiceType       *string `json:"service_type,omitempty"`        
	Channel           *string `json:"channel,omitempty"`             
	Reason            *string `json:"reason,omitempty"`              
	Resolution        *string `json:"resolution,omitempty"`          
}

type Paging struct {
	Count *int64 `json:"count,omitempty"`
	Total *int64 `json:"total,omitempty"`
	Page  *int64 `json:"page,omitempty"` 
	Pages *int64 `json:"pages,omitempty"`
}
