package datastoragerepository

import (
	"context"

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

	for rows.Next() {
		err := rows.Scan(&ds.Id, &ds.Host, &ds.Classification)
		if err != nil {
			return nil, err
		}
	}

	return &ds, nil

}
