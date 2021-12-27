// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    migrationExchange, err := UnmarshalMigrationExchange(bytes)
//    bytes, err = migrationExchange.Marshal()

package migration_exchange

import "encoding/json"

func UnmarshalMigrationExchange(data []byte) (MigrationExchange, error) {
	var r MigrationExchange
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MigrationExchange) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MigrationExchange struct {
	Ok             *bool      `json:"ok,omitempty"`              
	Warning        *string    `json:"warning,omitempty"`         
	Error          *string    `json:"error,omitempty"`           
	Needed         *string    `json:"needed,omitempty"`          
	Provided       *string    `json:"provided,omitempty"`        
	TeamID         *string    `json:"team_id,omitempty"`         
	EnterpriseID   *string    `json:"enterprise_id,omitempty"`   
	InvalidUserIDS []string   `json:"invalid_user_ids,omitempty"`
	UserIDMap      *UserIDMap `json:"user_id_map,omitempty"`     
}

type UserIDMap struct {
}
