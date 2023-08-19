package scanservice

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/AugustoEMoreira/datasec-challange/core/dto"
	"github.com/gorilla/mux"
)

func (service service) Create(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	ctx := context.Background()

	dsInfo, err := service.usecase.GetDatabaseConnInfo(id)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	regexRules, err := service.usecase.GetRegexRules()
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	conn, err := service.usecase.GetDbConnection(dsInfo)

	var scanResult []dto.ScanResult

	for _, r := range regexRules {
		query := `select column_name from information_schema.columns where column_name REGEXP ? and table_schema not in ('information_schema', 'performance_schema', 'mysql', 'sys')`
		stmt, err := conn.PrepareContext(ctx, query)
		if err != nil {
			response.WriteHeader(500)
			response.Write([]byte(err.Error()))
			return
		}

		rows, err := stmt.Query(r.Rule)
		for rows.Next() {
			var row dto.ScanResult
			err := rows.Scan(&row.Column)
			if err != nil {
				response.WriteHeader(500)
				response.Write([]byte(err.Error()))
				return
			}
			row.Class = r.Annotation
			scanResult = append(scanResult, row)
		}
	}

	err = service.usecase.SaveScan(&scanResult)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(response).Encode(scanResult)
}
