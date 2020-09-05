// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    usersSetActive, err := UnmarshalUsersSetActive(bytes)
//    bytes, err = usersSetActive.Marshal()

package users_setActive

import "encoding/json"

func UnmarshalUsersSetActive(data []byte) (UsersSetActive, error) {
	var r UsersSetActive
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UsersSetActive) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type UsersSetActive struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
