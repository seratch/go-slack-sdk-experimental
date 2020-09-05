// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    usergroupsUsersList, err := UnmarshalUsergroupsUsersList(bytes)
//    bytes, err = usergroupsUsersList.Marshal()

package usergroups_users_list

import "encoding/json"

func UnmarshalUsergroupsUsersList(data []byte) (UsergroupsUsersList, error) {
	var r UsergroupsUsersList
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UsergroupsUsersList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type UsergroupsUsersList struct {
	Ok       *bool    `json:"ok,omitempty"`      
	Users    []string `json:"users,omitempty"`   
	Error    *string  `json:"error,omitempty"`   
	Needed   *string  `json:"needed,omitempty"`  
	Provided *string  `json:"provided,omitempty"`
}
