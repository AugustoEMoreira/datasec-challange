package dto

import (
	"encoding/json"
	"io"
)

type CreateDatastoreRequest struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

func FromJSONCreateDatastoreRequest(body io.Reader) (*CreateDatastoreRequest, error) {
	createDatastoreRequest := CreateDatastoreRequest{}
	if err := json.NewDecoder(body).Decode(&createDatastoreRequest); err != nil {
		return nil, err
	}
	return &createDatastoreRequest, nil
}
