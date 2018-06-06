package present

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

type Result struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (r *Result) Format2Json() []byte {
	data, err := json.Marshal(r)
	if err != nil {
		log.Errorf("format result to json failed :", r)
		return nil
	}
	return data
}
