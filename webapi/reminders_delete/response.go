// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    remindersDelete, err := UnmarshalRemindersDelete(bytes)
//    bytes, err = remindersDelete.Marshal()

package reminders_delete

import "encoding/json"

func UnmarshalRemindersDelete(data []byte) (RemindersDelete, error) {
	var r RemindersDelete
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RemindersDelete) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type RemindersDelete struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
