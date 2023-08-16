package datastoragerepository

import (
	"github.com/AugustoEMoreira/datasec-challange/adapter/mysql"
	"github.com/AugustoEMoreira/datasec-challange/core/domain"
)

type repository struct {
	db mysql.PoolInterface
}

func New(db mysql.PoolInterface) domain.DatastorageRepository {
	return &repository{
		db: db,
	}
}
