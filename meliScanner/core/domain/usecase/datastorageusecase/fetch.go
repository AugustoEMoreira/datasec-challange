package datastorageusecase

import (
	"github.com/AugustoEMoreira/datasec-challange/core/domain"
	"github.com/AugustoEMoreira/datasec-challange/core/dto"
)

func (usecase usecase) Fetch(fetchDatastoreRequest *dto.FetchDatastoreRequest) (*domain.Datastorage, error) {
	ds, err := usecase.repository.Fetch(fetchDatastoreRequest)
	if err != nil {
		return nil, err
	}
	return ds, nil
}
