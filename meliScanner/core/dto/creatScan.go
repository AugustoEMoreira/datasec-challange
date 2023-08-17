package dto

type CreateScanRequest struct {
	Id string `json:"id"`
}

func FromJSONCreateScanRequest(id string) (*CreateScanRequest, error) {
	createScanRequest := CreateScanRequest{}
	createScanRequest.Id = id
	return &createScanRequest, nil
}
