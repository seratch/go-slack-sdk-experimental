// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    authTest, err := UnmarshalAuthTest(bytes)
//    bytes, err = authTest.Marshal()

package auth_test

import "encoding/json"

func UnmarshalAuthTest(data []byte) (AuthTest, error) {
	var r AuthTest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AuthTest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AuthTest struct {
	Ok           *bool   `json:"ok,omitempty"`           
	URL          *string `json:"url,omitempty"`          
	Team         *string `json:"team,omitempty"`         
	User         *string `json:"user,omitempty"`         
	TeamID       *string `json:"team_id,omitempty"`      
	UserID       *string `json:"user_id,omitempty"`      
	BotID        *string `json:"bot_id,omitempty"`       
	EnterpriseID *string `json:"enterprise_id,omitempty"`
	Error        *string `json:"error,omitempty"`        
	AppName      *string `json:"app_name,omitempty"`     
	AppID        *string `json:"app_id,omitempty"`       
	Needed       *string `json:"needed,omitempty"`       
	Provided     *string `json:"provided,omitempty"`     
}
