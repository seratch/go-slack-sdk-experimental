// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminConversationsEkmListOriginalConnectedChannelInfo, err := UnmarshalAdminConversationsEkmListOriginalConnectedChannelInfo(bytes)
//    bytes, err = adminConversationsEkmListOriginalConnectedChannelInfo.Marshal()

package admin_conversations_ekm_listOriginalConnectedChannelInfo

import "encoding/json"

func UnmarshalAdminConversationsEkmListOriginalConnectedChannelInfo(data []byte) (AdminConversationsEkmListOriginalConnectedChannelInfo, error) {
	var r AdminConversationsEkmListOriginalConnectedChannelInfo
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminConversationsEkmListOriginalConnectedChannelInfo) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminConversationsEkmListOriginalConnectedChannelInfo struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
