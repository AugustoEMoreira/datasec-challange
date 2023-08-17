package scanrepository

import (
	"github.com/AugustoEMoreira/datasec-challange/core/domain"
	"github.com/AugustoEMoreira/datasec-challange/core/dto"
)

// Create implements domain.ScanRepository.
func (*repository) Create(scanRequest *dto.CreateScanRequest) (*domain.Datastorage, error) {
	panic("unimplemented")
}
