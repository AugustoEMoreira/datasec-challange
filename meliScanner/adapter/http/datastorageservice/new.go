package datastorageservice

import (
	"github.com/AugustoEMoreira/datasec-challange/core/domain"
)

type service struct {
	usecase domain.DatastorageUseCase
}

func New(usecase domain.DatastorageUseCase) domain.DatastorageService {
	return &service{
		usecase: usecase,
	}
}
