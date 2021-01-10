// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminBarriersList, err := UnmarshalAdminBarriersList(bytes)
//    bytes, err = adminBarriersList.Marshal()

package admin_barriers_list

import "encoding/json"

func UnmarshalAdminBarriersList(data []byte) (AdminBarriersList, error) {
	var r AdminBarriersList
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminBarriersList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminBarriersList struct {
	Ok       *bool     `json:"ok,omitempty"`      
	Barriers []Barrier `json:"barriers,omitempty"`
	Error    *string   `json:"error,omitempty"`   
	Needed   *string   `json:"needed,omitempty"`  
	Provided *string   `json:"provided,omitempty"`
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
