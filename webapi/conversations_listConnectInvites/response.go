// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    conversationsListConnectInvites, err := UnmarshalConversationsListConnectInvites(bytes)
//    bytes, err = conversationsListConnectInvites.Marshal()

package conversations_listConnectInvites

import "encoding/json"

func UnmarshalConversationsListConnectInvites(data []byte) (ConversationsListConnectInvites, error) {
	var r ConversationsListConnectInvites
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationsListConnectInvites) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ConversationsListConnectInvites struct {
	Ok               *bool             `json:"ok,omitempty"`               
	Error            *string           `json:"error,omitempty"`            
	Arg              *string           `json:"arg,omitempty"`              
	Invites          []InviteElement   `json:"invites,omitempty"`          
	ResponseMetadata *ResponseMetadata `json:"response_metadata,omitempty"`
	Needed           *string           `json:"needed,omitempty"`           
	Provided         *string           `json:"provided,omitempty"`         
}

type InviteElement struct {
	Direction       *string       `json:"direction,omitempty"`        
	Status          *string       `json:"status,omitempty"`           
	DateLastUpdated *int64        `json:"date_last_updated,omitempty"`
	InviteType      *string       `json:"invite_type,omitempty"`      
	Invite          *InviteInvite `json:"invite,omitempty"`           
	Channel         *Channel      `json:"channel,omitempty"`          
	Acceptances     []Acceptance  `json:"acceptances,omitempty"`      
}

type Acceptance struct {
	ApprovalStatus  *string   `json:"approval_status,omitempty"`  
	DateAccepted    *int64    `json:"date_accepted,omitempty"`    
	DateInvalid     *int64    `json:"date_invalid,omitempty"`     
	DateLastUpdated *int64    `json:"date_last_updated,omitempty"`
	AcceptingTeam   *IngTeam  `json:"accepting_team,omitempty"`   
	AcceptingUser   *TingUser `json:"accepting_user,omitempty"`   
	Reviews         []Review  `json:"reviews,omitempty"`          
}

type IngTeam struct {
	ID          *string `json:"id,omitempty"`          
	Name        *string `json:"name,omitempty"`        
	Icon        *Icon   `json:"icon,omitempty"`        
	IsVerified  *bool   `json:"is_verified,omitempty"` 
	Domain      *string `json:"domain,omitempty"`      
	DateCreated *int64  `json:"date_created,omitempty"`
}

type Icon struct {
	Image102      *string `json:"image_102,omitempty"`     
	Image132      *string `json:"image_132,omitempty"`     
	Image230      *string `json:"image_230,omitempty"`     
	Image34       *string `json:"image_34,omitempty"`      
	Image44       *string `json:"image_44,omitempty"`      
	Image68       *string `json:"image_68,omitempty"`      
	Image88       *string `json:"image_88,omitempty"`      
	ImageOriginal *string `json:"image_original,omitempty"`
	ImageDefault  *bool   `json:"image_default,omitempty"` 
}

type TingUser struct {
	ID      *string  `json:"id,omitempty"`     
	TeamID  *string  `json:"team_id,omitempty"`
	Name    *string  `json:"name,omitempty"`   
	Updated *int64   `json:"updated,omitempty"`
	Profile *Profile `json:"profile,omitempty"`
}

type Profile struct {
	RealName              *string `json:"real_name,omitempty"`              
	DisplayName           *string `json:"display_name,omitempty"`           
	RealNameNormalized    *string `json:"real_name_normalized,omitempty"`   
	DisplayNameNormalized *string `json:"display_name_normalized,omitempty"`
	Team                  *string `json:"team,omitempty"`                   
	AvatarHash            *string `json:"avatar_hash,omitempty"`            
	Email                 *string `json:"email,omitempty"`                  
	Image24               *string `json:"image_24,omitempty"`               
	Image32               *string `json:"image_32,omitempty"`               
	Image48               *string `json:"image_48,omitempty"`               
	Image72               *string `json:"image_72,omitempty"`               
	Image192              *string `json:"image_192,omitempty"`              
	Image512              *string `json:"image_512,omitempty"`              
	Image1024             *string `json:"image_1024,omitempty"`             
	ImageOriginal         *string `json:"image_original,omitempty"`         
	IsCustomImage         *bool   `json:"is_custom_image,omitempty"`        
}

type Review struct {
	Type          *string  `json:"type,omitempty"`          
	DateReview    *int64   `json:"date_review,omitempty"`   
	ReviewingTeam *IngTeam `json:"reviewing_team,omitempty"`
}

type Channel struct {
	ID        *string `json:"id,omitempty"`        
	IsIM      *bool   `json:"is_im,omitempty"`     
	IsPrivate *bool   `json:"is_private,omitempty"`
	Name      *string `json:"name,omitempty"`      
}

type InviteInvite struct {
	ID              *string   `json:"id,omitempty"`               
	DateCreated     *int64    `json:"date_created,omitempty"`     
	DateInvalid     *int64    `json:"date_invalid,omitempty"`     
	InvitingTeam    *IngTeam  `json:"inviting_team,omitempty"`    
	InvitingUser    *TingUser `json:"inviting_user,omitempty"`    
	Link            *string   `json:"link,omitempty"`             
	RecipientUserID *string   `json:"recipient_user_id,omitempty"`
	RecipientEmail  *string   `json:"recipient_email,omitempty"`  
}

type ResponseMetadata struct {
	NextCursor *string  `json:"next_cursor,omitempty"`
	Messages   []string `json:"messages,omitempty"`   
}
