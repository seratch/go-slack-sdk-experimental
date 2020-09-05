// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    usersProfileSet, err := UnmarshalUsersProfileSet(bytes)
//    bytes, err = usersProfileSet.Marshal()

package users_profile_set

import "encoding/json"

func UnmarshalUsersProfileSet(data []byte) (UsersProfileSet, error) {
	var r UsersProfileSet
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UsersProfileSet) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type UsersProfileSet struct {
	Profile  *Profile `json:"profile,omitempty"` 
	Ok       *bool    `json:"ok,omitempty"`      
	Username *string  `json:"username,omitempty"`
	Error    *string  `json:"error,omitempty"`   
	Needed   *string  `json:"needed,omitempty"`  
	Provided *string  `json:"provided,omitempty"`
}

type Profile struct {
	Title                 *string          `json:"title,omitempty"`                  
	Phone                 *string          `json:"phone,omitempty"`                  
	Skype                 *string          `json:"skype,omitempty"`                  
	RealName              *string          `json:"real_name,omitempty"`              
	RealNameNormalized    *string          `json:"real_name_normalized,omitempty"`   
	DisplayName           *string          `json:"display_name,omitempty"`           
	DisplayNameNormalized *string          `json:"display_name_normalized,omitempty"`
	Fields                map[string]Field `json:"fields,omitempty"`                 
	StatusText            *string          `json:"status_text,omitempty"`            
	StatusEmoji           *string          `json:"status_emoji,omitempty"`           
	StatusExpiration      *int64           `json:"status_expiration,omitempty"`      
	AvatarHash            *string          `json:"avatar_hash,omitempty"`            
	ImageOriginal         *string          `json:"image_original,omitempty"`         
	IsCustomImage         *bool            `json:"is_custom_image,omitempty"`        
	Email                 *string          `json:"email,omitempty"`                  
	FirstName             *string          `json:"first_name,omitempty"`             
	LastName              *string          `json:"last_name,omitempty"`              
	Image24               *string          `json:"image_24,omitempty"`               
	Image32               *string          `json:"image_32,omitempty"`               
	Image48               *string          `json:"image_48,omitempty"`               
	Image72               *string          `json:"image_72,omitempty"`               
	Image192              *string          `json:"image_192,omitempty"`              
	Image512              *string          `json:"image_512,omitempty"`              
	Image1024             *string          `json:"image_1024,omitempty"`             
	StatusTextCanonical   *string          `json:"status_text_canonical,omitempty"`  
}

type Field struct {
	Value *string `json:"value,omitempty"`
	Alt   *string `json:"alt,omitempty"`  
}
