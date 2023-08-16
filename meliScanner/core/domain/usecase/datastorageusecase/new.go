package datastorageusecase

import "github.com/AugustoEMoreira/datasec-challange/core/domain"

type usecase struct {
	repository domain.DatastorageRepository
}

func New(repository domain.DatastorageRepository) domain.DatastorageUseCase {
	return &usecase{
		repository: repository,
	}
}
