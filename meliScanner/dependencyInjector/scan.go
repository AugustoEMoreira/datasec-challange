package dependencyinjector

import (
	"github.com/AugustoEMoreira/datasec-challange/adapter/http/scanservice"
	"github.com/AugustoEMoreira/datasec-challange/adapter/mysql"
	scanrepository "github.com/AugustoEMoreira/datasec-challange/adapter/mysql/scanmysqlrepository"
	"github.com/AugustoEMoreira/datasec-challange/core/domain"
	"github.com/AugustoEMoreira/datasec-challange/core/domain/usecase/scanusecase"
)

func ConfigScan(conn mysql.PoolInterface) domain.ScanService {
	scRepository := scanrepository.New(conn)
	scUseCase := scanusecase.New(scRepository)
	scService := scanservice.New(scUseCase)
	return scService
}
