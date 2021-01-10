// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminBarriersCreate, err := UnmarshalAdminBarriersCreate(bytes)
//    bytes, err = adminBarriersCreate.Marshal()

package admin_barriers_create

import "encoding/json"

func UnmarshalAdminBarriersCreate(data []byte) (AdminBarriersCreate, error) {
	var r AdminBarriersCreate
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminBarriersCreate) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminBarriersCreate struct {
	Ok               *bool             `json:"ok,omitempty"`               
	Error            *string           `json:"error,omitempty"`            
	Needed           *string           `json:"needed,omitempty"`           
	Provided         *string           `json:"provided,omitempty"`         
	ResponseMetadata *ResponseMetadata `json:"response_metadata,omitempty"`
	Barrier          *Barrier          `json:"barrier,omitempty"`          
}

type Barrier struct {
	ID                      *string     `json:"id,omitempty"`                       
	EnterpriseID            *string     `json:"enterprise_id,omitempty"`            
	PrimaryUsergroup        *Usergroup  `json:"primary_usergroup,omitempty"`        
	BarrieredFromUsergroups []Usergroup `json:"barriered_from_usergroups,omitempty"`
	RestrictedSubjects      []string    `json:"restricted_subjects,omitempty"`      
	DateUpdate              *int64      `json:"date_update,omitempty"`              
}

type Usergroup struct {
	ID   *string `json:"id,omitempty"`  
	Name *string `json:"name,omitempty"`
}

type ResponseMetadata struct {
	Messages []string `json:"messages,omitempty"`
}
