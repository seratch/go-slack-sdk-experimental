// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    usersGetPresence, err := UnmarshalUsersGetPresence(bytes)
//    bytes, err = usersGetPresence.Marshal()

package users_getPresence

import "encoding/json"

func UnmarshalUsersGetPresence(data []byte) (UsersGetPresence, error) {
	var r UsersGetPresence
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UsersGetPresence) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type UsersGetPresence struct {
	Ok              *bool   `json:"ok,omitempty"`              
	Warning         *string `json:"warning,omitempty"`         
	Error           *string `json:"error,omitempty"`           
	Needed          *string `json:"needed,omitempty"`          
	Provided        *string `json:"provided,omitempty"`        
	Presence        *string `json:"presence,omitempty"`        
	Online          *bool   `json:"online,omitempty"`          
	AutoAway        *bool   `json:"auto_away,omitempty"`       
	ManualAway      *bool   `json:"manual_away,omitempty"`     
	ConnectionCount *int64  `json:"connection_count,omitempty"`
	LastActivity    *int64  `json:"last_activity,omitempty"`   
}
