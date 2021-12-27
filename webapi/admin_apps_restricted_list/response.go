// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminAppsRestrictedList, err := UnmarshalAdminAppsRestrictedList(bytes)
//    bytes, err = adminAppsRestrictedList.Marshal()

package admin_apps_restricted_list

import "encoding/json"

func UnmarshalAdminAppsRestrictedList(data []byte) (AdminAppsRestrictedList, error) {
	var r AdminAppsRestrictedList
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminAppsRestrictedList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminAppsRestrictedList struct {
	Ok               *bool             `json:"ok,omitempty"`               
	Warning          *string           `json:"warning,omitempty"`          
	Error            *string           `json:"error,omitempty"`            
	Needed           *string           `json:"needed,omitempty"`           
	Provided         *string           `json:"provided,omitempty"`         
	RestrictedApps   []RestrictedApp   `json:"restricted_apps,omitempty"`  
	ResponseMetadata *ResponseMetadata `json:"response_metadata,omitempty"`
}

type ResponseMetadata struct {
	NextCursor *string  `json:"next_cursor,omitempty"`
	Messages   []string `json:"messages,omitempty"`   
	Warnings   []string `json:"warnings,omitempty"`   
}

type RestrictedApp struct {
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
	AdditionalInfo         *string `json:"additional_info,omitempty"`          
	Icons                  *Icons  `json:"icons,omitempty"`                    
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
