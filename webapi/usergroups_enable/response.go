// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    usergroupsEnable, err := UnmarshalUsergroupsEnable(bytes)
//    bytes, err = usergroupsEnable.Marshal()

package usergroups_enable

import "encoding/json"

func UnmarshalUsergroupsEnable(data []byte) (UsergroupsEnable, error) {
	var r UsergroupsEnable
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UsergroupsEnable) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type UsergroupsEnable struct {
	Ok        *bool      `json:"ok,omitempty"`       
	Usergroup *Usergroup `json:"usergroup,omitempty"`
	Error     *string    `json:"error,omitempty"`    
	Needed    *string    `json:"needed,omitempty"`   
	Provided  *string    `json:"provided,omitempty"` 
}

type Usergroup struct {
	ID                  *string `json:"id,omitempty"`                   
	TeamID              *string `json:"team_id,omitempty"`              
	IsUsergroup         *bool   `json:"is_usergroup,omitempty"`         
	IsSubteam           *bool   `json:"is_subteam,omitempty"`           
	Name                *string `json:"name,omitempty"`                 
	Description         *string `json:"description,omitempty"`          
	Handle              *string `json:"handle,omitempty"`               
	IsExternal          *bool   `json:"is_external,omitempty"`          
	DateCreate          *int64  `json:"date_create,omitempty"`          
	DateUpdate          *int64  `json:"date_update,omitempty"`          
	DateDelete          *int64  `json:"date_delete,omitempty"`          
	AutoProvision       *bool   `json:"auto_provision,omitempty"`       
	EnterpriseSubteamID *string `json:"enterprise_subteam_id,omitempty"`
	CreatedBy           *string `json:"created_by,omitempty"`           
	UpdatedBy           *string `json:"updated_by,omitempty"`           
	Prefs               *Prefs  `json:"prefs,omitempty"`                
	ChannelCount        *int64  `json:"channel_count,omitempty"`        
}

type Prefs struct {
	Channels []string `json:"channels,omitempty"`
	Groups   []string `json:"groups,omitempty"`  
}
