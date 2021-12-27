// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminAuthPolicyAssignEntities, err := UnmarshalAdminAuthPolicyAssignEntities(bytes)
//    bytes, err = adminAuthPolicyAssignEntities.Marshal()

package admin_auth_policy_assignEntities

import "encoding/json"

func UnmarshalAdminAuthPolicyAssignEntities(data []byte) (AdminAuthPolicyAssignEntities, error) {
	var r AdminAuthPolicyAssignEntities
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminAuthPolicyAssignEntities) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminAuthPolicyAssignEntities struct {
	Ok               *bool   `json:"ok,omitempty"`                
	EntityTotalCount *int64  `json:"entity_total_count,omitempty"`
	Error            *string `json:"error,omitempty"`             
	Needed           *string `json:"needed,omitempty"`            
	Provided         *string `json:"provided,omitempty"`          
}
