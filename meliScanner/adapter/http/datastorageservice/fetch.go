package datastorageservice

import (
	"encoding/json"
	"net/http"

	"github.com/AugustoEMoreira/datasec-challange/core/dto"
	"github.com/gorilla/mux"
)

func (service service) Fetch(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	dsId, err := dto.FromJSONFetchDatastoreRequest(id)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}
	ds, err := service.usecase.Fetch(dsId)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(response).Encode(ds)

}
