// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    reactionsGet, err := UnmarshalReactionsGet(bytes)
//    bytes, err = reactionsGet.Marshal()

package reactions_get

import "encoding/json"

func UnmarshalReactionsGet(data []byte) (ReactionsGet, error) {
	var r ReactionsGet
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ReactionsGet) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ReactionsGet struct {
	Ok       *bool    `json:"ok,omitempty"`      
	Type     *string  `json:"type,omitempty"`    
	Channel  *string  `json:"channel,omitempty"` 
	Message  *Message `json:"message,omitempty"` 
	Error    *string  `json:"error,omitempty"`   
	Needed   *string  `json:"needed,omitempty"`  
	Provided *string  `json:"provided,omitempty"`
}

type Message struct {
	BotID      *string     `json:"bot_id,omitempty"`     
	Type       *string     `json:"type,omitempty"`       
	Text       *string     `json:"text,omitempty"`       
	User       *string     `json:"user,omitempty"`       
	Ts         *string     `json:"ts,omitempty"`         
	Team       *string     `json:"team,omitempty"`       
	BotProfile *BotProfile `json:"bot_profile,omitempty"`
	Reactions  []Reaction  `json:"reactions,omitempty"`  
	Permalink  *string     `json:"permalink,omitempty"`  
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

type Reaction struct {
	Name  *string  `json:"name,omitempty"` 
	Users []string `json:"users,omitempty"`
	Count *int64   `json:"count,omitempty"`
}
