// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminAuthPolicyGetEntities, err := UnmarshalAdminAuthPolicyGetEntities(bytes)
//    bytes, err = adminAuthPolicyGetEntities.Marshal()

package admin_auth_policy_getEntities

import "encoding/json"

func UnmarshalAdminAuthPolicyGetEntities(data []byte) (AdminAuthPolicyGetEntities, error) {
	var r AdminAuthPolicyGetEntities
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminAuthPolicyGetEntities) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminAuthPolicyGetEntities struct {
	Ok               *bool    `json:"ok,omitempty"`                
	Error            *string  `json:"error,omitempty"`             
	Entities         []Entity `json:"entities,omitempty"`          
	EntityTotalCount *int64   `json:"entity_total_count,omitempty"`
	Needed           *string  `json:"needed,omitempty"`            
	Provided         *string  `json:"provided,omitempty"`          
}

type Entity struct {
	EntityID   *string `json:"entity_id,omitempty"`  
	EntityType *string `json:"entity_type,omitempty"`
	DateAdded  *int64  `json:"date_added,omitempty"` 
}
