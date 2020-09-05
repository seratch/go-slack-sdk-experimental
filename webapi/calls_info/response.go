// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    callsInfo, err := UnmarshalCallsInfo(bytes)
//    bytes, err = callsInfo.Marshal()

package calls_info

import "encoding/json"

func UnmarshalCallsInfo(data []byte) (CallsInfo, error) {
	var r CallsInfo
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CallsInfo) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CallsInfo struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Call     *Call   `json:"call,omitempty"`    
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}

type Call struct {
	ID                *string  `json:"id,omitempty"`                  
	DateStart         *int64   `json:"date_start,omitempty"`          
	ExternalUniqueID  *string  `json:"external_unique_id,omitempty"`  
	JoinURL           *string  `json:"join_url,omitempty"`            
	Channels          []string `json:"channels,omitempty"`            
	ExternalDisplayID *string  `json:"external_display_id,omitempty"` 
	Title             *string  `json:"title,omitempty"`               
	DesktopAppJoinURL *string  `json:"desktop_app_join_url,omitempty"`
	Users             []User   `json:"users,omitempty"`               
}

type User struct {
	ExternalID  *string `json:"external_id,omitempty"` 
	AvatarURL   *string `json:"avatar_url,omitempty"`  
	DisplayName *string `json:"display_name,omitempty"`
	SlackID     *string `json:"slack_id,omitempty"`    
}
