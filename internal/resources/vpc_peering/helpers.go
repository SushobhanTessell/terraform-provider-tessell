package vpc_peering

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-tessell/internal/models"
)

func validate(d *schema.ResourceData) error {
	return nil
}

func formPayload(d *schema.ResourceData) (models.VPCPeeringCreationPayload, error) {
	cloudType := d.Get("cloud_type").(string)

	payload := models.VPCPeeringCreationPayload{
		SubscriptionName: d.Get("subscription_name").(string),
		CloudType:        d.Get("cloud_type").(string),
		Region:           d.Get("region").(string),
	}
	if cloudType == "AWS" {
		awsClientVpcInfo := ((d.Get("aws_client_vpc_info").([]interface{}))[0]).(map[string]interface{})
		payload.AWSClientVpcInfo = models.AWSClientInfo{
			ClientAccountId: awsClientVpcInfo["client_account_id"].(string),
			ClientVpcId:     awsClientVpcInfo["client_vpc_id"].(string),
			ClientVpcRegion: awsClientVpcInfo["client_vpc_region"].(string),
		}
	} else if cloudType == "AZURE" {
		azureClientVpcInfo := ((d.Get("azure_client_vpc_info").([]interface{}))[0]).(map[string]interface{})
		payload.AzureClientVpcInfo = models.AzureClientVpcInfo{
			ClientSubscriptionId:          azureClientVpcInfo["client_subscription_id"].(string),
			ClientResourceGroup:           azureClientVpcInfo["client_resource_group"].(string),
			ClientVpcName:                 azureClientVpcInfo["client_vpc_name"].(string),
			ClientActiveDirectoryTenantId: azureClientVpcInfo["client_active_directory_tenant_id"].(string),
			ClientApplicationObjectId:     azureClientVpcInfo["client_application_object_id"].(string),
		}
	}

	return payload, nil
}

func setResourceData(d *schema.ResourceData, vpcPeering *models.VPCPeering) error {
	if err := d.Set("name", vpcPeering.Name); err != nil {
		return err
	}

	if err := d.Set("cloud_id", vpcPeering.CloudId); err != nil {
		return err
	}

	if err := d.Set("status", vpcPeering.Status); err != nil {
		return err
	}

	if err := d.Set("metadata", parseMetadata(vpcPeering)); err != nil {
		return err
	}

	cloudType := d.Get("cloud_type").(string)
	if cloudType == "AWS" {
		if err := d.Set("aws_client_info", parseAWSClientInfo(&vpcPeering.AWSClientInfo)); err != nil {
			return err
		}
	} else if cloudType == "AZURE" {
		if err := d.Set("azure_client_info", parseAzureClientInfo(&vpcPeering.AzureClientInfo)); err != nil {
			return err
		}
	}

	return nil
}

func parseMetadata(vpcPeering *models.VPCPeering) []interface{} {
	return []interface{}{map[string]interface{}{
		"azure_pending_peer_metadata": []interface{}{map[string]interface{}{
			"tenant_id":        vpcPeering.Metadata.AzurePendingPeerMetadata.TenantId,
			"vnet_resource_id": vpcPeering.Metadata.AzurePendingPeerMetadata.VNetResourceId,
		}},
	}}
}

func parseAWSClientInfo(awsClientInfo *models.AWSClientInfo) []interface{} {
	return []interface{}{map[string]interface{}{
		"client_account_id": awsClientInfo.ClientAccountId,
		"client_vpc_id":     awsClientInfo.ClientVpcId,
		"client_vpc_region": awsClientInfo.ClientVpcRegion,
	}}
}

func parseAzureClientInfo(azureClientInfo *models.AzureClientInfo) []interface{} {
	return []interface{}{map[string]interface{}{
		"client_subscription_id": azureClientInfo.ClientSubscriptionId,
		"client_resource_group":  azureClientInfo.ClientResourceGroup,
		"client_vpc_name":        azureClientInfo.ClientVpcName,
	}}
}
