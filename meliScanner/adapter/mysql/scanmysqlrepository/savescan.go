package scanrepository

import (
	"context"
	"encoding/json"

	"github.com/AugustoEMoreira/datasec-challange/core/dto"
)

// SaveScan implements domain.ScanRepository.
func (respository repository) SaveScan(sc *[]dto.ScanResult, dbId string) error {
	ctx := context.Background()

	query := `UPDATE datastorage SET classification = ? WHERE id = ?;`
	stmt, err := respository.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	jsonScan, err := json.Marshal(sc)
	_, err = stmt.Exec(jsonScan, dbId)
	if err != nil {
		return err
	}

	return nil
}
