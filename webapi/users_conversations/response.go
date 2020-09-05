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
	EnterpriseID       *string  `json:"enterprise_id,omitempty"`        
	ID                 *string  `json:"id,omitempty"`                   
	Name               *string  `json:"name,omitempty"`                 
	Created            *string  `json:"created,omitempty"`              
	Creator            *string  `json:"creator,omitempty"`              
	Unlinked           *int64   `json:"unlinked,omitempty"`             
	NameNormalized     *string  `json:"name_normalized,omitempty"`      
	LastRead           *string  `json:"last_read,omitempty"`            
	Topic              *Purpose `json:"topic,omitempty"`                
	Purpose            *Purpose `json:"purpose,omitempty"`              
	NumMembers         *int64   `json:"num_members,omitempty"`          
	Latest             *Latest  `json:"latest,omitempty"`               
	Locale             *string  `json:"locale,omitempty"`               
	UnreadCount        *int64   `json:"unread_count,omitempty"`         
	UnreadCountDisplay *int64   `json:"unread_count_display,omitempty"` 
	User               *string  `json:"user,omitempty"`                 
	Priority           *float64 `json:"priority,omitempty"`             
	DateConnected      *int64   `json:"date_connected,omitempty"`       
	ParentConversation *string  `json:"parent_conversation,omitempty"`  
	ConversationHostID *string  `json:"conversation_host_id,omitempty"` 
	IsChannel          *bool    `json:"is_channel,omitempty"`           
	IsGroup            *bool    `json:"is_group,omitempty"`             
	IsIM               *bool    `json:"is_im,omitempty"`                
	IsArchived         *bool    `json:"is_archived,omitempty"`          
	IsGeneral          *bool    `json:"is_general,omitempty"`           
	IsReadOnly         *bool    `json:"is_read_only,omitempty"`         
	IsThreadOnly       *bool    `json:"is_thread_only,omitempty"`       
	IsNonThreadable    *bool    `json:"is_non_threadable,omitempty"`    
	IsShared           *bool    `json:"is_shared,omitempty"`            
	IsEXTShared        *bool    `json:"is_ext_shared,omitempty"`        
	IsOrgShared        *bool    `json:"is_org_shared,omitempty"`        
	IsPendingEXTShared *bool    `json:"is_pending_ext_shared,omitempty"`
	IsGlobalShared     *bool    `json:"is_global_shared,omitempty"`     
	IsOrgDefault       *bool    `json:"is_org_default,omitempty"`       
	IsOrgMandatory     *bool    `json:"is_org_mandatory,omitempty"`     
	IsMoved            *int64   `json:"is_moved,omitempty"`             
	IsMember           *bool    `json:"is_member,omitempty"`            
	IsOpen             *bool    `json:"is_open,omitempty"`              
	IsPrivate          *bool    `json:"is_private,omitempty"`           
	IsMpim             *bool    `json:"is_mpim,omitempty"`              
}

type Latest struct {
	ClientMsgID  *string      `json:"client_msg_id,omitempty"` 
	Type         *string      `json:"type,omitempty"`          
	Subtype      *string      `json:"subtype,omitempty"`       
	Team         *string      `json:"team,omitempty"`          
	User         *string      `json:"user,omitempty"`          
	Username     *string      `json:"username,omitempty"`      
	ParentUserID *string      `json:"parent_user_id,omitempty"`
	Text         *string      `json:"text,omitempty"`          
	Topic        *string      `json:"topic,omitempty"`         
	Root         *Root        `json:"root,omitempty"`          
	Upload       *bool        `json:"upload,omitempty"`        
	DisplayAsBot *bool        `json:"display_as_bot,omitempty"`
	BotID        *string      `json:"bot_id,omitempty"`        
	BotLink      *string      `json:"bot_link,omitempty"`      
	BotProfile   *BotProfile  `json:"bot_profile,omitempty"`   
	ThreadTs     *string      `json:"thread_ts,omitempty"`     
	Ts           *string      `json:"ts,omitempty"`            
	Icons        *LatestIcons `json:"icons,omitempty"`         
	Edited       *Edited      `json:"edited,omitempty"`        
}

type BotProfile struct {
	ID      *string          `json:"id,omitempty"`     
	Deleted *bool            `json:"deleted,omitempty"`
	Name    *string          `json:"name,omitempty"`   
	Updated *int64           `json:"updated,omitempty"`
	AppID   *string          `json:"app_id,omitempty"` 
	Icons   *BotProfileIcons `json:"icons,omitempty"`  
	TeamID  *string          `json:"team_id,omitempty"`
}

type BotProfileIcons struct {
	Image36 *string `json:"image_36,omitempty"`
	Image48 *string `json:"image_48,omitempty"`
	Image72 *string `json:"image_72,omitempty"`
}

type Edited struct {
	User *string `json:"user,omitempty"`
	Ts   *string `json:"ts,omitempty"`  
}

type LatestIcons struct {
	Emoji   *string `json:"emoji,omitempty"`   
	Image36 *string `json:"image_36,omitempty"`
	Image48 *string `json:"image_48,omitempty"`
	Image64 *string `json:"image_64,omitempty"`
	Image72 *string `json:"image_72,omitempty"`
}

type Root struct {
	Text            *string      `json:"text,omitempty"`             
	User            *string      `json:"user,omitempty"`             
	ParentUserID    *string      `json:"parent_user_id,omitempty"`   
	Username        *string      `json:"username,omitempty"`         
	Team            *string      `json:"team,omitempty"`             
	BotID           *string      `json:"bot_id,omitempty"`           
	Mrkdwn          *bool        `json:"mrkdwn,omitempty"`           
	Type            *string      `json:"type,omitempty"`             
	Subtype         *string      `json:"subtype,omitempty"`          
	ThreadTs        *string      `json:"thread_ts,omitempty"`        
	Icons           *LatestIcons `json:"icons,omitempty"`            
	BotProfile      *BotProfile  `json:"bot_profile,omitempty"`      
	ReplyCount      *int64       `json:"reply_count,omitempty"`      
	ReplyUsersCount *int64       `json:"reply_users_count,omitempty"`
	LatestReply     *string      `json:"latest_reply,omitempty"`     
	Subscribed      *bool        `json:"subscribed,omitempty"`       
	LastRead        *string      `json:"last_read,omitempty"`        
	UnreadCount     *int64       `json:"unread_count,omitempty"`     
	Ts              *string      `json:"ts,omitempty"`               
}

type Purpose struct {
	Value   *string `json:"value,omitempty"`   
	Creator *string `json:"creator,omitempty"` 
	LastSet *int64  `json:"last_set,omitempty"`
}

type ResponseMetadata struct {
	NextCursor *string `json:"next_cursor,omitempty"`
}
