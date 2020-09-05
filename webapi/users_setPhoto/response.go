// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    usersSetPhoto, err := UnmarshalUsersSetPhoto(bytes)
//    bytes, err = usersSetPhoto.Marshal()

package users_setPhoto

import "encoding/json"

func UnmarshalUsersSetPhoto(data []byte) (UsersSetPhoto, error) {
	var r UsersSetPhoto
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UsersSetPhoto) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type UsersSetPhoto struct {
	Ok       *bool    `json:"ok,omitempty"`      
	Profile  *Profile `json:"profile,omitempty"` 
	Error    *string  `json:"error,omitempty"`   
	Needed   *string  `json:"needed,omitempty"`  
	Provided *string  `json:"provided,omitempty"`
}

type Profile struct {
	Image24       *string `json:"image_24,omitempty"`      
	Image32       *string `json:"image_32,omitempty"`      
	Image48       *string `json:"image_48,omitempty"`      
	Image72       *string `json:"image_72,omitempty"`      
	Image192      *string `json:"image_192,omitempty"`     
	Image512      *string `json:"image_512,omitempty"`     
	Image1024     *string `json:"image_1024,omitempty"`    
	ImageOriginal *string `json:"image_original,omitempty"`
	AvatarHash    *string `json:"avatar_hash,omitempty"`   
}
