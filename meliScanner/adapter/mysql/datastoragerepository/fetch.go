package datastoragerepository

import (
	"github.com/AugustoEMoreira/datasec-challange/core/domain"
	"github.com/AugustoEMoreira/datasec-challange/core/dto"
)

// Fetch implements domain.DatastorageRepository.
func (*repository) Fetch(datastorageRequest *dto.FetchDatastoreRequest) (*domain.Datastorage, error) {
	panic("unimplemented")
}
