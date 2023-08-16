package dto

import (
	"encoding/json"
	"io"
)

type CreateScanRequest struct {
	Id string `json:"id"`
}

func FromJSONCreateScanRequest(body io.Reader) (*CreateScanRequest, error) {
	createScanRequest := CreateScanRequest{}
	if err := json.NewDecoder(body).Decode(&createScanRequest); err != nil {
		return nil, err
	}
	return &createScanRequest, nil
}
