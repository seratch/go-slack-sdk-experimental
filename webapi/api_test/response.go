// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    aPITest, err := UnmarshalAPITest(bytes)
//    bytes, err = aPITest.Marshal()

package api_test

import "encoding/json"

func UnmarshalAPITest(data []byte) (APITest, error) {
	var r APITest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *APITest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type APITest struct {
	Ok       *bool   `json:"ok,omitempty"`      
	Error    *string `json:"error,omitempty"`   
	Args     *Args   `json:"args,omitempty"`    
	Needed   *string `json:"needed,omitempty"`  
	Provided *string `json:"provided,omitempty"`
}

type Args struct {
	Error *string `json:"error,omitempty"`
	Foo   *string `json:"foo,omitempty"`  
}
