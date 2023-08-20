package domain

import (
	"net/http"

	"github.com/AugustoEMoreira/datasec-challange/core/dto"
)

type ScanService interface {
	Create(response http.ResponseWriter, request *http.Request)
	GetDatabaseConnInfo(id string) (dbconnUrl *dto.DbConnectionInfo, err error)
	GetRegexRules() (rules *[]dto.RulesRegex, err error)
	SaveScan(sc *[]dto.ScanResult, dbId string) error
}

type ScanUseCase interface {
	GetDatabaseConnInfo(id string) (dbconnUrl *dto.DbConnectionInfo, err error)
	GetRegexRules() (rules *[]dto.RulesRegex, err error)
	SaveScan(sc *[]dto.ScanResult, dbId string) error
}

type ScanRepository interface {
	GetDatabaseConnInfo(id string) (dbconnUrl *dto.DbConnectionInfo, err error)
	GetRegexRules() (rules *[]dto.RulesRegex, err error)
	SaveScan(sc *[]dto.ScanResult, dbId string) error
}
