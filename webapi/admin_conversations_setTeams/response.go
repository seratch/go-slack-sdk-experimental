// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminConversationsSetTeams, err := UnmarshalAdminConversationsSetTeams(bytes)
//    bytes, err = adminConversationsSetTeams.Marshal()

package admin_conversations_setTeams

import "encoding/json"

func UnmarshalAdminConversationsSetTeams(data []byte) (AdminConversationsSetTeams, error) {
	var r AdminConversationsSetTeams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminConversationsSetTeams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminConversationsSetTeams struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Channel  *string `json:"channel,omitempty"` 
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
