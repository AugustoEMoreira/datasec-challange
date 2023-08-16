package domain

type Scan struct {
	Db      string   `json:"id"`
	Columns []Column `json:"columns"`
}
type Column struct {
	Value          string `json:"value"`
	Classification string `json:"classification"`
}
