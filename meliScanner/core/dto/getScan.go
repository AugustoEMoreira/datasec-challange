package dto

import (
	"net/http"
)

type ScanRequestParams struct {
	DBId string `json:"id"`
}

func FromUrlScanRequestParams(request *http.Request) (*ScanRequestParams, error) {
	params := ScanRequestParams{}
	id := request.URL.Path[len("/v1/db/"):]
	params.DBId = id
	return &params, nil
}
