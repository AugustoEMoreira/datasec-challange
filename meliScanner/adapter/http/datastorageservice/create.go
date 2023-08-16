package datastorageservice

import (
	"encoding/json"
	"net/http"

	"github.com/AugustoEMoreira/datasec-challange/core/dto"
)

func (service service) Create(response http.ResponseWriter, request *http.Request) {
	dsRequest, err := dto.FromJSONCreateDatastoreRequest(request.Body)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}
	ds, err := service.usecase.Create(dsRequest)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(response).Encode(ds)
}
