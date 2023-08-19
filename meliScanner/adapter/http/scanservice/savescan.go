package scanservice

import "github.com/AugustoEMoreira/datasec-challange/core/dto"

// SaveScan implements domain.ScanService.
func (service service) SaveScan(scanResult *[]dto.ScanResult) error {
	err := service.usecase.SaveScan(scanResult)
	if err != nil {
		return err
	}
	return nil
}
