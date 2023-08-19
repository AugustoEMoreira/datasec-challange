package scanusecase

import "github.com/AugustoEMoreira/datasec-challange/core/dto"

// GetRegexRules implements domain.ScanUseCase.
func (usecase usecase) GetRegexRules() (rules []*dto.RulesRegex, err error) {
	rgx, err := usecase.repository.GetRegexRules()
	if err != nil {
		return nil, err
	}
	return rgx, nil
}
