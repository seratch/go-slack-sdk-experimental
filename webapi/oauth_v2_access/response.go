// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    oauthV2Access, err := UnmarshalOauthV2Access(bytes)
//    bytes, err = oauthV2Access.Marshal()

package oauth_v2_access

import "encoding/json"

func UnmarshalOauthV2Access(data []byte) (OauthV2Access, error) {
	var r OauthV2Access
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *OauthV2Access) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type OauthV2Access struct {
	Ok                  *bool            `json:"ok,omitempty"`                   
	Warning             *string          `json:"warning,omitempty"`              
	Error               *string          `json:"error,omitempty"`                
	Needed              *string          `json:"needed,omitempty"`               
	Provided            *string          `json:"provided,omitempty"`             
	AppID               *string          `json:"app_id,omitempty"`               
	AuthedUser          *AuthedUser      `json:"authed_user,omitempty"`          
	Scope               *string          `json:"scope,omitempty"`                
	TokenType           *string          `json:"token_type,omitempty"`           
	AccessToken         *string          `json:"access_token,omitempty"`         
	RefreshToken        *string          `json:"refresh_token,omitempty"`        
	ExpiresIn           *int64           `json:"expires_in,omitempty"`           
	BotUserID           *string          `json:"bot_user_id,omitempty"`          
	Team                *Enterprise      `json:"team,omitempty"`                 
	Enterprise          *Enterprise      `json:"enterprise,omitempty"`           
	IsEnterpriseInstall *bool            `json:"is_enterprise_install,omitempty"`
	IncomingWebhook     *IncomingWebhook `json:"incoming_webhook,omitempty"`     
}

type AuthedUser struct {
	ID           *string `json:"id,omitempty"`           
	Scope        *string `json:"scope,omitempty"`        
	TokenType    *string `json:"token_type,omitempty"`   
	AccessToken  *string `json:"access_token,omitempty"` 
	RefreshToken *string `json:"refresh_token,omitempty"`
	ExpiresIn    *int64  `json:"expires_in,omitempty"`   
}

type Enterprise struct {
	ID   *string `json:"id,omitempty"`  
	Name *string `json:"name,omitempty"`
}

type IncomingWebhook struct {
	URL              *string `json:"url,omitempty"`              
	Channel          *string `json:"channel,omitempty"`          
	ChannelID        *string `json:"channel_id,omitempty"`       
	ConfigurationURL *string `json:"configuration_url,omitempty"`
}
