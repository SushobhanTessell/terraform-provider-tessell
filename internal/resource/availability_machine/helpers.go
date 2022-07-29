package availability_machine

import (
	//"fmt"
	//"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-tessell/internal/model"
)

func setResourceData(d *schema.ResourceData, tessellDmmServiceConsumerDTO *model.TessellDmmServiceConsumerDTO) error {
	if err := d.Set("id", tessellDmmServiceConsumerDTO.Id); err != nil {
		return err
	}

	if err := d.Set("tessell_service_id", tessellDmmServiceConsumerDTO.TessellServiceId); err != nil {
		return err
	}

	if err := d.Set("service_name", tessellDmmServiceConsumerDTO.ServiceName); err != nil {
		return err
	}

	if err := d.Set("tenant", tessellDmmServiceConsumerDTO.Tenant); err != nil {
		return err
	}

	if err := d.Set("subscription", tessellDmmServiceConsumerDTO.Subscription); err != nil {
		return err
	}

	if err := d.Set("engine_type", tessellDmmServiceConsumerDTO.EngineType); err != nil {
		return err
	}

	if err := d.Set("status", tessellDmmServiceConsumerDTO.Status); err != nil {
		return err
	}

	if err := d.Set("user_id", tessellDmmServiceConsumerDTO.UserId); err != nil {
		return err
	}

	if err := d.Set("owner", tessellDmmServiceConsumerDTO.Owner); err != nil {
		return err
	}

	if err := d.Set("logged_in_user_role", tessellDmmServiceConsumerDTO.LoggedInUserRole); err != nil {
		return err
	}

	if err := d.Set("shared_with", parseEntityAclSharingInfoWithResData(tessellDmmServiceConsumerDTO.SharedWith, d)); err != nil {
		return err
	}

	if err := d.Set("cloud_availability", parseCloudRegionInfo1ListWithResData(tessellDmmServiceConsumerDTO.CloudAvailability, d)); err != nil {
		return err
	}

	if err := d.Set("rpo_sla", parseTessellDmmAvailabilityServiceViewWithResData(tessellDmmServiceConsumerDTO.RpoSla, d)); err != nil {
		return err
	}

	if err := d.Set("daps", parseTessellDapServiceDTOListWithResData(tessellDmmServiceConsumerDTO.Daps, d)); err != nil {
		return err
	}

	if err := d.Set("clones", parseTessellCloneSummaryInfoListWithResData(tessellDmmServiceConsumerDTO.Clones, d)); err != nil {
		return err
	}

	if err := d.Set("date_created", tessellDmmServiceConsumerDTO.DateCreated); err != nil {
		return err
	}

	if err := d.Set("date_modified", tessellDmmServiceConsumerDTO.DateModified); err != nil {
		return err
	}

	return nil
}

func parseEntityAclSharingInfoWithResData(sharedWith *model.EntityAclSharingInfo, d *schema.ResourceData) []interface{} {
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

	var users *[]model.EntityUserAclSharingInfo
	if sharedWith.Users != users {
		parsedSharedWith["users"] = parseEntityUserAclSharingInfoList(sharedWith.Users)
	}

	return []interface{}{parsedSharedWith}
}

func parseEntityAclSharingInfo(sharedWith *model.EntityAclSharingInfo) interface{} {
	if sharedWith == nil {
		return nil
	}
	parsedSharedWith := make(map[string]interface{})

	var users *[]model.EntityUserAclSharingInfo
	if sharedWith.Users != users {
		parsedSharedWith["users"] = parseEntityUserAclSharingInfoList(sharedWith.Users)
	}

	return parsedSharedWith
}

func parseEntityUserAclSharingInfoList(entityUserAclSharingInfo *[]model.EntityUserAclSharingInfo) []interface{} {
	if entityUserAclSharingInfo == nil {
		return nil
	}
	entityUserAclSharingInfoList := make([]interface{}, 0)

	if entityUserAclSharingInfo != nil {
		entityUserAclSharingInfoList = make([]interface{}, len(*entityUserAclSharingInfo))
		for i, entityUserAclSharingInfoItem := range *entityUserAclSharingInfo {
			entityUserAclSharingInfoList[i] = parseEntityUserAclSharingInfo(&entityUserAclSharingInfoItem)
		}
	}

	return entityUserAclSharingInfoList
}

func parseEntityUserAclSharingInfo(entityUserAclSharingInfo *model.EntityUserAclSharingInfo) interface{} {
	if entityUserAclSharingInfo == nil {
		return nil
	}
	parsedEntityUserAclSharingInfo := make(map[string]interface{})
	parsedEntityUserAclSharingInfo["email_id"] = entityUserAclSharingInfo.EmailId
	parsedEntityUserAclSharingInfo["role"] = entityUserAclSharingInfo.Role

	return parsedEntityUserAclSharingInfo
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

func parseTessellDmmAvailabilityServiceViewWithResData(rpoSla *model.TessellDmmAvailabilityServiceView, d *schema.ResourceData) []interface{} {
	if rpoSla == nil {
		return nil
	}
	parsedRpoSla := make(map[string]interface{})
	if d.Get("rpo_sla") != nil {
		rpoSlaResourceData := d.Get("rpo_sla").([]interface{})
		if len(rpoSlaResourceData) > 0 {
			parsedRpoSla = (rpoSlaResourceData[0]).(map[string]interface{})
		}
	}
	parsedRpoSla["availability_machine_id"] = rpoSla.AvailabilityMachineId
	parsedRpoSla["availability_machine"] = rpoSla.AvailabilityMachine

	parsedRpoSla["rpo_sla_status"] = rpoSla.RpoSlaStatus
	parsedRpoSla["sla"] = rpoSla.Sla

	var cloudAvailability *[]model.CloudRegionInfo1
	if rpoSla.CloudAvailability != cloudAvailability {
		parsedRpoSla["cloud_availability"] = parseCloudRegionInfo1List(rpoSla.CloudAvailability)
	}

	var schedule *model.ScheduleInfo
	if rpoSla.Schedule != schedule {
		parsedRpoSla["schedule"] = []interface{}{parseScheduleInfo(rpoSla.Schedule)}
	}

	return []interface{}{parsedRpoSla}
}

func parseTessellDmmAvailabilityServiceView(rpoSla *model.TessellDmmAvailabilityServiceView) interface{} {
	if rpoSla == nil {
		return nil
	}
	parsedRpoSla := make(map[string]interface{})
	parsedRpoSla["availability_machine_id"] = rpoSla.AvailabilityMachineId
	parsedRpoSla["availability_machine"] = rpoSla.AvailabilityMachine

	parsedRpoSla["rpo_sla_status"] = rpoSla.RpoSlaStatus
	parsedRpoSla["sla"] = rpoSla.Sla

	var cloudAvailability *[]model.CloudRegionInfo1
	if rpoSla.CloudAvailability != cloudAvailability {
		parsedRpoSla["cloud_availability"] = parseCloudRegionInfo1List(rpoSla.CloudAvailability)
	}

	var schedule *model.ScheduleInfo
	if rpoSla.Schedule != schedule {
		parsedRpoSla["schedule"] = []interface{}{parseScheduleInfo(rpoSla.Schedule)}
	}

	return parsedRpoSla
}

func parseScheduleInfo(scheduleInfo *model.ScheduleInfo) interface{} {
	if scheduleInfo == nil {
		return nil
	}
	parsedScheduleInfo := make(map[string]interface{})

	var backupStartTime *model.ScheduleTimeFormat
	if scheduleInfo.BackupStartTime != backupStartTime {
		parsedScheduleInfo["backup_start_time"] = []interface{}{parseScheduleTimeFormat(scheduleInfo.BackupStartTime)}
	}

	var dailySchedule *model.DailySchedule
	if scheduleInfo.DailySchedule != dailySchedule {
		parsedScheduleInfo["daily_schedule"] = []interface{}{parseDailySchedule(scheduleInfo.DailySchedule)}
	}

	return parsedScheduleInfo
}

func parseScheduleTimeFormatList(scheduleTimeFormat *[]model.ScheduleTimeFormat) []interface{} {
	if scheduleTimeFormat == nil {
		return nil
	}
	scheduleTimeFormatList := make([]interface{}, 0)

	if scheduleTimeFormat != nil {
		scheduleTimeFormatList = make([]interface{}, len(*scheduleTimeFormat))
		for i, scheduleTimeFormatItem := range *scheduleTimeFormat {
			scheduleTimeFormatList[i] = parseScheduleTimeFormat(&scheduleTimeFormatItem)
		}
	}

	return scheduleTimeFormatList
}

func parseScheduleTimeFormat(scheduleTimeFormat *model.ScheduleTimeFormat) interface{} {
	if scheduleTimeFormat == nil {
		return nil
	}
	parsedScheduleTimeFormat := make(map[string]interface{})
	parsedScheduleTimeFormat["hour"] = scheduleTimeFormat.Hour
	parsedScheduleTimeFormat["minute"] = scheduleTimeFormat.Minute

	return parsedScheduleTimeFormat
}

func parseDailySchedule(dailySchedule *model.DailySchedule) interface{} {
	if dailySchedule == nil {
		return nil
	}
	parsedDailySchedule := make(map[string]interface{})
	parsedDailySchedule["backups_per_day"] = dailySchedule.BackupsPerDay

	var backupStartTimes *[]model.ScheduleTimeFormat
	if dailySchedule.BackupStartTimes != backupStartTimes {
		parsedDailySchedule["backup_start_times"] = parseScheduleTimeFormatList(dailySchedule.BackupStartTimes)
	}

	return parsedDailySchedule
}

func parseTessellDapServiceDTOListWithResData(daps *[]model.TessellDapServiceDTO, d *schema.ResourceData) []interface{} {
	if daps == nil {
		return nil
	}
	tessellDapServiceDTOList := make([]interface{}, 0)

	if daps != nil {
		tessellDapServiceDTOList = make([]interface{}, len(*daps))
		for i, tessellDapServiceDTOItem := range *daps {
			tessellDapServiceDTOList[i] = parseTessellDapServiceDTO(&tessellDapServiceDTOItem)
		}
	}

	return tessellDapServiceDTOList
}

func parseTessellDapServiceDTOList(daps *[]model.TessellDapServiceDTO) []interface{} {
	if daps == nil {
		return nil
	}
	tessellDapServiceDTOList := make([]interface{}, 0)

	if daps != nil {
		tessellDapServiceDTOList = make([]interface{}, len(*daps))
		for i, tessellDapServiceDTOItem := range *daps {
			tessellDapServiceDTOList[i] = parseTessellDapServiceDTO(&tessellDapServiceDTOItem)
		}
	}

	return tessellDapServiceDTOList
}

func parseTessellDapServiceDTO(daps *model.TessellDapServiceDTO) interface{} {
	if daps == nil {
		return nil
	}
	parsedDaps := make(map[string]interface{})
	parsedDaps["id"] = daps.Id
	parsedDaps["name"] = daps.Name
	parsedDaps["availability_machine_id"] = daps.AvailabilityMachineId
	parsedDaps["tessell_service_id"] = daps.TessellServiceId
	parsedDaps["service_name"] = daps.ServiceName
	parsedDaps["engine_type"] = daps.EngineType
	parsedDaps["content_type"] = daps.ContentType
	parsedDaps["status"] = daps.Status

	parsedDaps["owner"] = daps.Owner
	parsedDaps["logged_in_user_role"] = daps.LoggedInUserRole

	parsedDaps["date_created"] = daps.DateCreated
	parsedDaps["date_modified"] = daps.DateModified

	var contentInfo *model.DapContentInfo
	if daps.ContentInfo != contentInfo {
		parsedDaps["content_info"] = []interface{}{parseDapContentInfo(daps.ContentInfo)}
	}

	var cloudAvailability *[]model.CloudRegionInfo1
	if daps.CloudAvailability != cloudAvailability {
		parsedDaps["cloud_availability"] = parseCloudRegionInfo1List(daps.CloudAvailability)
	}

	var dataAccessConfig *model.RetentionAndScheduleInfo
	if daps.DataAccessConfig != dataAccessConfig {
		parsedDaps["data_access_config"] = []interface{}{parseRetentionAndScheduleInfo(daps.DataAccessConfig)}
	}

	var sharedWith *model.EntityAclSharingInfo
	if daps.SharedWith != sharedWith {
		parsedDaps["shared_with"] = []interface{}{parseEntityAclSharingInfo(daps.SharedWith)}
	}

	return parsedDaps
}

func parseDapContentInfo(dapContentInfo *model.DapContentInfo) interface{} {
	if dapContentInfo == nil {
		return nil
	}
	parsedDapContentInfo := make(map[string]interface{})

	var sanitizedContent *model.SanitizationDapContent
	if dapContentInfo.SanitizedContent != sanitizedContent {
		parsedDapContentInfo["sanitized_content"] = []interface{}{parseSanitizationDapContent(dapContentInfo.SanitizedContent)}
	}

	return parsedDapContentInfo
}

func parseSanitizationDapContent(sanitizationDapContent *model.SanitizationDapContent) interface{} {
	if sanitizationDapContent == nil {
		return nil
	}
	parsedSanitizationDapContent := make(map[string]interface{})

	var automated *model.SanitizationDapContentAutomated
	if sanitizationDapContent.Automated != automated {
		parsedSanitizationDapContent["automated"] = []interface{}{parseSanitizationDapContentAutomated(sanitizationDapContent.Automated)}
	}

	return parsedSanitizationDapContent
}

func parseSanitizationDapContentAutomated(sanitizationDapContent_automated *model.SanitizationDapContentAutomated) interface{} {
	if sanitizationDapContent_automated == nil {
		return nil
	}
	parsedSanitizationDapContent_automated := make(map[string]interface{})
	parsedSanitizationDapContent_automated["sanitization_schedule_id"] = sanitizationDapContent_automated.SanitizationScheduleId

	return parsedSanitizationDapContent_automated
}

func parseRetentionAndScheduleInfo(retentionAndScheduleInfo *model.RetentionAndScheduleInfo) interface{} {
	if retentionAndScheduleInfo == nil {
		return nil
	}
	parsedRetentionAndScheduleInfo := make(map[string]interface{})
	parsedRetentionAndScheduleInfo["daily_backups"] = retentionAndScheduleInfo.DailyBackups

	return parsedRetentionAndScheduleInfo
}

func parseTessellCloneSummaryInfoListWithResData(clones *[]model.TessellCloneSummaryInfo, d *schema.ResourceData) []interface{} {
	if clones == nil {
		return nil
	}
	tessellCloneSummaryInfoList := make([]interface{}, 0)

	if clones != nil {
		tessellCloneSummaryInfoList = make([]interface{}, len(*clones))
		for i, tessellCloneSummaryInfoItem := range *clones {
			tessellCloneSummaryInfoList[i] = parseTessellCloneSummaryInfo(&tessellCloneSummaryInfoItem)
		}
	}

	return tessellCloneSummaryInfoList
}

func parseTessellCloneSummaryInfoList(clones *[]model.TessellCloneSummaryInfo) []interface{} {
	if clones == nil {
		return nil
	}
	tessellCloneSummaryInfoList := make([]interface{}, 0)

	if clones != nil {
		tessellCloneSummaryInfoList = make([]interface{}, len(*clones))
		for i, tessellCloneSummaryInfoItem := range *clones {
			tessellCloneSummaryInfoList[i] = parseTessellCloneSummaryInfo(&tessellCloneSummaryInfoItem)
		}
	}

	return tessellCloneSummaryInfoList
}

func parseTessellCloneSummaryInfo(clones *model.TessellCloneSummaryInfo) interface{} {
	if clones == nil {
		return nil
	}
	parsedClones := make(map[string]interface{})
	parsedClones["id"] = clones.Id
	parsedClones["name"] = clones.Name
	parsedClones["subscription"] = clones.Subscription
	parsedClones["status"] = clones.Status

	parsedClones["clone_info"] = clones.CloneInfo
	parsedClones["owner"] = clones.Owner
	parsedClones["date_created"] = clones.DateCreated

	var cloudAvailability *[]model.CloudRegionInfo1
	if clones.CloudAvailability != cloudAvailability {
		parsedClones["cloud_availability"] = parseCloudRegionInfo1List(clones.CloudAvailability)
	}

	return parsedClones
}
