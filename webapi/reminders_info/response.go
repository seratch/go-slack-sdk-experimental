// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    remindersInfo, err := UnmarshalRemindersInfo(bytes)
//    bytes, err = remindersInfo.Marshal()

package reminders_info

import "encoding/json"

func UnmarshalRemindersInfo(data []byte) (RemindersInfo, error) {
	var r RemindersInfo
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RemindersInfo) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type RemindersInfo struct {
	Ok       *bool     `json:"ok,omitempty"`      
	Reminder *Reminder `json:"reminder,omitempty"`
	Error    *string   `json:"error,omitempty"`   
	Needed   *string   `json:"needed,omitempty"`  
	Provided *string   `json:"provided,omitempty"`
}

type Reminder struct {
	ID         *string `json:"id,omitempty"`         
	Creator    *string `json:"creator,omitempty"`    
	Text       *string `json:"text,omitempty"`       
	User       *string `json:"user,omitempty"`       
	Recurring  *bool   `json:"recurring,omitempty"`  
	Time       *int64  `json:"time,omitempty"`       
	CompleteTs *int64  `json:"complete_ts,omitempty"`
}
