package dto

import (
	"encoding/json"
	"io"
)

type CreateRegexRequest struct {
	RegexSelector string `json:"regex"`
	Notation      string `json:"notation"`
}

func FromJSONCreateRegexRequest(body io.Reader) (*CreateRegexRequest, error) {
	createRegexRequest := CreateRegexRequest{}
	if err := json.NewDecoder(body).Decode(&createRegexRequest); err != nil {
		return nil, err
	}
	return &createRegexRequest, nil
}
