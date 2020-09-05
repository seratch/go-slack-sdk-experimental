// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    botsInfo, err := UnmarshalBotsInfo(bytes)
//    bytes, err = botsInfo.Marshal()

package bots_info

import "encoding/json"

func UnmarshalBotsInfo(data []byte) (BotsInfo, error) {
	var r BotsInfo
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BotsInfo) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type BotsInfo struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Bot      *Bot    `json:"bot,omitempty"`     
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}

type Bot struct {
	ID      *string `json:"id,omitempty"`     
	Deleted *bool   `json:"deleted,omitempty"`
	Name    *string `json:"name,omitempty"`   
	Updated *int64  `json:"updated,omitempty"`
	AppID   *string `json:"app_id,omitempty"` 
	UserID  *string `json:"user_id,omitempty"`
	Icons   *Icons  `json:"icons,omitempty"`  
}

type Icons struct {
	Image36 *string `json:"image_36,omitempty"`
	Image48 *string `json:"image_48,omitempty"`
	Image72 *string `json:"image_72,omitempty"`
}
