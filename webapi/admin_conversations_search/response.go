// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminConversationsSearch, err := UnmarshalAdminConversationsSearch(bytes)
//    bytes, err = adminConversationsSearch.Marshal()

package admin_conversations_search

import "encoding/json"

func UnmarshalAdminConversationsSearch(data []byte) (AdminConversationsSearch, error) {
	var r AdminConversationsSearch
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminConversationsSearch) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminConversationsSearch struct {
	Ok            *bool          `json:"ok,omitempty"`           
	Conversations []Conversation `json:"conversations,omitempty"`
	NextCursor    *string        `json:"next_cursor,omitempty"`  
	Error         *string        `json:"error,omitempty"`        
	Needed        *string        `json:"needed,omitempty"`       
	Provided      *string        `json:"provided,omitempty"`     
}

type Conversation struct {
	ID                        *string  `json:"id,omitempty"`                           
	Name                      *string  `json:"name,omitempty"`                         
	Purpose                   *string  `json:"purpose,omitempty"`                      
	MemberCount               *int64   `json:"member_count,omitempty"`                 
	Created                   *int64   `json:"created,omitempty"`                      
	CreatorID                 *string  `json:"creator_id,omitempty"`                   
	IsPrivate                 *bool    `json:"is_private,omitempty"`                   
	IsArchived                *bool    `json:"is_archived,omitempty"`                  
	IsGeneral                 *bool    `json:"is_general,omitempty"`                   
	LastActivityTs            *int64   `json:"last_activity_ts,omitempty"`             
	IsEXTShared               *bool    `json:"is_ext_shared,omitempty"`                
	IsGlobalShared            *bool    `json:"is_global_shared,omitempty"`             
	IsOrgDefault              *bool    `json:"is_org_default,omitempty"`               
	IsOrgMandatory            *bool    `json:"is_org_mandatory,omitempty"`             
	IsOrgShared               *bool    `json:"is_org_shared,omitempty"`                
	IsFrozen                  *bool    `json:"is_frozen,omitempty"`                    
	InternalTeamIDSCount      *int64   `json:"internal_team_ids_count,omitempty"`      
	InternalTeamIDSSampleTeam *string  `json:"internal_team_ids_sample_team,omitempty"`
	PendingConnectedTeamIDS   []string `json:"pending_connected_team_ids,omitempty"`   
	IsPendingEXTShared        *bool    `json:"is_pending_ext_shared,omitempty"`        
	ConnectedTeamIDS          []string `json:"connected_team_ids,omitempty"`           
	ConversationHostID        *string  `json:"conversation_host_id,omitempty"`         
	ChannelEmailAddresses     []string `json:"channel_email_addresses,omitempty"`      
}
