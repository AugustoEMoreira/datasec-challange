package scanservice

import "github.com/AugustoEMoreira/datasec-challange/core/dto"

// GetRegexRules implements domain.ScanService.
func (service service) GetRegexRules() (rules *[]dto.RulesRegex, err error) {

	r, err := service.usecase.GetRegexRules()
	if err != nil {
		return nil, err
	}
	return r, nil
}
