package scanservice

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/AugustoEMoreira/datasec-challange/adapter/mysql"
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
	if dsInfo == nil {
		response.WriteHeader(404)
		return
	}

	regexRules, err := service.usecase.GetRegexRules()
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	conn, err := mysql.GetCustomDBConnection(ctx, dsInfo.UrlString)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}
	defer conn.Close()

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	var scanResult []dto.ScanResult

	for _, r := range *regexRules {
		query := `select column_name,table_name from information_schema.columns where column_name REGEXP ? and table_schema not in ('information_schema', 'performance_schema', 'mysql', 'sys')`
		stmt, err := conn.PrepareContext(ctx, query)
		if err != nil {
			response.WriteHeader(500)
			response.Write([]byte(err.Error()))
			return
		}

		rows, err := stmt.Query(r.Rule)
		for rows.Next() {
			var row dto.ScanResult
			err := rows.Scan(&row.Column, &row.Table)
			if err != nil {
				response.WriteHeader(500)
				response.Write([]byte(err.Error()))
				return
			}
			row.Class = r.Annotation
			scanResult = append(scanResult, row)
		}
	}

	err = service.usecase.SaveScan(&scanResult, id)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(response).Encode(scanResult)
}
