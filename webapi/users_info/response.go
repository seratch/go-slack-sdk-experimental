// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    usersInfo, err := UnmarshalUsersInfo(bytes)
//    bytes, err = usersInfo.Marshal()

package users_info

import "encoding/json"

func UnmarshalUsersInfo(data []byte) (UsersInfo, error) {
	var r UsersInfo
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UsersInfo) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type UsersInfo struct {
	Ok       *bool   `json:"ok,omitempty"`      
	User     *User   `json:"user,omitempty"`    
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}

type User struct {
	ID                     *string         `json:"id,omitempty"`                        
	TeamID                 *string         `json:"team_id,omitempty"`                   
	Name                   *string         `json:"name,omitempty"`                      
	Deleted                *bool           `json:"deleted,omitempty"`                   
	Color                  *string         `json:"color,omitempty"`                     
	RealName               *string         `json:"real_name,omitempty"`                 
	Tz                     *string         `json:"tz,omitempty"`                        
	TzLabel                *string         `json:"tz_label,omitempty"`                  
	TzOffset               *int64          `json:"tz_offset,omitempty"`                 
	Profile                *Profile        `json:"profile,omitempty"`                   
	IsAdmin                *bool           `json:"is_admin,omitempty"`                  
	IsOwner                *bool           `json:"is_owner,omitempty"`                  
	IsPrimaryOwner         *bool           `json:"is_primary_owner,omitempty"`          
	IsRestricted           *bool           `json:"is_restricted,omitempty"`             
	IsUltraRestricted      *bool           `json:"is_ultra_restricted,omitempty"`       
	IsBot                  *bool           `json:"is_bot,omitempty"`                    
	IsAppUser              *bool           `json:"is_app_user,omitempty"`               
	Updated                *int64          `json:"updated,omitempty"`                   
	Locale                 *string         `json:"locale,omitempty"`                    
	Has2Fa                 *bool           `json:"has_2fa,omitempty"`                   
	IsEmailConfirmed       *bool           `json:"is_email_confirmed,omitempty"`        
	EnterpriseUser         *EnterpriseUser `json:"enterprise_user,omitempty"`           
	IsStranger             *bool           `json:"is_stranger,omitempty"`               
	WhoCanShareContactCard *string         `json:"who_can_share_contact_card,omitempty"`
}

type EnterpriseUser struct {
	ID             *string  `json:"id,omitempty"`             
	EnterpriseID   *string  `json:"enterprise_id,omitempty"`  
	EnterpriseName *string  `json:"enterprise_name,omitempty"`
	IsAdmin        *bool    `json:"is_admin,omitempty"`       
	IsOwner        *bool    `json:"is_owner,omitempty"`       
	Teams          []string `json:"teams,omitempty"`          
}

type Profile struct {
	Title                  *string                  `json:"title,omitempty"`                    
	Phone                  *string                  `json:"phone,omitempty"`                    
	Skype                  *string                  `json:"skype,omitempty"`                    
	RealName               *string                  `json:"real_name,omitempty"`                
	RealNameNormalized     *string                  `json:"real_name_normalized,omitempty"`     
	DisplayName            *string                  `json:"display_name,omitempty"`             
	DisplayNameNormalized  *string                  `json:"display_name_normalized,omitempty"`  
	StatusText             *string                  `json:"status_text,omitempty"`              
	StatusEmoji            *string                  `json:"status_emoji,omitempty"`             
	StatusExpiration       *int64                   `json:"status_expiration,omitempty"`        
	AvatarHash             *string                  `json:"avatar_hash,omitempty"`              
	APIAppID               *string                  `json:"api_app_id,omitempty"`               
	AlwaysActive           *bool                    `json:"always_active,omitempty"`            
	BotID                  *string                  `json:"bot_id,omitempty"`                   
	Image24                *string                  `json:"image_24,omitempty"`                 
	Image32                *string                  `json:"image_32,omitempty"`                 
	Image48                *string                  `json:"image_48,omitempty"`                 
	Image72                *string                  `json:"image_72,omitempty"`                 
	Image192               *string                  `json:"image_192,omitempty"`                
	Image512               *string                  `json:"image_512,omitempty"`                
	StatusTextCanonical    *string                  `json:"status_text_canonical,omitempty"`    
	Team                   *string                  `json:"team,omitempty"`                     
	ImageOriginal          *string                  `json:"image_original,omitempty"`           
	IsCustomImage          *bool                    `json:"is_custom_image,omitempty"`          
	Email                  *string                  `json:"email,omitempty"`                    
	FirstName              *string                  `json:"first_name,omitempty"`               
	LastName               *string                  `json:"last_name,omitempty"`                
	Image1024              *string                  `json:"image_1024,omitempty"`               
	StatusEmojiURL         *string                  `json:"status_emoji_url,omitempty"`         
	Pronouns               *string                  `json:"pronouns,omitempty"`                 
	StatusEmojiDisplayInfo []StatusEmojiDisplayInfo `json:"status_emoji_display_info,omitempty"`
}

type StatusEmojiDisplayInfo struct {
	EmojiName    *string `json:"emoji_name,omitempty"`   
	DisplayAlias *string `json:"display_alias,omitempty"`
	DisplayURL   *string `json:"display_url,omitempty"`  
}
