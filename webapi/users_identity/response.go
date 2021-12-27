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
	Warning  *string `json:"warning,omitempty"` 
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
	User     *User   `json:"user,omitempty"`    
	Team     *Team   `json:"team,omitempty"`    
}

type Team struct {
	Name *string `json:"name,omitempty"`
	ID   *string `json:"id,omitempty"`  
}

type User struct {
	Name     *string `json:"name,omitempty"`     
	ID       *string `json:"id,omitempty"`       
	Email    *string `json:"email,omitempty"`    
	Image24  *string `json:"image_24,omitempty"` 
	Image32  *string `json:"image_32,omitempty"` 
	Image48  *string `json:"image_48,omitempty"` 
	Image72  *string `json:"image_72,omitempty"` 
	Image192 *string `json:"image_192,omitempty"`
	Image512 *string `json:"image_512,omitempty"`
}
