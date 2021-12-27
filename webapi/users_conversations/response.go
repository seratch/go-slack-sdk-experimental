// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    usersConversations, err := UnmarshalUsersConversations(bytes)
//    bytes, err = usersConversations.Marshal()

package users_conversations

import "encoding/json"

func UnmarshalUsersConversations(data []byte) (UsersConversations, error) {
	var r UsersConversations
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UsersConversations) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type UsersConversations struct {
	Ok               *bool             `json:"ok,omitempty"`               
	Channels         []Channel         `json:"channels,omitempty"`         
	ResponseMetadata *ResponseMetadata `json:"response_metadata,omitempty"`
	Error            *string           `json:"error,omitempty"`            
	Needed           *string           `json:"needed,omitempty"`           
	Provided         *string           `json:"provided,omitempty"`         
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
	IsEXTShared             *bool    `json:"is_ext_shared,omitempty"`             
	IsOrgShared             *bool    `json:"is_org_shared,omitempty"`             
	SharedTeamIDS           []string `json:"shared_team_ids,omitempty"`           
	PendingShared           []string `json:"pending_shared,omitempty"`            
	PendingConnectedTeamIDS []string `json:"pending_connected_team_ids,omitempty"`
	IsPendingEXTShared      *bool    `json:"is_pending_ext_shared,omitempty"`     
	IsPrivate               *bool    `json:"is_private,omitempty"`                
	IsMpim                  *bool    `json:"is_mpim,omitempty"`                   
	Topic                   *Purpose `json:"topic,omitempty"`                     
	Purpose                 *Purpose `json:"purpose,omitempty"`                   
	PreviousNames           []string `json:"previous_names,omitempty"`            
	ConversationHostID      *string  `json:"conversation_host_id,omitempty"`      
	IsMoved                 *int64   `json:"is_moved,omitempty"`                  
	InternalTeamIDS         []string `json:"internal_team_ids,omitempty"`         
	IsGlobalShared          *bool    `json:"is_global_shared,omitempty"`          
	IsOrgDefault            *bool    `json:"is_org_default,omitempty"`            
	IsOrgMandatory          *bool    `json:"is_org_mandatory,omitempty"`          
	EnterpriseID            *string  `json:"enterprise_id,omitempty"`             
	LastRead                *string  `json:"last_read,omitempty"`                 
	IsOpen                  *bool    `json:"is_open,omitempty"`                   
	Priority                *int64   `json:"priority,omitempty"`                  
	User                    *string  `json:"user,omitempty"`                      
	IsUserDeleted           *bool    `json:"is_user_deleted,omitempty"`           
}

type Purpose struct {
	Value   *string `json:"value,omitempty"`   
	Creator *string `json:"creator,omitempty"` 
	LastSet *int64  `json:"last_set,omitempty"`
}

type ResponseMetadata struct {
	NextCursor *string `json:"next_cursor,omitempty"`
}
