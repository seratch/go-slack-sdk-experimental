// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminAuthPolicyRemoveEntities, err := UnmarshalAdminAuthPolicyRemoveEntities(bytes)
//    bytes, err = adminAuthPolicyRemoveEntities.Marshal()

package admin_auth_policy_removeEntities

import "encoding/json"

func UnmarshalAdminAuthPolicyRemoveEntities(data []byte) (AdminAuthPolicyRemoveEntities, error) {
	var r AdminAuthPolicyRemoveEntities
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminAuthPolicyRemoveEntities) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminAuthPolicyRemoveEntities struct {
	Ok               *bool   `json:"ok,omitempty"`                
	EntityTotalCount *int64  `json:"entity_total_count,omitempty"`
	Error            *string `json:"error,omitempty"`             
	Needed           *string `json:"needed,omitempty"`            
	Provided         *string `json:"provided,omitempty"`          
}
