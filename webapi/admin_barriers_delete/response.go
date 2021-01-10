// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminBarriersDelete, err := UnmarshalAdminBarriersDelete(bytes)
//    bytes, err = adminBarriersDelete.Marshal()

package admin_barriers_delete

import "encoding/json"

func UnmarshalAdminBarriersDelete(data []byte) (AdminBarriersDelete, error) {
	var r AdminBarriersDelete
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminBarriersDelete) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminBarriersDelete struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
