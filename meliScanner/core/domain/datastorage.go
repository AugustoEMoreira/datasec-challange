package domain

import (
	"net/http"

	"github.com/AugustoEMoreira/datasec-challange/core/dto"
)

type Datastorage struct {
	Id             int64       `json:"id"`
	Host           string      `json:"host"`
	Classification interface{} `json:"classes"`
}
type DatastorageService interface {
	Create(response http.ResponseWriter, request *http.Request)
	Fetch(response http.ResponseWriter, request *http.Request)
}

type DatastorageUseCase interface {
	Create(datastorageRequest *dto.CreateDatastoreRequest) (*Datastorage, error)
	Fetch(datastorageRequest *dto.FetchDatastoreRequest) (*Datastorage, error)
}

type DatastorageRepository interface {
	Create(datastorageRequest *dto.CreateDatastoreRequest) (*Datastorage, error)
	Fetch(datastorageRequest *dto.FetchDatastoreRequest) (*Datastorage, error)
}
