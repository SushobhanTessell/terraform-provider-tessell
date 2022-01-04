package models

type APIMetadata struct {
	Records    int32  `json:"records"`
	TimeZone   string `json:"timeZone"`
	Pagination struct {
		PageSize   int32 `json:"pageSize"`
		PageOffset int32 `json:"pageOffset"`
	} `json:"pagination"`
}

type APIStatus struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
