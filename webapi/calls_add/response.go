// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    callsAdd, err := UnmarshalCallsAdd(bytes)
//    bytes, err = callsAdd.Marshal()

package calls_add

import "encoding/json"

func UnmarshalCallsAdd(data []byte) (CallsAdd, error) {
	var r CallsAdd
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CallsAdd) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CallsAdd struct {
	Ok               *bool             `json:"ok,omitempty"`               
	Call             *Call             `json:"call,omitempty"`             
	Error            *string           `json:"error,omitempty"`            
	ResponseMetadata *ResponseMetadata `json:"response_metadata,omitempty"`
	Needed           *string           `json:"needed,omitempty"`           
	Provided         *string           `json:"provided,omitempty"`         
}

type Call struct {
	ID                *string `json:"id,omitempty"`                  
	DateStart         *int64  `json:"date_start,omitempty"`          
	ExternalUniqueID  *string `json:"external_unique_id,omitempty"`  
	JoinURL           *string `json:"join_url,omitempty"`            
	ExternalDisplayID *string `json:"external_display_id,omitempty"` 
	Title             *string `json:"title,omitempty"`               
	Users             []User  `json:"users,omitempty"`               
	DesktopAppJoinURL *string `json:"desktop_app_join_url,omitempty"`
}

type User struct {
	SlackID     *string `json:"slack_id,omitempty"`    
	ExternalID  *string `json:"external_id,omitempty"` 
	DisplayName *string `json:"display_name,omitempty"`
	AvatarURL   *string `json:"avatar_url,omitempty"`  
}

type ResponseMetadata struct {
	Messages []string `json:"messages,omitempty"`
}
