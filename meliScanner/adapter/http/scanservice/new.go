package scanservice

import (
	"github.com/AugustoEMoreira/datasec-challange/core/domain"
)

type service struct {
	usecase domain.ScanUseCase
}

func New(usecase domain.ScanUseCase) domain.ScanService {
	return &service{
		usecase: usecase,
	}
}
