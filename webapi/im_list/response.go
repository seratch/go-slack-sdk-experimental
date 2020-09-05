// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    iMList, err := UnmarshalIMList(bytes)
//    bytes, err = iMList.Marshal()

package im_list

import "encoding/json"

func UnmarshalIMList(data []byte) (IMList, error) {
	var r IMList
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *IMList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type IMList struct {
	Ok               *bool             `json:"ok,omitempty"`               
	Ims              []IM              `json:"ims,omitempty"`              
	Warning          *string           `json:"warning,omitempty"`          
	ResponseMetadata *ResponseMetadata `json:"response_metadata,omitempty"`
	Error            *string           `json:"error,omitempty"`            
	Needed           *string           `json:"needed,omitempty"`           
	Provided         *string           `json:"provided,omitempty"`         
}

type IM struct {
	ID            *string `json:"id,omitempty"`             
	Created       *int64  `json:"created,omitempty"`        
	IsArchived    *bool   `json:"is_archived,omitempty"`    
	IsIM          *bool   `json:"is_im,omitempty"`          
	IsOrgShared   *bool   `json:"is_org_shared,omitempty"`  
	User          *string `json:"user,omitempty"`           
	IsUserDeleted *bool   `json:"is_user_deleted,omitempty"`
	Priority      *int64  `json:"priority,omitempty"`       
}

type ResponseMetadata struct {
	NextCursor *string  `json:"next_cursor,omitempty"`
	Messages   []string `json:"messages,omitempty"`   
	Warnings   []string `json:"warnings,omitempty"`   
}
