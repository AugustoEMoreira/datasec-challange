package dto

import (
	"encoding/json"
	"io"
)

type CreateScanRequest struct {
	Id string `json:"id"`
}
type DbConnectionInfo struct {
	UrlString string
}

type RulesRegex struct {
	Rule       string
	Annotation string
}

type ScanResult struct {
	Table  string
	Column string
	Class  string
}

func FromIdCreateScanRequest(id string) (*CreateScanRequest, error) {
	createScanRequest := CreateScanRequest{}
	createScanRequest.Id = id
	return &createScanRequest, nil
}

type FetchScanRequest struct {
	Id string `json:"id"`
}

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

type FetchRegexRequest struct {
	RegexSelector string `json:"regex"`
}
