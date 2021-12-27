// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    openidConnectUserInfo, err := UnmarshalOpenidConnectUserInfo(bytes)
//    bytes, err = openidConnectUserInfo.Marshal()

package openid_connect_userInfo

import "encoding/json"

func UnmarshalOpenidConnectUserInfo(data []byte) (OpenidConnectUserInfo, error) {
	var r OpenidConnectUserInfo
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *OpenidConnectUserInfo) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type OpenidConnectUserInfo struct {
	Ok                            *bool   `json:"ok,omitempty"`                                 
	Warning                       *string `json:"warning,omitempty"`                            
	Error                         *string `json:"error,omitempty"`                              
	Needed                        *string `json:"needed,omitempty"`                             
	Provided                      *string `json:"provided,omitempty"`                           
	Sub                           *string `json:"sub,omitempty"`                                
	HTTPSSlackCOMUserID           *string `json:"https://slack.com/user_id,omitempty"`          
	HTTPSSlackCOMTeamID           *string `json:"https://slack.com/team_id,omitempty"`          
	HTTPSSlackCOMEnterpriseID     *string `json:"https://slack.com/enterprise_id,omitempty"`    
	Email                         *string `json:"email,omitempty"`                              
	EmailVerified                 *bool   `json:"email_verified,omitempty"`                     
	DateEmailVerified             *int64  `json:"date_email_verified,omitempty"`                
	Name                          *string `json:"name,omitempty"`                               
	Picture                       *string `json:"picture,omitempty"`                            
	GivenName                     *string `json:"given_name,omitempty"`                         
	FamilyName                    *string `json:"family_name,omitempty"`                        
	Locale                        *string `json:"locale,omitempty"`                             
	HTTPSSlackCOMTeamName         *string `json:"https://slack.com/team_name,omitempty"`        
	HTTPSSlackCOMTeamDomain       *string `json:"https://slack.com/team_domain,omitempty"`      
	HTTPSSlackCOMEnterpriseName   *string `json:"https://slack.com/enterprise_name,omitempty"`  
	HTTPSSlackCOMEnterpriseDomain *string `json:"https://slack.com/enterprise_domain,omitempty"`
	HTTPSSlackCOMUserImage24      *string `json:"https://slack.com/user_image_24,omitempty"`    
	HTTPSSlackCOMUserImage32      *string `json:"https://slack.com/user_image_32,omitempty"`    
	HTTPSSlackCOMUserImage48      *string `json:"https://slack.com/user_image_48,omitempty"`    
	HTTPSSlackCOMUserImage72      *string `json:"https://slack.com/user_image_72,omitempty"`    
	HTTPSSlackCOMUserImage192     *string `json:"https://slack.com/user_image_192,omitempty"`   
	HTTPSSlackCOMUserImage512     *string `json:"https://slack.com/user_image_512,omitempty"`   
	HTTPSSlackCOMUserImage1024    *string `json:"https://slack.com/user_image_1024,omitempty"`  
	HTTPSSlackCOMTeamImage34      *string `json:"https://slack.com/team_image_34,omitempty"`    
	HTTPSSlackCOMTeamImage44      *string `json:"https://slack.com/team_image_44,omitempty"`    
	HTTPSSlackCOMTeamImage68      *string `json:"https://slack.com/team_image_68,omitempty"`    
	HTTPSSlackCOMTeamImage88      *string `json:"https://slack.com/team_image_88,omitempty"`    
	HTTPSSlackCOMTeamImage102     *string `json:"https://slack.com/team_image_102,omitempty"`   
	HTTPSSlackCOMTeamImage132     *string `json:"https://slack.com/team_image_132,omitempty"`   
	HTTPSSlackCOMTeamImage230     *string `json:"https://slack.com/team_image_230,omitempty"`   
}
