package models

type VPCs struct {
	VPCs []VPC
}

type VPC struct {
	Name             string `json:"name"`
	CidrBlock        string `json:"cidrBlock"`
	CloudType        string `json:"cloudType"`
	Region           string `json:"region"`
	Status           string `json:"status"`
	SubscriptionName string `json:"subscriptionName"`
	Metadata         struct {
		ValidationFailureReason string `json:"validationFailureReason"`
	} `json:"metadata"`
}

type VPCCreationPayload struct {
	Name             string `json:"name"`
	CidrBlock        string `json:"cidrBlock"`
	CloudType        string `json:"cloudType"`
	Region           string `json:"region"`
	SubscriptionName string `json:"subscriptionName"`
}

type VPCDeletionResponse struct {
}
