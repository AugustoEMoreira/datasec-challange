package scanusecase

import (
	"github.com/AugustoEMoreira/datasec-challange/core/domain"
)

type usecase struct {
	repository domain.ScanRepository
}

func New(repository domain.ScanRepository) domain.ScanUseCase {
	return &usecase{
		repository: repository,
	}
}
