package scanrepository

import (
	"context"

	"github.com/AugustoEMoreira/datasec-challange/core/dto"
)

// GetRegexRules implements domain.ScanRepository.
func (repository repository) GetRegexRules() (rules *[]dto.RulesRegex, err error) {
	ctx := context.Background()

	var rArray []dto.RulesRegex
	query := `select regex, annotation from rules`

	rows, err := repository.db.QueryContext(ctx, query)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var r dto.RulesRegex
		err := rows.Scan(&r.Rule, &r.Annotation)
		if err != nil {
			return nil, err
		}
		rArray = append(rArray, r)
	}
	return &rArray, nil
}
