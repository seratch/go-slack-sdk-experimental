// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminConversationsGetCustomRetention, err := UnmarshalAdminConversationsGetCustomRetention(bytes)
//    bytes, err = adminConversationsGetCustomRetention.Marshal()

package admin_conversations_getCustomRetention

import "encoding/json"

func UnmarshalAdminConversationsGetCustomRetention(data []byte) (AdminConversationsGetCustomRetention, error) {
	var r AdminConversationsGetCustomRetention
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminConversationsGetCustomRetention) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminConversationsGetCustomRetention struct {
	Ok              *bool   `json:"ok,omitempty"`               
	Error           *string `json:"error,omitempty"`            
	Needed          *string `json:"needed,omitempty"`           
	Provided        *string `json:"provided,omitempty"`         
	IsPolicyEnabled *bool   `json:"is_policy_enabled,omitempty"`
	DurationDays    *int64  `json:"duration_days,omitempty"`    
}
