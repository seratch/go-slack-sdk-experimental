// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    adminAnalyticsGetFile, err := UnmarshalAdminAnalyticsGetFile(bytes)
//    bytes, err = adminAnalyticsGetFile.Marshal()

package admin_analytics_getFile

import "encoding/json"

func UnmarshalAdminAnalyticsGetFile(data []byte) (AdminAnalyticsGetFile, error) {
	var r AdminAnalyticsGetFile
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminAnalyticsGetFile) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AdminAnalyticsGetFile struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}
