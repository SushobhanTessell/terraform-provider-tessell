package dataflix_catalog

import (
	//"fmt"
	//"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-tessell/internal/model"
)

func setResourceData(d *schema.ResourceData, tessellDmmDataflixServiceView *model.TessellDmmDataflixServiceView) error {
	if err := d.Set("availability_machine_id", tessellDmmDataflixServiceView.AvailabilityMachineId); err != nil {
		return err
	}

	if err := d.Set("tessell_service_id", tessellDmmDataflixServiceView.TessellServiceId); err != nil {
		return err
	}

	if err := d.Set("service_name", tessellDmmDataflixServiceView.ServiceName); err != nil {
		return err
	}

	if err := d.Set("engine_type", tessellDmmDataflixServiceView.EngineType); err != nil {
		return err
	}

	if err := d.Set("time_zone", tessellDmmDataflixServiceView.TimeZone); err != nil {
		return err
	}

	if err := d.Set("owner", tessellDmmDataflixServiceView.Owner); err != nil {
		return err
	}

	if err := d.Set("pitr_catalog", parseTessellDataflixPitrInfoListWithResData(tessellDmmDataflixServiceView.PitrCatalog, d)); err != nil {
		return err
	}

	if err := d.Set("snapshot_catalog", parseTessellDmmDataflixBackupDTOListWithResData(tessellDmmDataflixServiceView.SnapshotCatalog, d)); err != nil {
		return err
	}

	return nil
}

func parseTessellDataflixPitrInfoListWithResData(pitrCatalog *[]model.TessellDataflixPitrInfo, d *schema.ResourceData) []interface{} {
	if pitrCatalog == nil {
		return nil
	}
	tessellDataflixPitrInfoList := make([]interface{}, 0)

	if pitrCatalog != nil {
		tessellDataflixPitrInfoList = make([]interface{}, len(*pitrCatalog))
		for i, tessellDataflixPitrInfoItem := range *pitrCatalog {
			tessellDataflixPitrInfoList[i] = parseTessellDataflixPitrInfo(&tessellDataflixPitrInfoItem)
		}
	}

	return tessellDataflixPitrInfoList
}

func parseTessellDataflixPitrInfo(pitrCatalog *model.TessellDataflixPitrInfo) interface{} {
	if pitrCatalog == nil {
		return nil
	}
	parsedPitrCatalog := make(map[string]interface{})
	parsedPitrCatalog["cloud"] = pitrCatalog.Cloud

	var regions *[]model.TessellDataflixPitrInfoForRegion
	if pitrCatalog.Regions != regions {
		parsedPitrCatalog["regions"] = parseTessellDataflixPitrInfoForRegionList(pitrCatalog.Regions)
	}

	return parsedPitrCatalog
}

func parseTessellDataflixPitrInfoForRegionList(regions *[]model.TessellDataflixPitrInfoForRegion) []interface{} {
	if regions == nil {
		return nil
	}
	tessellDataflixPitrInfoForRegionList := make([]interface{}, 0)

	if regions != nil {
		tessellDataflixPitrInfoForRegionList = make([]interface{}, len(*regions))
		for i, tessellDataflixPitrInfoForRegionItem := range *regions {
			tessellDataflixPitrInfoForRegionList[i] = parseTessellDataflixPitrInfoForRegion(&tessellDataflixPitrInfoForRegionItem)
		}
	}

	return tessellDataflixPitrInfoForRegionList
}

func parseTessellDataflixPitrInfoForRegion(regions *model.TessellDataflixPitrInfoForRegion) interface{} {
	if regions == nil {
		return nil
	}
	parsedRegions := make(map[string]interface{})
	parsedRegions["region"] = regions.Region

	var timeRanges *[]model.TessellDataflixFromTimeInfo
	if regions.TimeRanges != timeRanges {
		parsedRegions["time_ranges"] = parseTessellDataflixFromTimeInfoList(regions.TimeRanges)
	}

	return parsedRegions
}

func parseTessellDataflixFromTimeInfoList(timeRanges *[]model.TessellDataflixFromTimeInfo) []interface{} {
	if timeRanges == nil {
		return nil
	}
	tessellDataflixFromTimeInfoList := make([]interface{}, 0)

	if timeRanges != nil {
		tessellDataflixFromTimeInfoList = make([]interface{}, len(*timeRanges))
		for i, tessellDataflixFromTimeInfoItem := range *timeRanges {
			tessellDataflixFromTimeInfoList[i] = parseTessellDataflixFromTimeInfo(&tessellDataflixFromTimeInfoItem)
		}
	}

	return tessellDataflixFromTimeInfoList
}

func parseTessellDataflixFromTimeInfo(timeRanges *model.TessellDataflixFromTimeInfo) interface{} {
	if timeRanges == nil {
		return nil
	}
	parsedTimeRanges := make(map[string]interface{})
	parsedTimeRanges["from_time"] = timeRanges.FromTime
	parsedTimeRanges["to_time"] = timeRanges.ToTime

	var sharedWith *model.EntityAclSharingSummaryInfo
	if timeRanges.SharedWith != sharedWith {
		parsedTimeRanges["shared_with"] = []interface{}{parseEntityAclSharingSummaryInfo(timeRanges.SharedWith)}
	}

	return parsedTimeRanges
}

func parseEntityAclSharingSummaryInfo(entityAclSharingSummaryInfo *model.EntityAclSharingSummaryInfo) interface{} {
	if entityAclSharingSummaryInfo == nil {
		return nil
	}
	parsedEntityAclSharingSummaryInfo := make(map[string]interface{})
	parsedEntityAclSharingSummaryInfo["users"] = entityAclSharingSummaryInfo.Users

	return parsedEntityAclSharingSummaryInfo
}

func parseTessellDmmDataflixBackupDTOListWithResData(snapshotCatalog *[]model.TessellDmmDataflixBackupDTO, d *schema.ResourceData) []interface{} {
	if snapshotCatalog == nil {
		return nil
	}
	tessellDmmDataflixBackupDTOList := make([]interface{}, 0)

	if snapshotCatalog != nil {
		tessellDmmDataflixBackupDTOList = make([]interface{}, len(*snapshotCatalog))
		for i, tessellDmmDataflixBackupDTOItem := range *snapshotCatalog {
			tessellDmmDataflixBackupDTOList[i] = parseTessellDmmDataflixBackupDTO(&tessellDmmDataflixBackupDTOItem)
		}
	}

	return tessellDmmDataflixBackupDTOList
}

func parseTessellDmmDataflixBackupDTO(snapshotCatalog *model.TessellDmmDataflixBackupDTO) interface{} {
	if snapshotCatalog == nil {
		return nil
	}
	parsedSnapshotCatalog := make(map[string]interface{})
	parsedSnapshotCatalog["id"] = snapshotCatalog.Id
	parsedSnapshotCatalog["name"] = snapshotCatalog.Name
	parsedSnapshotCatalog["description"] = snapshotCatalog.Description
	parsedSnapshotCatalog["snapshot_time"] = snapshotCatalog.SnapshotTime
	parsedSnapshotCatalog["status"] = snapshotCatalog.Status
	parsedSnapshotCatalog["size"] = snapshotCatalog.Size
	parsedSnapshotCatalog["manual"] = snapshotCatalog.Manual

	var cloudAvailability *[]model.CloudRegionInfo1
	if snapshotCatalog.CloudAvailability != cloudAvailability {
		parsedSnapshotCatalog["cloud_availability"] = parseCloudRegionInfo1List(snapshotCatalog.CloudAvailability)
	}

	var databases *[]model.BackupDatabaseInfo
	if snapshotCatalog.Databases != databases {
		parsedSnapshotCatalog["databases"] = parseBackupDatabaseInfoList(snapshotCatalog.Databases)
	}

	var sharedWith *model.EntityAclSharingSummaryInfo
	if snapshotCatalog.SharedWith != sharedWith {
		parsedSnapshotCatalog["shared_with"] = []interface{}{parseEntityAclSharingSummaryInfo(snapshotCatalog.SharedWith)}
	}

	return parsedSnapshotCatalog
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

func parseBackupDatabaseInfoList(databases *[]model.BackupDatabaseInfo) []interface{} {
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
