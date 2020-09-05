// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    remindersComplete, err := UnmarshalRemindersComplete(bytes)
//    bytes, err = remindersComplete.Marshal()

package reminders_complete

import "encoding/json"

func UnmarshalRemindersComplete(data []byte) (RemindersComplete, error) {
	var r RemindersComplete
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RemindersComplete) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type RemindersComplete struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
