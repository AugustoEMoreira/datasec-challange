package dependencyinjector

import (
	"github.com/AugustoEMoreira/datasec-challange/adapter/http/datastorageservice"
	"github.com/AugustoEMoreira/datasec-challange/adapter/mysql"
	"github.com/AugustoEMoreira/datasec-challange/adapter/mysql/datastoragerepository"
	"github.com/AugustoEMoreira/datasec-challange/core/domain"
	"github.com/AugustoEMoreira/datasec-challange/core/domain/usecase/datastorageusecase"
)

var DatastorageService domain.DatastorageService

func ConfigDatastorage(conn mysql.PoolInterface) domain.DatastorageService {
	dsRepository := datastoragerepository.New(conn)
	dsUseCase := datastorageusecase.New(dsRepository)
	dsService := datastorageservice.New(dsUseCase)

	DatastorageService = dsService
	return dsService
}
