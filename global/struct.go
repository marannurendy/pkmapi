package global

import "encoding/json"

type APIResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

type APIGetResponse struct {
	Code int `json:"responseCode"`
	Message string `json:"responseDescription"`
	Data json.RawMessage `json:"data"`
}