// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminUsersList, err := UnmarshalAdminUsersList(bytes)
//    bytes, err = adminUsersList.Marshal()

package admin_users_list

import "encoding/json"

func UnmarshalAdminUsersList(data []byte) (AdminUsersList, error) {
	var r AdminUsersList
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminUsersList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminUsersList struct {
	Ok               *bool             `json:"ok,omitempty"`               
	Users            []User            `json:"users,omitempty"`            
	ResponseMetadata *ResponseMetadata `json:"response_metadata,omitempty"`
	Error            *string           `json:"error,omitempty"`            
	Needed           *string           `json:"needed,omitempty"`           
	Provided         *string           `json:"provided,omitempty"`         
}

type ResponseMetadata struct {
	NextCursor *string  `json:"next_cursor,omitempty"`
	Messages   []string `json:"messages,omitempty"`   
}

type User struct {
	ID                *string `json:"id,omitempty"`                 
	Email             *string `json:"email,omitempty"`              
	IsAdmin           *bool   `json:"is_admin,omitempty"`           
	IsOwner           *bool   `json:"is_owner,omitempty"`           
	IsPrimaryOwner    *bool   `json:"is_primary_owner,omitempty"`   
	IsRestricted      *bool   `json:"is_restricted,omitempty"`      
	IsUltraRestricted *bool   `json:"is_ultra_restricted,omitempty"`
	IsBot             *bool   `json:"is_bot,omitempty"`             
	ExpirationTs      *int64  `json:"expiration_ts,omitempty"`      
}
