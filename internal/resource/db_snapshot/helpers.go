package db_snapshot

import (
	//"fmt"
	//"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-tessell/internal/helper"
	"terraform-provider-tessell/internal/model"
)

func setResourceData(d *schema.ResourceData, tessellDmmDataflixBackupDTO *model.TessellDmmDataflixBackupDTO) error {
	if err := d.Set("id", tessellDmmDataflixBackupDTO.Id); err != nil {
		return err
	}

	if err := d.Set("name", tessellDmmDataflixBackupDTO.Name); err != nil {
		return err
	}

	if err := d.Set("description", tessellDmmDataflixBackupDTO.Description); err != nil {
		return err
	}

	if err := d.Set("snapshot_time", tessellDmmDataflixBackupDTO.SnapshotTime); err != nil {
		return err
	}

	if err := d.Set("status", tessellDmmDataflixBackupDTO.Status); err != nil {
		return err
	}

	if err := d.Set("size", tessellDmmDataflixBackupDTO.Size); err != nil {
		return err
	}

	if err := d.Set("manual", tessellDmmDataflixBackupDTO.Manual); err != nil {
		return err
	}

	if err := d.Set("cloud_availability", parseCloudRegionInfo1ListWithResData(tessellDmmDataflixBackupDTO.CloudAvailability, d)); err != nil {
		return err
	}

	if err := d.Set("databases", parseBackupDatabaseInfoListWithResData(tessellDmmDataflixBackupDTO.Databases, d)); err != nil {
		return err
	}

	if err := d.Set("shared_with", parseEntityAclSharingSummaryInfoWithResData(tessellDmmDataflixBackupDTO.SharedWith, d)); err != nil {
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

func parseRegionInfoList(regions *[]model.RegionInfo) []interface{} {
	if regions == nil {
		return nil
	}
	regionInfoList := make([]interface{}, 0)

	if regions != nil {
		regionInfoList = make([]interface{}, len(*regions))
		for i, regionInfoItem := range *regions {
			regionInfoList[i] = parseRegionInfo(&regionInfoItem)
		}
	}

	return regionInfoList
}

func parseRegionInfo(regions *model.RegionInfo) interface{} {
	if regions == nil {
		return nil
	}
	parsedRegions := make(map[string]interface{})
	parsedRegions["region"] = regions.Region
	parsedRegions["availability_zones"] = regions.AvailabilityZones

	return parsedRegions
}

func parseBackupDatabaseInfoListWithResData(databases *[]model.BackupDatabaseInfo, d *schema.ResourceData) []interface{} {
	if databases == nil {
		return nil
	}
	backupDatabaseInfoList := make([]interface{}, 0)

	if databases != nil {
		backupDatabaseInfoList = make([]interface{}, len(*databases))
		for i, backupDatabaseInfoItem := range *databases {
			backupDatabaseInfoList[i] = parseBackupDatabaseInfo(&backupDatabaseInfoItem)
		}
	}

	return backupDatabaseInfoList
}

func parseBackupDatabaseInfo(databases *model.BackupDatabaseInfo) interface{} {
	if databases == nil {
		return nil
	}
	parsedDatabases := make(map[string]interface{})
	parsedDatabases["id"] = databases.Id
	parsedDatabases["name"] = databases.Name
	parsedDatabases["status"] = databases.Status

	return parsedDatabases
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

func formPayloadForCreateTessellServiceBackupRequest(d *schema.ResourceData) model.CreateBackupTaskPayload {
	createBackupTaskPayloadFormed := model.CreateBackupTaskPayload{
		Name:        helper.GetStringPointer(d.Get("name")),
		Description: helper.GetStringPointer(d.Get("description")),
	}

	return createBackupTaskPayloadFormed
}
