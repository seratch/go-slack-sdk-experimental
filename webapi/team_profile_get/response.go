// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    teamProfileGet, err := UnmarshalTeamProfileGet(bytes)
//    bytes, err = teamProfileGet.Marshal()

package team_profile_get

import "encoding/json"

func UnmarshalTeamProfileGet(data []byte) (TeamProfileGet, error) {
	var r TeamProfileGet
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TeamProfileGet) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TeamProfileGet struct {
	Ok       *bool    `json:"ok,omitempty"`      
	Profile  *Profile `json:"profile,omitempty"` 
	Error    *string  `json:"error,omitempty"`   
	Needed   *string  `json:"needed,omitempty"`  
	Provided *string  `json:"provided,omitempty"`
}

type Profile struct {
	Fields []Field `json:"fields,omitempty"`
}

type Field struct {
	ID        *string `json:"id,omitempty"`        
	Ordering  *int64  `json:"ordering,omitempty"`  
	FieldName *string `json:"field_name,omitempty"`
	Label     *string `json:"label,omitempty"`     
	Hint      *string `json:"hint,omitempty"`      
	Type      *string `json:"type,omitempty"`      
	IsHidden  *bool   `json:"is_hidden,omitempty"` 
}
