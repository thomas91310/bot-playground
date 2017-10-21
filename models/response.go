package models

import (
	"encoding/json"
)

const (
	//Version (API Version)
	Version = "1.0"
)

//Response enveloppe
type Response struct {
	//Data contains the actual data
	Data interface{} `json: "data"`
	//Meta adds metadata to the response
	Meta Meta `json: "meta"`
}

//Meta contains a status, a HTTP code and a version
type Meta struct {
	Status  string `json "status"`
	Code    int    `json: "code"`
	Version string `json: "version"`
}

//MakeBadResp returns a fromatted bad response
func MakeBadResp(code int) string {
	resp := Response{
		Meta: Meta{
			Status:  "ERROR",
			Code:    code,
			Version: Version,
		},
		Data: nil,
	}
	badResp, _ := json.Marshal(resp)
	return string(badResp)
}

//MakeResp makes a good response.
//If it fails to encode json it will return MakeBadResp
func MakeResp(code int, data interface{}) string {
	resp := Response{
		Meta: Meta{
			Status:  "OK",
			Code:    200,
			Version: Version,
		},
		Data: data,
	}
	goodResp, err := json.Marshal(resp)
	if err != nil {
		return MakeBadResp(400)
	}
	return string(goodResp)
}
