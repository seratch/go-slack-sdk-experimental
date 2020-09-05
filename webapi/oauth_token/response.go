// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    oauthToken, err := UnmarshalOauthToken(bytes)
//    bytes, err = oauthToken.Marshal()

package oauth_token

import "encoding/json"

func UnmarshalOauthToken(data []byte) (OauthToken, error) {
	var r OauthToken
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *OauthToken) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type OauthToken struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
