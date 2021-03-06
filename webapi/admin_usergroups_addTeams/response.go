// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminUsergroupsAddTeams, err := UnmarshalAdminUsergroupsAddTeams(bytes)
//    bytes, err = adminUsergroupsAddTeams.Marshal()

package admin_usergroups_addTeams

import "encoding/json"

func UnmarshalAdminUsergroupsAddTeams(data []byte) (AdminUsergroupsAddTeams, error) {
	var r AdminUsergroupsAddTeams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminUsergroupsAddTeams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminUsergroupsAddTeams struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
