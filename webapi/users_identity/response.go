// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    usersIdentity, err := UnmarshalUsersIdentity(bytes)
//    bytes, err = usersIdentity.Marshal()

package users_identity

import "encoding/json"

func UnmarshalUsersIdentity(data []byte) (UsersIdentity, error) {
	var r UsersIdentity
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UsersIdentity) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type UsersIdentity struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
