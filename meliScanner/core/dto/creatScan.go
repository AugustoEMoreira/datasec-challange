package dto

type CreateScanRequest struct {
	Id string `json:"id"`
}
type DbConnectionInfo struct {
	UrlString string
}

type RulesRegex struct {
	Rule       string
	Annotation string
}

type ScanResult struct {
	Column string
	Class  string
}

func FromIdCreateScanRequest(id string) (*CreateScanRequest, error) {
	createScanRequest := CreateScanRequest{}
	createScanRequest.Id = id
	return &createScanRequest, nil
}
