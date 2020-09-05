// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    iMReplies, err := UnmarshalIMReplies(bytes)
//    bytes, err = iMReplies.Marshal()

package im_replies

import "encoding/json"

func UnmarshalIMReplies(data []byte) (IMReplies, error) {
	var r IMReplies
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *IMReplies) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type IMReplies struct {
	Messages         []Message         `json:"messages,omitempty"`         
	HasMore          *bool             `json:"has_more,omitempty"`         
	Ok               *bool             `json:"ok,omitempty"`               
	Warning          *string           `json:"warning,omitempty"`          
	ResponseMetadata *ResponseMetadata `json:"response_metadata,omitempty"`
	Error            *string           `json:"error,omitempty"`            
	Needed           *string           `json:"needed,omitempty"`           
	Provided         *string           `json:"provided,omitempty"`         
}

type Message struct {
	BotID           *string     `json:"bot_id,omitempty"`           
	Type            *string     `json:"type,omitempty"`             
	Text            *string     `json:"text,omitempty"`             
	User            *string     `json:"user,omitempty"`             
	Ts              *string     `json:"ts,omitempty"`               
	Team            *string     `json:"team,omitempty"`             
	BotProfile      *BotProfile `json:"bot_profile,omitempty"`      
	ThreadTs        *string     `json:"thread_ts,omitempty"`        
	ReplyCount      *int64      `json:"reply_count,omitempty"`      
	ReplyUsersCount *int64      `json:"reply_users_count,omitempty"`
	LatestReply     *string     `json:"latest_reply,omitempty"`     
	ReplyUsers      []string    `json:"reply_users,omitempty"`      
	Subscribed      *bool       `json:"subscribed,omitempty"`       
	LastRead        *string     `json:"last_read,omitempty"`        
	ParentUserID    *string     `json:"parent_user_id,omitempty"`   
}

type BotProfile struct {
	ID      *string `json:"id,omitempty"`     
	Deleted *bool   `json:"deleted,omitempty"`
	Name    *string `json:"name,omitempty"`   
	Updated *int64  `json:"updated,omitempty"`
	AppID   *string `json:"app_id,omitempty"` 
	Icons   *Icons  `json:"icons,omitempty"`  
	TeamID  *string `json:"team_id,omitempty"`
}

type Icons struct {
	Image36 *string `json:"image_36,omitempty"`
	Image48 *string `json:"image_48,omitempty"`
	Image72 *string `json:"image_72,omitempty"`
}

type ResponseMetadata struct {
	Messages []string `json:"messages,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
}