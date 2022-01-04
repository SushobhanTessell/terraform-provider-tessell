package models

type Workflow struct {
	CreatedBy          string      `json:"createdBy"`
	ExecutionEndTime   string      `json:"executionEndTime"`
	ExecutionStartTime string      `json:"executionStartTime"`
	Id                 string      `json:"id"`
	Status             string      `json:"status"`
	UpdateTime         string      `json:"updateTime"`
	WorkflowName       string      `json:"workflowName"`
	Data               interface{} `json:"data"`
}
