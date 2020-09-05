// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    authRevoke, err := UnmarshalAuthRevoke(bytes)
//    bytes, err = authRevoke.Marshal()

package auth_revoke

import "encoding/json"

func UnmarshalAuthRevoke(data []byte) (AuthRevoke, error) {
	var r AuthRevoke
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AuthRevoke) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AuthRevoke struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
