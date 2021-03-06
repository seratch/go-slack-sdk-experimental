// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    groupsClose, err := UnmarshalGroupsClose(bytes)
//    bytes, err = groupsClose.Marshal()

package groups_close

import "encoding/json"

func UnmarshalGroupsClose(data []byte) (GroupsClose, error) {
	var r GroupsClose
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GroupsClose) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type GroupsClose struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
