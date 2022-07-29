package dataflix

import (
	//"fmt"
	//"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-tessell/internal/model"
)

func setResourceData(d *schema.ResourceData, tessellAmDataflixDTO *model.TessellAmDataflixDTO) error {
	if err := d.Set("availability_machine_id", tessellAmDataflixDTO.AvailabilityMachineId); err != nil {
		return err
	}

	if err := d.Set("tessell_service_id", tessellAmDataflixDTO.TessellServiceId); err != nil {
		return err
	}

	if err := d.Set("service_name", tessellAmDataflixDTO.ServiceName); err != nil {
		return err
	}

	if err := d.Set("engine_type", tessellAmDataflixDTO.EngineType); err != nil {
		return err
	}

	if err := d.Set("cloud_availability", parseCloudRegionInfo1ListWithResData(tessellAmDataflixDTO.CloudAvailability, d)); err != nil {
		return err
	}

	if err := d.Set("owner", tessellAmDataflixDTO.Owner); err != nil {
		return err
	}

	if err := d.Set("shared_with", parseEntityAclSharingSummaryInfoWithResData(tessellAmDataflixDTO.SharedWith, d)); err != nil {
		return err
	}

	return nil
}

func parseCloudRegionInfo1ListWithResData(cloudAvailability *[]model.CloudRegionInfo1, d *schema.ResourceData) []interface{} {
	if cloudAvailability == nil {
		return nil
	}
	cloudRegionInfo1List := make([]interface{}, 0)

	if cloudAvailability != nil {
		cloudRegionInfo1List = make([]interface{}, len(*cloudAvailability))
		for i, cloudRegionInfo1Item := range *cloudAvailability {
			cloudRegionInfo1List[i] = parseCloudRegionInfo1(&cloudRegionInfo1Item)
		}
	}

	return cloudRegionInfo1List
}

func parseCloudRegionInfo1List(cloudAvailability *[]model.CloudRegionInfo1) []interface{} {
	if cloudAvailability == nil {
		return nil
	}
	cloudRegionInfo1List := make([]interface{}, 0)

	if cloudAvailability != nil {
		cloudRegionInfo1List = make([]interface{}, len(*cloudAvailability))
		for i, cloudRegionInfo1Item := range *cloudAvailability {
			cloudRegionInfo1List[i] = parseCloudRegionInfo1(&cloudRegionInfo1Item)
		}
	}

	return cloudRegionInfo1List
}

func parseCloudRegionInfo1(cloudAvailability *model.CloudRegionInfo1) interface{} {
	if cloudAvailability == nil {
		return nil
	}
	parsedCloudAvailability := make(map[string]interface{})
	parsedCloudAvailability["cloud"] = cloudAvailability.Cloud

	var regions *[]model.RegionInfo
	if cloudAvailability.Regions != regions {
		parsedCloudAvailability["regions"] = parseRegionInfoList(cloudAvailability.Regions)
	}

	return parsedCloudAvailability
}

func parseRegionInfoList(regionInfo *[]model.RegionInfo) []interface{} {
	if regionInfo == nil {
		return nil
	}
	regionInfoList := make([]interface{}, 0)

	if regionInfo != nil {
		regionInfoList = make([]interface{}, len(*regionInfo))
		for i, regionInfoItem := range *regionInfo {
			regionInfoList[i] = parseRegionInfo(&regionInfoItem)
		}
	}

	return regionInfoList
}

func parseRegionInfo(regionInfo *model.RegionInfo) interface{} {
	if regionInfo == nil {
		return nil
	}
	parsedRegionInfo := make(map[string]interface{})
	parsedRegionInfo["region"] = regionInfo.Region
	parsedRegionInfo["availability_zones"] = regionInfo.AvailabilityZones

	return parsedRegionInfo
}

func parseEntityAclSharingSummaryInfoWithResData(sharedWith *model.EntityAclSharingSummaryInfo, d *schema.ResourceData) []interface{} {
	if sharedWith == nil {
		return nil
	}
	parsedSharedWith := make(map[string]interface{})
	if d.Get("shared_with") != nil {
		sharedWithResourceData := d.Get("shared_with").([]interface{})
		if len(sharedWithResourceData) > 0 {
			parsedSharedWith = (sharedWithResourceData[0]).(map[string]interface{})
		}
	}
	parsedSharedWith["users"] = sharedWith.Users

	return []interface{}{parsedSharedWith}
}

func parseEntityAclSharingSummaryInfo(sharedWith *model.EntityAclSharingSummaryInfo) interface{} {
	if sharedWith == nil {
		return nil
	}
	parsedSharedWith := make(map[string]interface{})
	parsedSharedWith["users"] = sharedWith.Users

	return parsedSharedWith
}
