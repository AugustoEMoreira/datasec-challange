package domain

import (
	"database/sql"
	"net/http"

	"github.com/AugustoEMoreira/datasec-challange/core/dto"
)

type Scan struct {
	Db      string        `json:"id"`
	Columns []interface{} `json:"columns"`
}
type ScanService interface {
	Create(response http.ResponseWriter, request *http.Request)
	GetDatabaseConnInfo(id string) (dbconnUrl *dto.DbConnectionInfo, err error)
	GetDbConnection(d *dto.DbConnectionInfo) (conn *sql.Conn, err error)
	GetRegexRules() (rules []*dto.RulesRegex, err error)
	SaveScan(*[]dto.ScanResult) error
}

type ScanUseCase interface {
	GetDatabaseConnInfo(id string) (dbconnUrl *dto.DbConnectionInfo, err error)
	GetDbConnection(d *dto.DbConnectionInfo) (conn *sql.Conn, err error)
	GetRegexRules() (rules []*dto.RulesRegex, err error)
	SaveScan(*[]dto.ScanResult) error
}

type ScanRepository interface {
	GetDatabaseConnInfo(id string) (dbconnUrl *dto.DbConnectionInfo, err error)
	GetRegexRules() (rules []*dto.RulesRegex, err error)
	GetDbConnection(d *dto.DbConnectionInfo) (conn *sql.Conn, err error)
	SaveScan(*[]dto.ScanResult) error
}
