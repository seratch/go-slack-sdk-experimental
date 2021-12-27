// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    openidConnectToken, err := UnmarshalOpenidConnectToken(bytes)
//    bytes, err = openidConnectToken.Marshal()

package openid_connect_token

import "encoding/json"

func UnmarshalOpenidConnectToken(data []byte) (OpenidConnectToken, error) {
	var r OpenidConnectToken
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *OpenidConnectToken) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type OpenidConnectToken struct {
	Ok           *bool   `json:"ok,omitempty"`           
	Warning      *string `json:"warning,omitempty"`      
	Error        *string `json:"error,omitempty"`        
	Needed       *string `json:"needed,omitempty"`       
	Provided     *string `json:"provided,omitempty"`     
	AccessToken  *string `json:"access_token,omitempty"` 
	TokenType    *string `json:"token_type,omitempty"`   
	IDToken      *string `json:"id_token,omitempty"`     
	RefreshToken *string `json:"refresh_token,omitempty"`
	ExpiresIn    *int64  `json:"expires_in,omitempty"`   
}
