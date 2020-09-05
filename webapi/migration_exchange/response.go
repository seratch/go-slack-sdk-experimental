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
	Error          *string    `json:"error,omitempty"`           
	TeamID         *string    `json:"team_id,omitempty"`         
	EnterpriseID   *string    `json:"enterprise_id,omitempty"`   
	UserIDMap      *UserIDMap `json:"user_id_map,omitempty"`     
	InvalidUserIDS []string   `json:"invalid_user_ids,omitempty"`
}

type UserIDMap struct {
	U06Ubsun5 *string `json:"U06UBSUN5,omitempty"`
	U06Ueb62U *string `json:"U06UEB62U,omitempty"`
	U06Ubsvb3 *string `json:"U06UBSVB3,omitempty"`
	U06Ubsvdx *string `json:"U06UBSVDX,omitempty"`
	W06Uaz65Q *string `json:"W06UAZ65Q,omitempty"`
}
