// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminAppsApprovedList, err := UnmarshalAdminAppsApprovedList(bytes)
//    bytes, err = adminAppsApprovedList.Marshal()

package admin_apps_approved_list

import "encoding/json"

func UnmarshalAdminAppsApprovedList(data []byte) (AdminAppsApprovedList, error) {
	var r AdminAppsApprovedList
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminAppsApprovedList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminAppsApprovedList struct {
	Ok               *bool             `json:"ok,omitempty"`               
	Error            *string           `json:"error,omitempty"`            
	ApprovedApps     []ApprovedApp     `json:"approved_apps,omitempty"`    
	ResponseMetadata *ResponseMetadata `json:"response_metadata,omitempty"`
	Needed           *string           `json:"needed,omitempty"`           
	Provided         *string           `json:"provided,omitempty"`         
}

type ApprovedApp struct {
	App            *App            `json:"app,omitempty"`             
	Scopes         []Scope         `json:"scopes,omitempty"`          
	DateUpdated    *int64          `json:"date_updated,omitempty"`    
	LastResolvedBy *LastResolvedBy `json:"last_resolved_by,omitempty"`
}

type App struct {
	ID                     *string `json:"id,omitempty"`                       
	Name                   *string `json:"name,omitempty"`                     
	Description            *string `json:"description,omitempty"`              
	HelpURL                *string `json:"help_url,omitempty"`                 
	PrivacyPolicyURL       *string `json:"privacy_policy_url,omitempty"`       
	AppHomepageURL         *string `json:"app_homepage_url,omitempty"`         
	AppDirectoryURL        *string `json:"app_directory_url,omitempty"`        
	IsAppDirectoryApproved *bool   `json:"is_app_directory_approved,omitempty"`
	IsInternal             *bool   `json:"is_internal,omitempty"`              
	Icons                  *Icons  `json:"icons,omitempty"`                    
	AdditionalInfo         *string `json:"additional_info,omitempty"`          
}

type Icons struct {
	Image32       *string `json:"image_32,omitempty"`      
	Image36       *string `json:"image_36,omitempty"`      
	Image48       *string `json:"image_48,omitempty"`      
	Image64       *string `json:"image_64,omitempty"`      
	Image72       *string `json:"image_72,omitempty"`      
	Image96       *string `json:"image_96,omitempty"`      
	Image128      *string `json:"image_128,omitempty"`     
	Image192      *string `json:"image_192,omitempty"`     
	Image512      *string `json:"image_512,omitempty"`     
	Image1024     *string `json:"image_1024,omitempty"`    
	ImageOriginal *string `json:"image_original,omitempty"`
}

type LastResolvedBy struct {
	ActorID   *string `json:"actor_id,omitempty"`  
	ActorType *string `json:"actor_type,omitempty"`
}

type Scope struct {
	Name        *string `json:"name,omitempty"`        
	Description *string `json:"description,omitempty"` 
	IsSensitive *bool   `json:"is_sensitive,omitempty"`
	TokenType   *string `json:"token_type,omitempty"`  
}

type ResponseMetadata struct {
	NextCursor *string `json:"next_cursor,omitempty"`
}
