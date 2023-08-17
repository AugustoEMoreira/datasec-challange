package datastoragerepository

import (
	"context"

	"github.com/AugustoEMoreira/datasec-challange/core/domain"
	"github.com/AugustoEMoreira/datasec-challange/core/dto"
)

func (repository repository) Create(datastorageRequest *dto.CreateDatastoreRequest) (*domain.Datastorage, error) {
	ctx := context.Background()
	ds := domain.Datastorage{
		Host:           datastorageRequest.Host,
		Classification: []string{},
	}

	query := `INSERT INTO datastorage (user, pass, host, port) values (?,?,?,?)`

	stmt, err := repository.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	result, err := stmt.Exec(datastorageRequest.User, datastorageRequest.Password, datastorageRequest.Host, datastorageRequest.Port)

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	ds.Id = id
	return &ds, nil
}
