// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    usersDeletePhoto, err := UnmarshalUsersDeletePhoto(bytes)
//    bytes, err = usersDeletePhoto.Marshal()

package users_deletePhoto

import "encoding/json"

func UnmarshalUsersDeletePhoto(data []byte) (UsersDeletePhoto, error) {
	var r UsersDeletePhoto
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UsersDeletePhoto) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type UsersDeletePhoto struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
