// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminBarriersUpdate, err := UnmarshalAdminBarriersUpdate(bytes)
//    bytes, err = adminBarriersUpdate.Marshal()

package admin_barriers_update

import "encoding/json"

func UnmarshalAdminBarriersUpdate(data []byte) (AdminBarriersUpdate, error) {
	var r AdminBarriersUpdate
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminBarriersUpdate) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminBarriersUpdate struct {
	Ok       *bool    `json:"ok,omitempty"`      
	Error    *string  `json:"error,omitempty"`   
	Barrier  *Barrier `json:"barrier,omitempty"` 
	Needed   *string  `json:"needed,omitempty"`  
	Provided *string  `json:"provided,omitempty"`
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
