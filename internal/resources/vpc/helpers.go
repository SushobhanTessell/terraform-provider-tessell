package vpc

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	apiClient "terraform-provider-tessell/internal/client"
	"terraform-provider-tessell/internal/models"
)

func validate(d *schema.ResourceData) error {
	return nil
}

func formPayload(d *schema.ResourceData) (models.VPCCreationPayload, error) {
	payload := models.VPCCreationPayload{
		Name:             d.Get("name").(string),
		CidrBlock:        d.Get("cidr_block").(string),
		CloudType:        d.Get("cloud_type").(string),
		Region:           d.Get("region").(string),
		SubscriptionName: d.Get("subscription_name").(string),
	}

	return payload, nil
}

func waitAndCheckIfCreationSucceeded(waitTime int, d *schema.ResourceData, meta interface{}) error {
	client := meta.(*apiClient.Client)

	loopCount := 0
	sleepCycleDuration, err := time.ParseDuration("10s")
	if err != nil {
		return err
	}
	loops := waitTime / int(sleepCycleDuration.Seconds())

	var vpcStatus string
	for {
		vpc, err := client.GetVPC(d.Get("name").(string), d.Get("cloud_type").(string), d.Get("region").(string), d.Get("subscription_name").(string))
		if err != nil {
			return err
		}
		vpcStatus = vpc.Status

		if vpcStatus == "ACTIVE" {
			return nil
		} else if vpcStatus == "CREATION_FAILED" {
			return fmt.Errorf("creation of VPC failed")
		} else if vpcStatus == "VALIDATION_FAILED" {
			return fmt.Errorf("validation for VPC creation failed")
		}
		loopCount = loopCount + 1
		if loopCount > loops {
			break
		}
		time.Sleep(sleepCycleDuration)
	}

	return fmt.Errorf("timed out waiting for VPC creation to complete; last known status: %s", vpcStatus)
}

func setResourceData(d *schema.ResourceData, vpc *models.VPC) error {
	if err := d.Set("name", vpc.Name); err != nil {
		return err
	}

	if err := d.Set("cidr_block", vpc.CidrBlock); err != nil {
		return err
	}

	if err := d.Set("cloud_type", vpc.CloudType); err != nil {
		return err
	}

	if err := d.Set("region", vpc.Region); err != nil {
		return err
	}

	if err := d.Set("status", vpc.Status); err != nil {
		return err
	}

	if err := d.Set("subscription_name", vpc.SubscriptionName); err != nil {
		return err
	}

	if err := d.Set("metadata", parseMetadata(vpc)); err != nil {
		return err
	}

	return nil
}

func parseMetadata(vpc *models.VPC) []interface{} {
	return []interface{}{map[string]interface{}{
		"validation_failure_reason": vpc.Metadata.ValidationFailureReason,
	}}
}
