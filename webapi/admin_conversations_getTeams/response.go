// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminConversationsGetTeams, err := UnmarshalAdminConversationsGetTeams(bytes)
//    bytes, err = adminConversationsGetTeams.Marshal()

package admin_conversations_getTeams

import "encoding/json"

func UnmarshalAdminConversationsGetTeams(data []byte) (AdminConversationsGetTeams, error) {
	var r AdminConversationsGetTeams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminConversationsGetTeams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminConversationsGetTeams struct {
	Ok       *bool    `json:"ok,omitempty"`      
	TeamIDS  []string `json:"team_ids,omitempty"`
	Error    *string  `json:"error,omitempty"`   
	Needed   *string  `json:"needed,omitempty"`  
	Provided *string  `json:"provided,omitempty"`
}
