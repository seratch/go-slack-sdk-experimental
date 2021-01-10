// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminUsersSessionList, err := UnmarshalAdminUsersSessionList(bytes)
//    bytes, err = adminUsersSessionList.Marshal()

package admin_users_session_list

import "encoding/json"

func UnmarshalAdminUsersSessionList(data []byte) (AdminUsersSessionList, error) {
	var r AdminUsersSessionList
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminUsersSessionList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminUsersSessionList struct {
	Ok               *bool             `json:"ok,omitempty"`               
	Error            *string           `json:"error,omitempty"`            
	ActiveSessions   []ActiveSession   `json:"active_sessions,omitempty"`  
	ResponseMetadata *ResponseMetadata `json:"response_metadata,omitempty"`
	Needed           *string           `json:"needed,omitempty"`           
	Provided         *string           `json:"provided,omitempty"`         
}

type ActiveSession struct {
	UserID    *string  `json:"user_id,omitempty"`   
	SessionID *int64   `json:"session_id,omitempty"`
	TeamID    *string  `json:"team_id,omitempty"`   
	Created   *Created `json:"created,omitempty"`   
}

type Created struct {
	DeviceHardware     *string `json:"device_hardware,omitempty"`     
	OS                 *string `json:"os,omitempty"`                  
	OSVersion          *string `json:"os_version,omitempty"`          
	SlackClientVersion *string `json:"slack_client_version,omitempty"`
	IP                 *string `json:"ip,omitempty"`                  
}

type ResponseMetadata struct {
	NextCursor *string `json:"next_cursor,omitempty"`
}
