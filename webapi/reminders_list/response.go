// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    remindersList, err := UnmarshalRemindersList(bytes)
//    bytes, err = remindersList.Marshal()

package reminders_list

import "encoding/json"

func UnmarshalRemindersList(data []byte) (RemindersList, error) {
	var r RemindersList
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RemindersList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type RemindersList struct {
	Ok        *bool      `json:"ok,omitempty"`       
	Reminders []Reminder `json:"reminders,omitempty"`
	Error     *string    `json:"error,omitempty"`    
	Needed    *string    `json:"needed,omitempty"`   
	Provided  *string    `json:"provided,omitempty"` 
}

type Reminder struct {
	ID         *string     `json:"id,omitempty"`         
	Creator    *string     `json:"creator,omitempty"`    
	Text       *string     `json:"text,omitempty"`       
	User       *string     `json:"user,omitempty"`       
	Recurring  *bool       `json:"recurring,omitempty"`  
	Time       *int64      `json:"time,omitempty"`       
	CompleteTs *int64      `json:"complete_ts,omitempty"`
	Channel    *string     `json:"channel,omitempty"`    
	Recurrence *Recurrence `json:"recurrence,omitempty"` 
}

type Recurrence struct {
	Frequency *string  `json:"frequency,omitempty"`
	Weekdays  []string `json:"weekdays,omitempty"` 
}
