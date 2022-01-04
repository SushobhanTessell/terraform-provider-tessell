package models

type VPCPeerings struct {
	VPCPeerings []VPCPeering
}

type VPCPeering struct {
	Name            string          `json:"name"`
	CloudId         string          `json:"cloudId"`
	Status          string          `json:"status"`
	AWSClientInfo   AWSClientInfo   `json:"awsClientInfo"`
	AzureClientInfo AzureClientInfo `json:"azureClientInfo"`
	Metadata        struct {
		AzurePendingPeerMetadata struct {
			TenantId       string `json:"tenantId"`
			VNetResourceId string `json:"vnetResourceId"`
		} `json:"azurePendingPeerMetadata"`
	} `json:"metadata"`
}

type AWSClientInfo struct {
	ClientAccountId string `json:"clientAccountId"`
	ClientVpcId     string `json:"clientVpcId"`
	ClientVpcRegion string `json:"clientVpcRegion"`
}

type AzureClientInfo struct {
	ClientSubscriptionId string `json:"clientSubscriptionId"`
	ClientResourceGroup  string `json:"clientResourceGroup"`
	ClientVpcName        string `json:"clientVpcName"`
}

type AzureClientVpcInfo struct {
	ClientSubscriptionId          string `json:"clientSubscriptionId"`
	ClientResourceGroup           string `json:"clientResourceGroup"`
	ClientVpcName                 string `json:"clientVpcName"`
	ClientActiveDirectoryTenantId string `json:"clientActiveDirectoryTenantId"`
	ClientApplicationObjectId     string `json:"clientApplicationObjectId"`
}

type VPCPeeringCreationPayload struct {
	SubscriptionName   string             `json:"subscriptionName"`
	CloudType          string             `json:"cloudType"`
	Region             string             `json:"region"`
	AWSClientVpcInfo   AWSClientInfo      `json:"awsClientVpcInfo"`
	AzureClientVpcInfo AzureClientVpcInfo `json:"azureClientVpcInfo"`
}
