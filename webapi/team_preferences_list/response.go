// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    teamPreferencesList, err := UnmarshalTeamPreferencesList(bytes)
//    bytes, err = teamPreferencesList.Marshal()

package team_preferences_list

import "encoding/json"

func UnmarshalTeamPreferencesList(data []byte) (TeamPreferencesList, error) {
	var r TeamPreferencesList
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TeamPreferencesList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TeamPreferencesList struct {
	Ok                   *bool   `json:"ok,omitempty"`                    
	MsgEditWindowMins    *int64  `json:"msg_edit_window_mins,omitempty"`  
	AllowMessageDeletion *bool   `json:"allow_message_deletion,omitempty"`
	DisplayRealNames     *bool   `json:"display_real_names,omitempty"`    
	DisableFileUploads   *string `json:"disable_file_uploads,omitempty"`  
	WhoCanPostGeneral    *string `json:"who_can_post_general,omitempty"`  
	Error                *string `json:"error,omitempty"`                 
	Needed               *string `json:"needed,omitempty"`                
	Provided             *string `json:"provided,omitempty"`              
}
