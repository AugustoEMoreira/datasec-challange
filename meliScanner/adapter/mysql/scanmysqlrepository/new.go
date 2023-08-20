package scanrepository

import (
	"github.com/AugustoEMoreira/datasec-challange/adapter/mysql"
	"github.com/AugustoEMoreira/datasec-challange/core/domain"
)

type repository struct {
	db mysql.PoolInterface
}

// GetDbConnection implements domain.ScanRepository.

func New(db mysql.PoolInterface) domain.ScanRepository {
	return &repository{
		db: db,
	}
}
