// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminUsergroupsListChannels, err := UnmarshalAdminUsergroupsListChannels(bytes)
//    bytes, err = adminUsergroupsListChannels.Marshal()

package admin_usergroups_listChannels

import "encoding/json"

func UnmarshalAdminUsergroupsListChannels(data []byte) (AdminUsergroupsListChannels, error) {
	var r AdminUsergroupsListChannels
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminUsergroupsListChannels) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminUsergroupsListChannels struct {
	Ok       *bool     `json:"ok,omitempty"`      
	Channels []Channel `json:"channels,omitempty"`
	Error    *string   `json:"error,omitempty"`   
	Needed   *string   `json:"needed,omitempty"`  
	Provided *string   `json:"provided,omitempty"`
}

type Channel struct {
	ID                      *string  `json:"id,omitempty"`                        
	Name                    *string  `json:"name,omitempty"`                      
	IsChannel               *bool    `json:"is_channel,omitempty"`                
	IsGroup                 *bool    `json:"is_group,omitempty"`                  
	IsIM                    *bool    `json:"is_im,omitempty"`                     
	Created                 *int64   `json:"created,omitempty"`                   
	IsArchived              *bool    `json:"is_archived,omitempty"`               
	IsGeneral               *bool    `json:"is_general,omitempty"`                
	Unlinked                *int64   `json:"unlinked,omitempty"`                  
	NameNormalized          *string  `json:"name_normalized,omitempty"`           
	IsShared                *bool    `json:"is_shared,omitempty"`                 
	Creator                 *string  `json:"creator,omitempty"`                   
	IsMoved                 *int64   `json:"is_moved,omitempty"`                  
	IsEXTShared             *bool    `json:"is_ext_shared,omitempty"`             
	EnterpriseID            *string  `json:"enterprise_id,omitempty"`             
	IsGlobalShared          *bool    `json:"is_global_shared,omitempty"`          
	IsOrgDefault            *bool    `json:"is_org_default,omitempty"`            
	IsOrgMandatory          *bool    `json:"is_org_mandatory,omitempty"`          
	IsOrgShared             *bool    `json:"is_org_shared,omitempty"`             
	PendingShared           []string `json:"pending_shared,omitempty"`            
	PendingConnectedTeamIDS []string `json:"pending_connected_team_ids,omitempty"`
	IsPendingEXTShared      *bool    `json:"is_pending_ext_shared,omitempty"`     
	IsMember                *bool    `json:"is_member,omitempty"`                 
	IsPrivate               *bool    `json:"is_private,omitempty"`                
	IsMpim                  *bool    `json:"is_mpim,omitempty"`                   
	Topic                   *Purpose `json:"topic,omitempty"`                     
	Purpose                 *Purpose `json:"purpose,omitempty"`                   
	PreviousNames           []string `json:"previous_names,omitempty"`            
	DateConnected           *int64   `json:"date_connected,omitempty"`            
}

type Purpose struct {
	Value   *string `json:"value,omitempty"`   
	Creator *string `json:"creator,omitempty"` 
	LastSet *int64  `json:"last_set,omitempty"`
}
