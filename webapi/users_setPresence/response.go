// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    usersSetPresence, err := UnmarshalUsersSetPresence(bytes)
//    bytes, err = usersSetPresence.Marshal()

package users_setPresence

import "encoding/json"

func UnmarshalUsersSetPresence(data []byte) (UsersSetPresence, error) {
	var r UsersSetPresence
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UsersSetPresence) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type UsersSetPresence struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
