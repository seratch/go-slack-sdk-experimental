// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminUsergroupsRemoveChannels, err := UnmarshalAdminUsergroupsRemoveChannels(bytes)
//    bytes, err = adminUsergroupsRemoveChannels.Marshal()

package admin_usergroups_removeChannels

import "encoding/json"

func UnmarshalAdminUsergroupsRemoveChannels(data []byte) (AdminUsergroupsRemoveChannels, error) {
	var r AdminUsergroupsRemoveChannels
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminUsergroupsRemoveChannels) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminUsergroupsRemoveChannels struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
