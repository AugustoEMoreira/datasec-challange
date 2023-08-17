package domain

import (
	"net/http"

	"github.com/AugustoEMoreira/datasec-challange/core/dto"
)

type Scan struct {
	Db      string        `json:"id"`
	Columns []interface{} `json:"columns"`
}

type ScanService interface {
	Create(response http.ResponseWriter, request *http.Request)
}

type ScanUseCase interface {
	Create(scanRequest *dto.CreateScanRequest) (*Datastorage, error)
}

type ScanRepository interface {
	Create(scanRequest *dto.CreateScanRequest) (*Datastorage, error)
}
