// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminUsergroupsAddChannels, err := UnmarshalAdminUsergroupsAddChannels(bytes)
//    bytes, err = adminUsergroupsAddChannels.Marshal()

package admin_usergroups_addChannels

import "encoding/json"

func UnmarshalAdminUsergroupsAddChannels(data []byte) (AdminUsergroupsAddChannels, error) {
	var r AdminUsergroupsAddChannels
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminUsergroupsAddChannels) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminUsergroupsAddChannels struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
