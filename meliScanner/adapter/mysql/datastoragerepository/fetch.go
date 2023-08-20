package datastoragerepository

import (
	"context"
	"encoding/json"

	"github.com/AugustoEMoreira/datasec-challange/core/domain"
	"github.com/AugustoEMoreira/datasec-challange/core/dto"
)

// Fetch implements domain.DatastorageRepository.
func (repository repository) Fetch(datastorageRequest *dto.FetchDatastoreRequest) (*domain.Datastorage, error) {
	ctx := context.Background()
	ds := domain.Datastorage{}
	query := `select id,host,classification from datastorage where id = ?`
	stmt, err := repository.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(datastorageRequest.DBId)
	defer rows.Close()

	var count int
	for rows.Next() {
		count++
		var jsonStr []byte
		err := rows.Scan(&ds.Id, &ds.Host, &jsonStr)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(jsonStr, &ds.Classification)
		if err != nil {
			return nil, err
		}
	}

	if count == 0 {
		return nil, nil
	}
	return &ds, nil

}
