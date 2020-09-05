// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    oauthAccess, err := UnmarshalOauthAccess(bytes)
//    bytes, err = oauthAccess.Marshal()

package oauth_access

import "encoding/json"

func UnmarshalOauthAccess(data []byte) (OauthAccess, error) {
	var r OauthAccess
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *OauthAccess) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type OauthAccess struct {
	Ok              *bool            `json:"ok,omitempty"`              
	Warning         *string          `json:"warning,omitempty"`         
	Error           *string          `json:"error,omitempty"`           
	Needed          *string          `json:"needed,omitempty"`          
	Provided        *string          `json:"provided,omitempty"`        
	TokenType       *string          `json:"token_type,omitempty"`      
	AccessToken     *string          `json:"access_token,omitempty"`    
	Scope           *string          `json:"scope,omitempty"`           
	EnterpriseID    *string          `json:"enterprise_id,omitempty"`   
	TeamName        *string          `json:"team_name,omitempty"`       
	TeamID          *string          `json:"team_id,omitempty"`         
	UserID          *string          `json:"user_id,omitempty"`         
	IncomingWebhook *IncomingWebhook `json:"incoming_webhook,omitempty"`
	Bot             *Bot             `json:"bot,omitempty"`             
	AuthorizingUser *User            `json:"authorizing_user,omitempty"`
	InstallerUser   *User            `json:"installer_user,omitempty"`  
	Scopes          *Scopes          `json:"scopes,omitempty"`          
}

type User struct {
	UserID  *string `json:"user_id,omitempty"` 
	AppHome *string `json:"app_home,omitempty"`
}

type Bot struct {
	BotUserID      *string `json:"bot_user_id,omitempty"`     
	BotAccessToken *string `json:"bot_access_token,omitempty"`
}

type IncomingWebhook struct {
	URL              *string `json:"url,omitempty"`              
	Channel          *string `json:"channel,omitempty"`          
	ChannelID        *string `json:"channel_id,omitempty"`       
	ConfigurationURL *string `json:"configuration_url,omitempty"`
}

type Scopes struct {
	AppHome []string `json:"app_home,omitempty"`
	Team    []string `json:"team,omitempty"`    
	Channel []string `json:"channel,omitempty"` 
	Group   []string `json:"group,omitempty"`   
	Mpim    []string `json:"mpim,omitempty"`    
	IM      []string `json:"im,omitempty"`      
	User    []string `json:"user,omitempty"`    
}
