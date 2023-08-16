package dto

type FetchDatastoreRequest struct {
	DBId string `json:"id"`
}

func FromJSONFetchDatastoreRequest(id string) (*FetchDatastoreRequest, error) {
	FetchDatastoreRequest := FetchDatastoreRequest{}
	FetchDatastoreRequest.DBId = id
	return &FetchDatastoreRequest, nil
}
