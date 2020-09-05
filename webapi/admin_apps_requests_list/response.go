// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminAppsRequestsList, err := UnmarshalAdminAppsRequestsList(bytes)
//    bytes, err = adminAppsRequestsList.Marshal()

package admin_apps_requests_list

import "encoding/json"

func UnmarshalAdminAppsRequestsList(data []byte) (AdminAppsRequestsList, error) {
	var r AdminAppsRequestsList
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminAppsRequestsList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminAppsRequestsList struct {
	Ok               *bool             `json:"ok,omitempty"`               
	AppRequests      []AppRequest      `json:"app_requests,omitempty"`     
	ResponseMetadata *ResponseMetadata `json:"response_metadata,omitempty"`
	Error            *string           `json:"error,omitempty"`            
	Needed           *string           `json:"needed,omitempty"`           
	Provided         *string           `json:"provided,omitempty"`         
}

type AppRequest struct {
	ID                 *string             `json:"id,omitempty"`                 
	App                *App                `json:"app,omitempty"`                
	User               *User               `json:"user,omitempty"`               
	Team               *Team               `json:"team,omitempty"`               
	PreviousResolution *PreviousResolution `json:"previous_resolution,omitempty"`
	Message            *string             `json:"message,omitempty"`            
	DateCreated        *int64              `json:"date_created,omitempty"`       
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

type PreviousResolution struct {
	Status *string `json:"status,omitempty"`
}

type Team struct {
	ID     *string `json:"id,omitempty"`    
	Name   *string `json:"name,omitempty"`  
	Domain *string `json:"domain,omitempty"`
}

type User struct {
	ID    *string `json:"id,omitempty"`   
	Name  *string `json:"name,omitempty"` 
	Email *string `json:"email,omitempty"`
}

type ResponseMetadata struct {
	NextCursor *string `json:"next_cursor,omitempty"`
}
