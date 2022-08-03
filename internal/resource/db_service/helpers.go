package db_service

import (
	//"fmt"
	//"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-tessell/internal/helper"
	"terraform-provider-tessell/internal/model"
)

func setResourceData(d *schema.ResourceData, tessellServiceDTO *model.TessellServiceDTO) error {
	if err := d.Set("id", tessellServiceDTO.Id); err != nil {
		return err
	}

	if err := d.Set("availability_machine_id", tessellServiceDTO.AvailabilityMachineId); err != nil {
		return err
	}

	if err := d.Set("name", tessellServiceDTO.Name); err != nil {
		return err
	}

	if err := d.Set("description", tessellServiceDTO.Description); err != nil {
		return err
	}

	if err := d.Set("engine_type", tessellServiceDTO.EngineType); err != nil {
		return err
	}

	if err := d.Set("topology", tessellServiceDTO.Topology); err != nil {
		return err
	}

	if err := d.Set("num_of_instances", tessellServiceDTO.NumOfInstances); err != nil {
		return err
	}

	if err := d.Set("status", tessellServiceDTO.Status); err != nil {
		return err
	}

	if err := d.Set("license_type", tessellServiceDTO.LicenseType); err != nil {
		return err
	}

	if err := d.Set("auto_minor_version_update", tessellServiceDTO.AutoMinorVersionUpdate); err != nil {
		return err
	}

	if err := d.Set("enable_deletion_protection", tessellServiceDTO.EnableDeletionProtection); err != nil {
		return err
	}

	if err := d.Set("software_image", tessellServiceDTO.SoftwareImage); err != nil {
		return err
	}

	if err := d.Set("software_image_version", tessellServiceDTO.SoftwareImageVersion); err != nil {
		return err
	}

	if err := d.Set("tenant_id", tessellServiceDTO.TenantId); err != nil {
		return err
	}

	if err := d.Set("subscription", tessellServiceDTO.Subscription); err != nil {
		return err
	}

	if err := d.Set("owner", tessellServiceDTO.Owner); err != nil {
		return err
	}

	if err := d.Set("logged_in_user_role", tessellServiceDTO.LoggedInUserRole); err != nil {
		return err
	}

	if err := d.Set("date_created", tessellServiceDTO.DateCreated); err != nil {
		return err
	}

	if err := d.Set("started_at", tessellServiceDTO.StartedAt); err != nil {
		return err
	}

	if err := d.Set("stopped_at", tessellServiceDTO.StoppedAt); err != nil {
		return err
	}

	if err := d.Set("cloned_from_info", parseTessellServiceClonedFromInfoWithResData(tessellServiceDTO.ClonedFromInfo, d)); err != nil {
		return err
	}

	if err := d.Set("service_connectivity", parseTessellServiceConnectivityInfoWithResData(tessellServiceDTO.ServiceConnectivity, d)); err != nil {
		return err
	}

	if err := d.Set("infrastructure", parseTessellServiceInfrastructureInfoWithResData(tessellServiceDTO.Infrastructure, d)); err != nil {
		return err
	}

	if err := d.Set("maintenance_window", parseTessellServiceMaintenanceWindowWithResData(tessellServiceDTO.MaintenanceWindow, d)); err != nil {
		return err
	}

	if err := d.Set("engine_configuration", parseTessellServiceEngineInfoWithResData(tessellServiceDTO.EngineConfiguration, d)); err != nil {
		return err
	}

	if err := d.Set("integrations_config", parseTessellServiceIntegrationsInfoWithResData(tessellServiceDTO.IntegrationsConfig, d)); err != nil {
		return err
	}

	if err := d.Set("deletion_config", parseTessellServiceDeletionConfigWithResData(tessellServiceDTO.DeletionConfig, d)); err != nil {
		return err
	}

	if err := d.Set("tags", parseTessellTagListWithResData(tessellServiceDTO.Tags, d)); err != nil {
		return err
	}

	if err := d.Set("instances", parseTessellServiceInstanceDTOListWithResData(tessellServiceDTO.Instances, d)); err != nil {
		return err
	}

	if err := d.Set("databases", parseTessellDatabaseDTOListWithResData(tessellServiceDTO.Databases, d)); err != nil {
		return err
	}

	if err := d.Set("shared_with", parseEntityAclSharingInfoWithResData(tessellServiceDTO.SharedWith, d)); err != nil {
		return err
	}

	return nil
}

func parseTessellServiceClonedFromInfoWithResData(clonedFromInfo *model.TessellServiceClonedFromInfo, d *schema.ResourceData) []interface{} {
	if clonedFromInfo == nil {
		return nil
	}
	parsedClonedFromInfo := make(map[string]interface{})
	if d.Get("cloned_from_info") != nil {
		clonedFromInfoResourceData := d.Get("cloned_from_info").([]interface{})
		if len(clonedFromInfoResourceData) > 0 {
			parsedClonedFromInfo = (clonedFromInfoResourceData[0]).(map[string]interface{})
		}
	}
	parsedClonedFromInfo["tessell_service_id"] = clonedFromInfo.TessellServiceId
	parsedClonedFromInfo["availability_machine_id"] = clonedFromInfo.AvailabilityMachineId
	parsedClonedFromInfo["tessell_service"] = clonedFromInfo.TessellService
	parsedClonedFromInfo["availability_machine"] = clonedFromInfo.AvailabilityMachine
	parsedClonedFromInfo["snapshot_name"] = clonedFromInfo.SnapshotName
	parsedClonedFromInfo["snapshot_id"] = clonedFromInfo.SnapshotId
	parsedClonedFromInfo["pitr_time"] = clonedFromInfo.PitrTime

	return []interface{}{parsedClonedFromInfo}
}

func parseTessellServiceClonedFromInfo(clonedFromInfo *model.TessellServiceClonedFromInfo) interface{} {
	if clonedFromInfo == nil {
		return nil
	}
	parsedClonedFromInfo := make(map[string]interface{})
	parsedClonedFromInfo["tessell_service_id"] = clonedFromInfo.TessellServiceId
	parsedClonedFromInfo["availability_machine_id"] = clonedFromInfo.AvailabilityMachineId
	parsedClonedFromInfo["tessell_service"] = clonedFromInfo.TessellService
	parsedClonedFromInfo["availability_machine"] = clonedFromInfo.AvailabilityMachine
	parsedClonedFromInfo["snapshot_name"] = clonedFromInfo.SnapshotName
	parsedClonedFromInfo["snapshot_id"] = clonedFromInfo.SnapshotId
	parsedClonedFromInfo["pitr_time"] = clonedFromInfo.PitrTime

	return parsedClonedFromInfo
}

func parseTessellServiceConnectivityInfoWithResData(serviceConnectivity *model.TessellServiceConnectivityInfo, d *schema.ResourceData) []interface{} {
	if serviceConnectivity == nil {
		return nil
	}
	parsedServiceConnectivity := make(map[string]interface{})
	if d.Get("service_connectivity") != nil {
		serviceConnectivityResourceData := d.Get("service_connectivity").([]interface{})
		if len(serviceConnectivityResourceData) > 0 {
			parsedServiceConnectivity = (serviceConnectivityResourceData[0]).(map[string]interface{})
		}
	}
	parsedServiceConnectivity["dns_prefix"] = serviceConnectivity.DnsPrefix
	parsedServiceConnectivity["service_port"] = serviceConnectivity.ServicePort
	parsedServiceConnectivity["enable_public_access"] = serviceConnectivity.EnablePublicAccess
	parsedServiceConnectivity["allowed_ip_addresses"] = serviceConnectivity.AllowedIpAddresses

	var connectStrings *[]model.TessellServiceConnectString
	if serviceConnectivity.ConnectStrings != connectStrings {
		parsedServiceConnectivity["connect_strings"] = parseTessellServiceConnectStringList(serviceConnectivity.ConnectStrings)
	}

	var updateInProgressInfo *model.TessellServiceConnectivityUpdateInProgressInfo
	if serviceConnectivity.UpdateInProgressInfo != updateInProgressInfo {
		parsedServiceConnectivity["update_in_progress_info"] = []interface{}{parseTessellServiceConnectivityUpdateInProgressInfo(serviceConnectivity.UpdateInProgressInfo)}
	}

	return []interface{}{parsedServiceConnectivity}
}

func parseTessellServiceConnectivityInfo(serviceConnectivity *model.TessellServiceConnectivityInfo) interface{} {
	if serviceConnectivity == nil {
		return nil
	}
	parsedServiceConnectivity := make(map[string]interface{})
	parsedServiceConnectivity["dns_prefix"] = serviceConnectivity.DnsPrefix
	parsedServiceConnectivity["service_port"] = serviceConnectivity.ServicePort
	parsedServiceConnectivity["enable_public_access"] = serviceConnectivity.EnablePublicAccess
	parsedServiceConnectivity["allowed_ip_addresses"] = serviceConnectivity.AllowedIpAddresses

	var connectStrings *[]model.TessellServiceConnectString
	if serviceConnectivity.ConnectStrings != connectStrings {
		parsedServiceConnectivity["connect_strings"] = parseTessellServiceConnectStringList(serviceConnectivity.ConnectStrings)
	}

	var updateInProgressInfo *model.TessellServiceConnectivityUpdateInProgressInfo
	if serviceConnectivity.UpdateInProgressInfo != updateInProgressInfo {
		parsedServiceConnectivity["update_in_progress_info"] = []interface{}{parseTessellServiceConnectivityUpdateInProgressInfo(serviceConnectivity.UpdateInProgressInfo)}
	}

	return parsedServiceConnectivity
}

func parseTessellServiceConnectStringList(tessellServiceConnectString *[]model.TessellServiceConnectString) []interface{} {
	if tessellServiceConnectString == nil {
		return nil
	}
	tessellServiceConnectStringList := make([]interface{}, 0)

	if tessellServiceConnectString != nil {
		tessellServiceConnectStringList = make([]interface{}, len(*tessellServiceConnectString))
		for i, tessellServiceConnectStringItem := range *tessellServiceConnectString {
			tessellServiceConnectStringList[i] = parseTessellServiceConnectString(&tessellServiceConnectStringItem)
		}
	}

	return tessellServiceConnectStringList
}

func parseTessellServiceConnectString(tessellServiceConnectString *model.TessellServiceConnectString) interface{} {
	if tessellServiceConnectString == nil {
		return nil
	}
	parsedTessellServiceConnectString := make(map[string]interface{})
	parsedTessellServiceConnectString["type"] = tessellServiceConnectString.Type
	parsedTessellServiceConnectString["connect_descriptor"] = tessellServiceConnectString.ConnectDescriptor
	parsedTessellServiceConnectString["endpoint"] = tessellServiceConnectString.Endpoint
	parsedTessellServiceConnectString["master_user"] = tessellServiceConnectString.MasterUser
	parsedTessellServiceConnectString["service_port"] = tessellServiceConnectString.ServicePort

	return parsedTessellServiceConnectString
}

func parseTessellServiceConnectivityUpdateInProgressInfo(tessellServiceConnectivityUpdateInProgressInfo *model.TessellServiceConnectivityUpdateInProgressInfo) interface{} {
	if tessellServiceConnectivityUpdateInProgressInfo == nil {
		return nil
	}
	parsedTessellServiceConnectivityUpdateInProgressInfo := make(map[string]interface{})
	parsedTessellServiceConnectivityUpdateInProgressInfo["dns_prefix"] = tessellServiceConnectivityUpdateInProgressInfo.DnsPrefix
	parsedTessellServiceConnectivityUpdateInProgressInfo["enable_public_access"] = tessellServiceConnectivityUpdateInProgressInfo.EnablePublicAccess
	parsedTessellServiceConnectivityUpdateInProgressInfo["allowed_ip_addresses"] = tessellServiceConnectivityUpdateInProgressInfo.AllowedIpAddresses

	return parsedTessellServiceConnectivityUpdateInProgressInfo
}

func parseTessellServiceInfrastructureInfoWithResData(infrastructure *model.TessellServiceInfrastructureInfo, d *schema.ResourceData) []interface{} {
	if infrastructure == nil {
		return nil
	}
	parsedInfrastructure := make(map[string]interface{})
	if d.Get("infrastructure") != nil {
		infrastructureResourceData := d.Get("infrastructure").([]interface{})
		if len(infrastructureResourceData) > 0 {
			parsedInfrastructure = (infrastructureResourceData[0]).(map[string]interface{})
		}
	}
	parsedInfrastructure["cloud"] = infrastructure.Cloud
	parsedInfrastructure["region"] = infrastructure.Region
	parsedInfrastructure["availability_zone"] = infrastructure.AvailabilityZone

	parsedInfrastructure["vpc"] = infrastructure.Vpc
	parsedInfrastructure["compute_type"] = infrastructure.ComputeType
	parsedInfrastructure["additional_storage"] = infrastructure.AdditionalStorage

	var cloudAvailability *[]model.CloudRegionInfo1
	if infrastructure.CloudAvailability != cloudAvailability {
		parsedInfrastructure["cloud_availability"] = parseCloudRegionInfo1List(infrastructure.CloudAvailability)
	}

	return []interface{}{parsedInfrastructure}
}

func parseTessellServiceInfrastructureInfo(infrastructure *model.TessellServiceInfrastructureInfo) interface{} {
	if infrastructure == nil {
		return nil
	}
	parsedInfrastructure := make(map[string]interface{})
	parsedInfrastructure["cloud"] = infrastructure.Cloud
	parsedInfrastructure["region"] = infrastructure.Region
	parsedInfrastructure["availability_zone"] = infrastructure.AvailabilityZone

	parsedInfrastructure["vpc"] = infrastructure.Vpc
	parsedInfrastructure["compute_type"] = infrastructure.ComputeType
	parsedInfrastructure["additional_storage"] = infrastructure.AdditionalStorage

	var cloudAvailability *[]model.CloudRegionInfo1
	if infrastructure.CloudAvailability != cloudAvailability {
		parsedInfrastructure["cloud_availability"] = parseCloudRegionInfo1List(infrastructure.CloudAvailability)
	}

	return parsedInfrastructure
}

func parseCloudRegionInfo1List(cloudRegionInfo_1 *[]model.CloudRegionInfo1) []interface{} {
	if cloudRegionInfo_1 == nil {
		return nil
	}
	cloudRegionInfo1List := make([]interface{}, 0)

	if cloudRegionInfo_1 != nil {
		cloudRegionInfo1List = make([]interface{}, len(*cloudRegionInfo_1))
		for i, cloudRegionInfo1Item := range *cloudRegionInfo_1 {
			cloudRegionInfo1List[i] = parseCloudRegionInfo1(&cloudRegionInfo1Item)
		}
	}

	return cloudRegionInfo1List
}

func parseCloudRegionInfo1(cloudRegionInfo_1 *model.CloudRegionInfo1) interface{} {
	if cloudRegionInfo_1 == nil {
		return nil
	}
	parsedCloudRegionInfo_1 := make(map[string]interface{})
	parsedCloudRegionInfo_1["cloud"] = cloudRegionInfo_1.Cloud

	var regions *[]model.RegionInfo
	if cloudRegionInfo_1.Regions != regions {
		parsedCloudRegionInfo_1["regions"] = parseRegionInfoList(cloudRegionInfo_1.Regions)
	}

	return parsedCloudRegionInfo_1
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

func parseTessellServiceMaintenanceWindowWithResData(maintenanceWindow *model.TessellServiceMaintenanceWindow, d *schema.ResourceData) []interface{} {
	if maintenanceWindow == nil {
		return nil
	}
	parsedMaintenanceWindow := make(map[string]interface{})
	if d.Get("maintenance_window") != nil {
		maintenanceWindowResourceData := d.Get("maintenance_window").([]interface{})
		if len(maintenanceWindowResourceData) > 0 {
			parsedMaintenanceWindow = (maintenanceWindowResourceData[0]).(map[string]interface{})
		}
	}
	parsedMaintenanceWindow["day"] = maintenanceWindow.Day
	parsedMaintenanceWindow["time"] = maintenanceWindow.Time
	parsedMaintenanceWindow["duration"] = maintenanceWindow.Duration

	return []interface{}{parsedMaintenanceWindow}
}

func parseTessellServiceMaintenanceWindow(maintenanceWindow *model.TessellServiceMaintenanceWindow) interface{} {
	if maintenanceWindow == nil {
		return nil
	}
	parsedMaintenanceWindow := make(map[string]interface{})
	parsedMaintenanceWindow["day"] = maintenanceWindow.Day
	parsedMaintenanceWindow["time"] = maintenanceWindow.Time
	parsedMaintenanceWindow["duration"] = maintenanceWindow.Duration

	return parsedMaintenanceWindow
}

func parseTessellServiceEngineInfoWithResData(engineConfiguration *model.TessellServiceEngineInfo, d *schema.ResourceData) []interface{} {
	if engineConfiguration == nil {
		return nil
	}
	parsedEngineConfiguration := make(map[string]interface{})
	if d.Get("engine_configuration") != nil {
		engineConfigurationResourceData := d.Get("engine_configuration").([]interface{})
		if len(engineConfigurationResourceData) > 0 {
			parsedEngineConfiguration = (engineConfigurationResourceData[0]).(map[string]interface{})
		}
	}

	var oracleConfig *model.TessellServiceOracleEngineConfig
	if engineConfiguration.OracleConfig != oracleConfig {
		parsedEngineConfiguration["oracle_config"] = []interface{}{parseTessellServiceOracleEngineConfig(engineConfiguration.OracleConfig)}
	}

	var postgresqlConfig *model.TessellServicePostgresqlEngineConfig
	if engineConfiguration.PostgresqlConfig != postgresqlConfig {
		parsedEngineConfiguration["postgresql_config"] = []interface{}{parseTessellServicePostgresqlEngineConfig(engineConfiguration.PostgresqlConfig)}
	}

	var mysqlConfig *model.TessellServiceMySqlEngineConfig
	if engineConfiguration.MysqlConfig != mysqlConfig {
		parsedEngineConfiguration["mysql_config"] = []interface{}{parseTessellServiceMySqlEngineConfig(engineConfiguration.MysqlConfig)}
	}

	var sqlServerConfig *model.TessellServiceSqlServerEngineConfig
	if engineConfiguration.SqlServerConfig != sqlServerConfig {
		parsedEngineConfiguration["sql_server_config"] = []interface{}{parseTessellServiceSqlServerEngineConfig(engineConfiguration.SqlServerConfig)}
	}

	var apacheKafkaConfig *model.TessellServiceApacheKafkaEngineConfig
	if engineConfiguration.ApacheKafkaConfig != apacheKafkaConfig {
		parsedEngineConfiguration["apache_kafka_config"] = []interface{}{parseTessellServiceApacheKafkaEngineConfig(engineConfiguration.ApacheKafkaConfig)}
	}

	var preScriptInfo *model.ScriptInfo
	if engineConfiguration.PreScriptInfo != preScriptInfo {
		parsedEngineConfiguration["pre_script_info"] = []interface{}{parseScriptInfo(engineConfiguration.PreScriptInfo)}
	}

	var postScriptInfo *model.ScriptInfo
	if engineConfiguration.PostScriptInfo != postScriptInfo {
		parsedEngineConfiguration["post_script_info"] = []interface{}{parseScriptInfo(engineConfiguration.PostScriptInfo)}
	}

	return []interface{}{parsedEngineConfiguration}
}

func parseTessellServiceEngineInfo(engineConfiguration *model.TessellServiceEngineInfo) interface{} {
	if engineConfiguration == nil {
		return nil
	}
	parsedEngineConfiguration := make(map[string]interface{})

	var oracleConfig *model.TessellServiceOracleEngineConfig
	if engineConfiguration.OracleConfig != oracleConfig {
		parsedEngineConfiguration["oracle_config"] = []interface{}{parseTessellServiceOracleEngineConfig(engineConfiguration.OracleConfig)}
	}

	var postgresqlConfig *model.TessellServicePostgresqlEngineConfig
	if engineConfiguration.PostgresqlConfig != postgresqlConfig {
		parsedEngineConfiguration["postgresql_config"] = []interface{}{parseTessellServicePostgresqlEngineConfig(engineConfiguration.PostgresqlConfig)}
	}

	var mysqlConfig *model.TessellServiceMySqlEngineConfig
	if engineConfiguration.MysqlConfig != mysqlConfig {
		parsedEngineConfiguration["mysql_config"] = []interface{}{parseTessellServiceMySqlEngineConfig(engineConfiguration.MysqlConfig)}
	}

	var sqlServerConfig *model.TessellServiceSqlServerEngineConfig
	if engineConfiguration.SqlServerConfig != sqlServerConfig {
		parsedEngineConfiguration["sql_server_config"] = []interface{}{parseTessellServiceSqlServerEngineConfig(engineConfiguration.SqlServerConfig)}
	}

	var apacheKafkaConfig *model.TessellServiceApacheKafkaEngineConfig
	if engineConfiguration.ApacheKafkaConfig != apacheKafkaConfig {
		parsedEngineConfiguration["apache_kafka_config"] = []interface{}{parseTessellServiceApacheKafkaEngineConfig(engineConfiguration.ApacheKafkaConfig)}
	}

	var preScriptInfo *model.ScriptInfo
	if engineConfiguration.PreScriptInfo != preScriptInfo {
		parsedEngineConfiguration["pre_script_info"] = []interface{}{parseScriptInfo(engineConfiguration.PreScriptInfo)}
	}

	var postScriptInfo *model.ScriptInfo
	if engineConfiguration.PostScriptInfo != postScriptInfo {
		parsedEngineConfiguration["post_script_info"] = []interface{}{parseScriptInfo(engineConfiguration.PostScriptInfo)}
	}

	return parsedEngineConfiguration
}

func parseTessellServiceOracleEngineConfig(tessellServiceOracleEngineConfig *model.TessellServiceOracleEngineConfig) interface{} {
	if tessellServiceOracleEngineConfig == nil {
		return nil
	}
	parsedTessellServiceOracleEngineConfig := make(map[string]interface{})
	parsedTessellServiceOracleEngineConfig["multi_tenant"] = tessellServiceOracleEngineConfig.MultiTenant
	parsedTessellServiceOracleEngineConfig["parameter_profile"] = tessellServiceOracleEngineConfig.ParameterProfile
	parsedTessellServiceOracleEngineConfig["options_profile"] = tessellServiceOracleEngineConfig.OptionsProfile
	parsedTessellServiceOracleEngineConfig["character_set"] = tessellServiceOracleEngineConfig.CharacterSet
	parsedTessellServiceOracleEngineConfig["national_character_set"] = tessellServiceOracleEngineConfig.NationalCharacterSet

	return parsedTessellServiceOracleEngineConfig
}

func parseTessellServicePostgresqlEngineConfig(tessellServicePostgresqlEngineConfig *model.TessellServicePostgresqlEngineConfig) interface{} {
	if tessellServicePostgresqlEngineConfig == nil {
		return nil
	}
	parsedTessellServicePostgresqlEngineConfig := make(map[string]interface{})
	parsedTessellServicePostgresqlEngineConfig["parameter_profile"] = tessellServicePostgresqlEngineConfig.ParameterProfile

	return parsedTessellServicePostgresqlEngineConfig
}

func parseTessellServiceMySqlEngineConfig(tessellServiceMySqlEngineConfig *model.TessellServiceMySqlEngineConfig) interface{} {
	if tessellServiceMySqlEngineConfig == nil {
		return nil
	}
	parsedTessellServiceMySqlEngineConfig := make(map[string]interface{})
	parsedTessellServiceMySqlEngineConfig["parameter_profile"] = tessellServiceMySqlEngineConfig.ParameterProfile

	return parsedTessellServiceMySqlEngineConfig
}

func parseTessellServiceSqlServerEngineConfig(tessellServiceSqlServerEngineConfig *model.TessellServiceSqlServerEngineConfig) interface{} {
	if tessellServiceSqlServerEngineConfig == nil {
		return nil
	}
	parsedTessellServiceSqlServerEngineConfig := make(map[string]interface{})
	parsedTessellServiceSqlServerEngineConfig["parameter_profile"] = tessellServiceSqlServerEngineConfig.ParameterProfile

	return parsedTessellServiceSqlServerEngineConfig
}

func parseTessellServiceApacheKafkaEngineConfig(tessellServiceApacheKafkaEngineConfig *model.TessellServiceApacheKafkaEngineConfig) interface{} {
	if tessellServiceApacheKafkaEngineConfig == nil {
		return nil
	}
	parsedTessellServiceApacheKafkaEngineConfig := make(map[string]interface{})
	parsedTessellServiceApacheKafkaEngineConfig["parameter_profile"] = tessellServiceApacheKafkaEngineConfig.ParameterProfile

	return parsedTessellServiceApacheKafkaEngineConfig
}

func parseScriptInfo(scriptInfo *model.ScriptInfo) interface{} {
	if scriptInfo == nil {
		return nil
	}
	parsedScriptInfo := make(map[string]interface{})
	parsedScriptInfo["script_id"] = scriptInfo.ScriptId
	parsedScriptInfo["script_version"] = scriptInfo.ScriptVersion

	return parsedScriptInfo
}

func parseTessellServiceIntegrationsInfoWithResData(integrationsConfig *model.TessellServiceIntegrationsInfo, d *schema.ResourceData) []interface{} {
	if integrationsConfig == nil {
		return nil
	}
	parsedIntegrationsConfig := make(map[string]interface{})
	if d.Get("integrations_config") != nil {
		integrationsConfigResourceData := d.Get("integrations_config").([]interface{})
		if len(integrationsConfigResourceData) > 0 {
			parsedIntegrationsConfig = (integrationsConfigResourceData[0]).(map[string]interface{})
		}
	}
	parsedIntegrationsConfig["integrations"] = integrationsConfig.Integrations

	return []interface{}{parsedIntegrationsConfig}
}

func parseTessellServiceIntegrationsInfo(integrationsConfig *model.TessellServiceIntegrationsInfo) interface{} {
	if integrationsConfig == nil {
		return nil
	}
	parsedIntegrationsConfig := make(map[string]interface{})
	parsedIntegrationsConfig["integrations"] = integrationsConfig.Integrations

	return parsedIntegrationsConfig
}

func parseTessellServiceDeletionConfigWithResData(deletionConfig *model.TessellServiceDeletionConfig, d *schema.ResourceData) []interface{} {
	if deletionConfig == nil {
		return nil
	}
	parsedDeletionConfig := make(map[string]interface{})
	if d.Get("deletion_config") != nil {
		deletionConfigResourceData := d.Get("deletion_config").([]interface{})
		if len(deletionConfigResourceData) > 0 {
			parsedDeletionConfig = (deletionConfigResourceData[0]).(map[string]interface{})
		}
	}
	parsedDeletionConfig["retain_availability_machine"] = deletionConfig.RetainAvailabilityMachine

	return []interface{}{parsedDeletionConfig}
}

func parseTessellServiceDeletionConfig(deletionConfig *model.TessellServiceDeletionConfig) interface{} {
	if deletionConfig == nil {
		return nil
	}
	parsedDeletionConfig := make(map[string]interface{})
	parsedDeletionConfig["retain_availability_machine"] = deletionConfig.RetainAvailabilityMachine

	return parsedDeletionConfig
}

func parseTessellTagListWithResData(tags *[]model.TessellTag, d *schema.ResourceData) []interface{} {
	if tags == nil {
		return nil
	}
	tessellTagList := make([]interface{}, 0)

	if tags != nil {
		tessellTagList = make([]interface{}, len(*tags))
		for i, tessellTagItem := range *tags {
			tessellTagList[i] = parseTessellTag(&tessellTagItem)
		}
	}

	return tessellTagList
}

func parseTessellTagList(tags *[]model.TessellTag) []interface{} {
	if tags == nil {
		return nil
	}
	tessellTagList := make([]interface{}, 0)

	if tags != nil {
		tessellTagList = make([]interface{}, len(*tags))
		for i, tessellTagItem := range *tags {
			tessellTagList[i] = parseTessellTag(&tessellTagItem)
		}
	}

	return tessellTagList
}

func parseTessellTag(tags *model.TessellTag) interface{} {
	if tags == nil {
		return nil
	}
	parsedTags := make(map[string]interface{})
	parsedTags["name"] = tags.Name
	parsedTags["value"] = tags.Value

	return parsedTags
}

func parseTessellServiceInstanceDTOListWithResData(instances *[]model.TessellServiceInstanceDTO, d *schema.ResourceData) []interface{} {
	if instances == nil {
		return nil
	}
	tessellServiceInstanceDTOList := make([]interface{}, 0)

	if instances != nil {
		tessellServiceInstanceDTOList = make([]interface{}, len(*instances))
		for i, tessellServiceInstanceDTOItem := range *instances {
			tessellServiceInstanceDTOList[i] = parseTessellServiceInstanceDTO(&tessellServiceInstanceDTOItem)
		}
	}

	return tessellServiceInstanceDTOList
}

func parseTessellServiceInstanceDTOList(instances *[]model.TessellServiceInstanceDTO) []interface{} {
	if instances == nil {
		return nil
	}
	tessellServiceInstanceDTOList := make([]interface{}, 0)

	if instances != nil {
		tessellServiceInstanceDTOList = make([]interface{}, len(*instances))
		for i, tessellServiceInstanceDTOItem := range *instances {
			tessellServiceInstanceDTOList[i] = parseTessellServiceInstanceDTO(&tessellServiceInstanceDTOItem)
		}
	}

	return tessellServiceInstanceDTOList
}

func parseTessellServiceInstanceDTO(instances *model.TessellServiceInstanceDTO) interface{} {
	if instances == nil {
		return nil
	}
	parsedInstances := make(map[string]interface{})
	parsedInstances["id"] = instances.Id
	parsedInstances["name"] = instances.Name
	parsedInstances["role"] = instances.Role
	parsedInstances["status"] = instances.Status
	parsedInstances["tessell_service_id"] = instances.TessellServiceId
	parsedInstances["compute_type"] = instances.ComputeType
	parsedInstances["cloud"] = instances.Cloud
	parsedInstances["region"] = instances.Region
	parsedInstances["availability_zone"] = instances.AvailabilityZone
	parsedInstances["date_created"] = instances.DateCreated

	var connectString *model.TessellServiceInstanceConnectString
	if instances.ConnectString != connectString {
		parsedInstances["connect_string"] = []interface{}{parseTessellServiceInstanceConnectString(instances.ConnectString)}
	}

	var updatesInProgress *[]model.TessellResourceUpdateInfo
	if instances.UpdatesInProgress != updatesInProgress {
		parsedInstances["updates_in_progress"] = parseTessellResourceUpdateInfoList(instances.UpdatesInProgress)
	}

	return parsedInstances
}

func parseTessellServiceInstanceConnectString(tessellServiceInstanceConnectString *model.TessellServiceInstanceConnectString) interface{} {
	if tessellServiceInstanceConnectString == nil {
		return nil
	}
	parsedTessellServiceInstanceConnectString := make(map[string]interface{})
	parsedTessellServiceInstanceConnectString["connect_descriptor"] = tessellServiceInstanceConnectString.ConnectDescriptor
	parsedTessellServiceInstanceConnectString["master_user"] = tessellServiceInstanceConnectString.MasterUser
	parsedTessellServiceInstanceConnectString["endpoint"] = tessellServiceInstanceConnectString.Endpoint
	parsedTessellServiceInstanceConnectString["service_port"] = tessellServiceInstanceConnectString.ServicePort

	return parsedTessellServiceInstanceConnectString
}

func parseTessellResourceUpdateInfoList(tessellResourceUpdateInfo *[]model.TessellResourceUpdateInfo) []interface{} {
	if tessellResourceUpdateInfo == nil {
		return nil
	}
	tessellResourceUpdateInfoList := make([]interface{}, 0)

	if tessellResourceUpdateInfo != nil {
		tessellResourceUpdateInfoList = make([]interface{}, len(*tessellResourceUpdateInfo))
		for i, tessellResourceUpdateInfoItem := range *tessellResourceUpdateInfo {
			tessellResourceUpdateInfoList[i] = parseTessellResourceUpdateInfo(&tessellResourceUpdateInfoItem)
		}
	}

	return tessellResourceUpdateInfoList
}

func parseTessellResourceUpdateInfo(tessellResourceUpdateInfo *model.TessellResourceUpdateInfo) interface{} {
	if tessellResourceUpdateInfo == nil {
		return nil
	}
	parsedTessellResourceUpdateInfo := make(map[string]interface{})
	parsedTessellResourceUpdateInfo["update_type"] = tessellResourceUpdateInfo.UpdateType
	parsedTessellResourceUpdateInfo["reference_id"] = tessellResourceUpdateInfo.ReferenceId
	parsedTessellResourceUpdateInfo["submitted_at"] = tessellResourceUpdateInfo.SubmittedAt
	parsedTessellResourceUpdateInfo["update_info"] = tessellResourceUpdateInfo.UpdateInfo

	return parsedTessellResourceUpdateInfo
}

func parseTessellDatabaseDTOListWithResData(databases *[]model.TessellDatabaseDTO, d *schema.ResourceData) []interface{} {
	if databases == nil {
		return nil
	}
	tessellDatabaseDTOList := make([]interface{}, 0)

	if databases != nil {
		tessellDatabaseDTOList = make([]interface{}, len(*databases))
		for i, tessellDatabaseDTOItem := range *databases {
			tessellDatabaseDTOList[i] = parseTessellDatabaseDTO(&tessellDatabaseDTOItem)
		}
	}

	return tessellDatabaseDTOList
}

func parseTessellDatabaseDTOList(databases *[]model.TessellDatabaseDTO) []interface{} {
	if databases == nil {
		return nil
	}
	tessellDatabaseDTOList := make([]interface{}, 0)

	if databases != nil {
		tessellDatabaseDTOList = make([]interface{}, len(*databases))
		for i, tessellDatabaseDTOItem := range *databases {
			tessellDatabaseDTOList[i] = parseTessellDatabaseDTO(&tessellDatabaseDTOItem)
		}
	}

	return tessellDatabaseDTOList
}

func parseTessellDatabaseDTO(databases *model.TessellDatabaseDTO) interface{} {
	if databases == nil {
		return nil
	}
	parsedDatabases := make(map[string]interface{})
	parsedDatabases["id"] = databases.Id
	parsedDatabases["database_name"] = databases.DatabaseName
	parsedDatabases["description"] = databases.Description
	parsedDatabases["tessell_service_id"] = databases.TessellServiceId
	parsedDatabases["engine_type"] = databases.EngineType
	parsedDatabases["status"] = databases.Status
	parsedDatabases["date_created"] = databases.DateCreated

	var clonedFromInfo *model.TessellDatabaseClonedFromInfo
	if databases.ClonedFromInfo != clonedFromInfo {
		parsedDatabases["cloned_from_info"] = []interface{}{parseTessellDatabaseClonedFromInfo(databases.ClonedFromInfo)}
	}

	var databaseConfiguration *model.DatabaseConfiguration
	if databases.DatabaseConfiguration != databaseConfiguration {
		parsedDatabases["database_configuration"] = []interface{}{parseDatabaseConfiguration(databases.DatabaseConfiguration)}
	}

	return parsedDatabases
}

func parseTessellDatabaseClonedFromInfo(clonedFromInfo *model.TessellDatabaseClonedFromInfo) interface{} {
	if clonedFromInfo == nil {
		return nil
	}
	parsedClonedFromInfo := make(map[string]interface{})
	parsedClonedFromInfo["database_id"] = clonedFromInfo.DatabaseId

	return parsedClonedFromInfo
}

func parseDatabaseConfiguration(databaseConfiguration *model.DatabaseConfiguration) interface{} {
	if databaseConfiguration == nil {
		return nil
	}
	parsedDatabaseConfiguration := make(map[string]interface{})

	var oracleConfig *model.OracleDatabaseConfig
	if databaseConfiguration.OracleConfig != oracleConfig {
		parsedDatabaseConfiguration["oracle_config"] = []interface{}{parseOracleDatabaseConfig(databaseConfiguration.OracleConfig)}
	}

	var postgresqlConfig *model.PostgresqlDatabaseConfig
	if databaseConfiguration.PostgresqlConfig != postgresqlConfig {
		parsedDatabaseConfiguration["postgresql_config"] = []interface{}{parsePostgresqlDatabaseConfig(databaseConfiguration.PostgresqlConfig)}
	}

	var mySqlConfig *model.MySqlDatabaseConfig
	if databaseConfiguration.MySqlConfig != mySqlConfig {
		parsedDatabaseConfiguration["my_sql_config"] = []interface{}{parseMySqlDatabaseConfig(databaseConfiguration.MySqlConfig)}
	}

	var sqlServerConfig *model.SqlServerDatabaseConfig
	if databaseConfiguration.SqlServerConfig != sqlServerConfig {
		parsedDatabaseConfiguration["sql_server_config"] = []interface{}{parseSqlServerDatabaseConfig(databaseConfiguration.SqlServerConfig)}
	}

	return parsedDatabaseConfiguration
}

func parseOracleDatabaseConfig(oracleDatabaseConfig *model.OracleDatabaseConfig) interface{} {
	if oracleDatabaseConfig == nil {
		return nil
	}
	parsedOracleDatabaseConfig := make(map[string]interface{})
	parsedOracleDatabaseConfig["parameter_profile"] = oracleDatabaseConfig.ParameterProfile
	parsedOracleDatabaseConfig["options_profile"] = oracleDatabaseConfig.OptionsProfile

	return parsedOracleDatabaseConfig
}

func parsePostgresqlDatabaseConfig(postgresqlDatabaseConfig *model.PostgresqlDatabaseConfig) interface{} {
	if postgresqlDatabaseConfig == nil {
		return nil
	}
	parsedPostgresqlDatabaseConfig := make(map[string]interface{})
	parsedPostgresqlDatabaseConfig["parameter_profile"] = postgresqlDatabaseConfig.ParameterProfile

	return parsedPostgresqlDatabaseConfig
}

func parseMySqlDatabaseConfig(mySqlDatabaseConfig *model.MySqlDatabaseConfig) interface{} {
	if mySqlDatabaseConfig == nil {
		return nil
	}
	parsedMySqlDatabaseConfig := make(map[string]interface{})
	parsedMySqlDatabaseConfig["parameter_profile"] = mySqlDatabaseConfig.ParameterProfile

	return parsedMySqlDatabaseConfig
}

func parseSqlServerDatabaseConfig(sqlServerDatabaseConfig *model.SqlServerDatabaseConfig) interface{} {
	if sqlServerDatabaseConfig == nil {
		return nil
	}
	parsedSqlServerDatabaseConfig := make(map[string]interface{})
	parsedSqlServerDatabaseConfig["parameter_profile"] = sqlServerDatabaseConfig.ParameterProfile

	return parsedSqlServerDatabaseConfig
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

func formPayloadForCloneTessellService(d *schema.ResourceData) model.CloneTessellServicePayload {
	cloneTessellServicePayloadFormed := model.CloneTessellServicePayload{
		SnapshotId:               helper.GetStringPointer(d.Get("snapshot_id")),
		Pitr:                     helper.GetStringPointer(d.Get("pitr")),
		Name:                     helper.GetStringPointer(d.Get("name")),
		Description:              helper.GetStringPointer(d.Get("description")),
		Subscription:             helper.GetStringPointer(d.Get("subscription")),
		EngineType:               helper.GetStringPointer(d.Get("engine_type")),
		Topology:                 helper.GetStringPointer(d.Get("topology")),
		NumOfInstances:           helper.GetIntPointer(d.Get("num_of_instances")),
		SoftwareImage:            helper.GetStringPointer(d.Get("software_image")),
		SoftwareImageVersion:     helper.GetStringPointer(d.Get("software_image_version")),
		AutoMinorVersionUpdate:   helper.GetBoolPointer(d.Get("auto_minor_version_update")),
		EnableDeletionProtection: helper.GetBoolPointer(d.Get("enable_deletion_protection")),
		Infrastructure:           formTessellServiceInfrastructurePayload(d.Get("infrastructure")),
		ServiceConnectivity:      formTessellServiceConnectivityInfoPayload(d.Get("service_connectivity")),
		Creds:                    formTessellServiceCredsPayload(d.Get("creds")),
		MaintenanceWindow:        formTessellServiceMaintenanceWindow(d.Get("maintenance_window")),
		SnapshotConfiguration:    formTessellServiceBackupConfigurationPayload(d.Get("snapshot_configuration")),
		EngineConfiguration:      formTessellServiceEngineConfigurationPayload1(d.Get("engine_configuration")),
		Databases:                formCreateDatabasePayload1List(d.Get("databases")),
		IntegrationsConfig:       formTessellServiceIntegrationsPayload(d.Get("integrations_config")),
		Tags:                     formTessellTagList(d.Get("tags")),
	}

	return cloneTessellServicePayloadFormed
}

func formPayloadForDeleteTessellService(d *schema.ResourceData) model.DeleteTessellServicePayload {
	deleteTessellServicePayloadFormed := model.DeleteTessellServicePayload{
		DeletionConfig: formTessellServiceDeletionConfig(d.Get("deletion_config")),
	}

	return deleteTessellServicePayloadFormed
}

func formPayloadForProvisionTessellService(d *schema.ResourceData) model.ProvisionTessellServicePayload {
	provisionTessellServicePayloadFormed := model.ProvisionTessellServicePayload{
		Name:                     helper.GetStringPointer(d.Get("name")),
		Description:              helper.GetStringPointer(d.Get("description")),
		Subscription:             helper.GetStringPointer(d.Get("subscription")),
		EngineType:               helper.GetStringPointer(d.Get("engine_type")),
		Topology:                 helper.GetStringPointer(d.Get("topology")),
		NumOfInstances:           helper.GetIntPointer(d.Get("num_of_instances")),
		SoftwareImage:            helper.GetStringPointer(d.Get("software_image")),
		SoftwareImageVersion:     helper.GetStringPointer(d.Get("software_image_version")),
		AutoMinorVersionUpdate:   helper.GetBoolPointer(d.Get("auto_minor_version_update")),
		EnableDeletionProtection: helper.GetBoolPointer(d.Get("enable_deletion_protection")),
		Infrastructure:           formTessellServiceInfrastructurePayload(d.Get("infrastructure")),
		ServiceConnectivity:      formTessellServiceConnectivityInfoPayload(d.Get("service_connectivity")),
		Creds:                    formTessellServiceCredsPayload(d.Get("creds")),
		MaintenanceWindow:        formTessellServiceMaintenanceWindow(d.Get("maintenance_window")),
		SnapshotConfiguration:    formTessellServiceBackupConfigurationPayload(d.Get("snapshot_configuration")),
		EngineConfiguration:      formTessellServiceEngineConfigurationPayload(d.Get("engine_configuration")),
		Databases:                formCreateDatabasePayloadList(d.Get("databases")),
		IntegrationsConfig:       formTessellServiceIntegrationsPayload(d.Get("integrations_config")),
		Tags:                     formTessellTagList(d.Get("tags")),
	}

	return provisionTessellServicePayloadFormed
}

func formTessellServiceInfrastructurePayload(tessellServiceInfrastructurePayloadRaw interface{}) *model.TessellServiceInfrastructurePayload {
	if tessellServiceInfrastructurePayloadRaw == nil || len(tessellServiceInfrastructurePayloadRaw.([]interface{})) == 0 {
		return nil
	}

	tessellServiceInfrastructurePayloadData := tessellServiceInfrastructurePayloadRaw.([]interface{})[0].(map[string]interface{})

	tessellServiceInfrastructurePayloadFormed := model.TessellServiceInfrastructurePayload{
		Cloud:             helper.GetStringPointer(tessellServiceInfrastructurePayloadData["cloud"]),
		Region:            helper.GetStringPointer(tessellServiceInfrastructurePayloadData["region"]),
		AvailabilityZone:  helper.GetStringPointer(tessellServiceInfrastructurePayloadData["availability_zone"]),
		Vpc:               helper.GetStringPointer(tessellServiceInfrastructurePayloadData["vpc"]),
		ComputeType:       helper.GetStringPointer(tessellServiceInfrastructurePayloadData["compute_type"]),
		AdditionalStorage: helper.GetIntPointer(tessellServiceInfrastructurePayloadData["additional_storage"]),
	}

	return &tessellServiceInfrastructurePayloadFormed
}

func formTessellServiceConnectivityInfoPayload(tessellServiceConnectivityInfoPayloadRaw interface{}) *model.TessellServiceConnectivityInfoPayload {
	if tessellServiceConnectivityInfoPayloadRaw == nil || len(tessellServiceConnectivityInfoPayloadRaw.([]interface{})) == 0 {
		return nil
	}

	tessellServiceConnectivityInfoPayloadData := tessellServiceConnectivityInfoPayloadRaw.([]interface{})[0].(map[string]interface{})

	tessellServiceConnectivityInfoPayloadFormed := model.TessellServiceConnectivityInfoPayload{
		DnsPrefix:          helper.GetStringPointer(tessellServiceConnectivityInfoPayloadData["dns_prefix"]),
		ServicePort:        helper.GetIntPointer(tessellServiceConnectivityInfoPayloadData["service_port"]),
		EnablePublicAccess: helper.GetBoolPointer(tessellServiceConnectivityInfoPayloadData["enable_public_access"]),
		AllowedIpAddresses: helper.InterfaceToStringSlice(tessellServiceConnectivityInfoPayloadData["allowed_ip_addresses"]),
	}

	return &tessellServiceConnectivityInfoPayloadFormed
}

func formTessellServiceCredsPayload(tessellServiceCredsPayloadRaw interface{}) *model.TessellServiceCredsPayload {
	if tessellServiceCredsPayloadRaw == nil || len(tessellServiceCredsPayloadRaw.([]interface{})) == 0 {
		return nil
	}

	tessellServiceCredsPayloadData := tessellServiceCredsPayloadRaw.([]interface{})[0].(map[string]interface{})

	tessellServiceCredsPayloadFormed := model.TessellServiceCredsPayload{
		MasterUser:     helper.GetStringPointer(tessellServiceCredsPayloadData["master_user"]),
		MasterPassword: helper.GetStringPointer(tessellServiceCredsPayloadData["master_password"]),
	}

	return &tessellServiceCredsPayloadFormed
}

func formTessellServiceMaintenanceWindow(tessellServiceMaintenanceWindowRaw interface{}) *model.TessellServiceMaintenanceWindow {
	if tessellServiceMaintenanceWindowRaw == nil || len(tessellServiceMaintenanceWindowRaw.([]interface{})) == 0 {
		return nil
	}

	tessellServiceMaintenanceWindowData := tessellServiceMaintenanceWindowRaw.([]interface{})[0].(map[string]interface{})

	tessellServiceMaintenanceWindowFormed := model.TessellServiceMaintenanceWindow{
		Day:      helper.GetStringPointer(tessellServiceMaintenanceWindowData["day"]),
		Time:     helper.GetStringPointer(tessellServiceMaintenanceWindowData["time"]),
		Duration: helper.GetIntPointer(tessellServiceMaintenanceWindowData["duration"]),
	}

	return &tessellServiceMaintenanceWindowFormed
}

func formTessellServiceBackupConfigurationPayload(tessellServiceBackupConfigurationPayloadRaw interface{}) *model.TessellServiceBackupConfigurationPayload {
	if tessellServiceBackupConfigurationPayloadRaw == nil || len(tessellServiceBackupConfigurationPayloadRaw.([]interface{})) == 0 {
		return nil
	}

	tessellServiceBackupConfigurationPayloadData := tessellServiceBackupConfigurationPayloadRaw.([]interface{})[0].(map[string]interface{})

	tessellServiceBackupConfigurationPayloadFormed := model.TessellServiceBackupConfigurationPayload{
		AutoSnapshot:   helper.GetBoolPointer(tessellServiceBackupConfigurationPayloadData["auto_snapshot"]),
		Sla:            helper.GetStringPointer(tessellServiceBackupConfigurationPayloadData["sla"]),
		SnapshotWindow: formTessellServiceBackupConfigurationPayloadSnapshotWindow(tessellServiceBackupConfigurationPayloadData["snapshot_window"]),
	}

	return &tessellServiceBackupConfigurationPayloadFormed
}

func formTessellServiceBackupConfigurationPayloadSnapshotWindow(tessellServiceBackupConfigurationPayloadSnapshotWindowRaw interface{}) *model.TessellServiceBackupConfigurationPayloadSnapshotWindow {
	if tessellServiceBackupConfigurationPayloadSnapshotWindowRaw == nil || len(tessellServiceBackupConfigurationPayloadSnapshotWindowRaw.([]interface{})) == 0 {
		return nil
	}

	tessellServiceBackupConfigurationPayloadSnapshotWindowData := tessellServiceBackupConfigurationPayloadSnapshotWindowRaw.([]interface{})[0].(map[string]interface{})

	tessellServiceBackupConfigurationPayloadSnapshotWindowFormed := model.TessellServiceBackupConfigurationPayloadSnapshotWindow{
		Time:     helper.GetStringPointer(tessellServiceBackupConfigurationPayloadSnapshotWindowData["time"]),
		Duration: helper.GetIntPointer(tessellServiceBackupConfigurationPayloadSnapshotWindowData["duration"]),
	}

	return &tessellServiceBackupConfigurationPayloadSnapshotWindowFormed
}

func formTessellServiceEngineConfigurationPayload1(engineConfigurationRaw interface{}) *model.TessellServiceEngineConfigurationPayload1 {
	if engineConfigurationRaw == nil || len(engineConfigurationRaw.([]interface{})) == 0 {
		return nil
	}

	engineConfigurationData := engineConfigurationRaw.([]interface{})[0].(map[string]interface{})

	tessellServiceEngineConfigurationPayload1Formed := model.TessellServiceEngineConfigurationPayload1{
		PreScriptInfo:     formScriptInfo(engineConfigurationData["pre_script_info"]),
		PostScriptInfo:    formScriptInfo(engineConfigurationData["post_script_info"]),
		OracleConfig:      formOracleEngineConfigPayload(engineConfigurationData["oracle_config"]),
		PostgresqlConfig:  formPostgresqlEngineConfigPayload(engineConfigurationData["postgresql_config"]),
		MysqlConfig:       formMySqlEngineConfigPayload(engineConfigurationData["mysql_config"]),
		SqlServerConfig:   formSqlServerEngineConfigPayload(engineConfigurationData["sql_server_config"]),
		ApacheKafkaConfig: formApacheKafkaEngineConfigPayload(engineConfigurationData["apache_kafka_config"]),
	}

	return &tessellServiceEngineConfigurationPayload1Formed
}

func formScriptInfo(scriptInfoRaw interface{}) *model.ScriptInfo {
	if scriptInfoRaw == nil || len(scriptInfoRaw.([]interface{})) == 0 {
		return nil
	}

	scriptInfoData := scriptInfoRaw.([]interface{})[0].(map[string]interface{})

	scriptInfoFormed := model.ScriptInfo{
		ScriptId:      helper.GetStringPointer(scriptInfoData["script_id"]),
		ScriptVersion: helper.GetStringPointer(scriptInfoData["script_version"]),
	}

	return &scriptInfoFormed
}

func formOracleEngineConfigPayload(oracleEngineConfigPayloadRaw interface{}) *model.OracleEngineConfigPayload {
	if oracleEngineConfigPayloadRaw == nil || len(oracleEngineConfigPayloadRaw.([]interface{})) == 0 {
		return nil
	}

	oracleEngineConfigPayloadData := oracleEngineConfigPayloadRaw.([]interface{})[0].(map[string]interface{})

	oracleEngineConfigPayloadFormed := model.OracleEngineConfigPayload{
		MultiTenant:          helper.GetBoolPointer(oracleEngineConfigPayloadData["multi_tenant"]),
		ParameterProfile:     helper.GetStringPointer(oracleEngineConfigPayloadData["parameter_profile"]),
		OptionsProfile:       helper.GetStringPointer(oracleEngineConfigPayloadData["options_profile"]),
		CharacterSet:         helper.GetStringPointer(oracleEngineConfigPayloadData["character_set"]),
		NationalCharacterSet: helper.GetStringPointer(oracleEngineConfigPayloadData["national_character_set"]),
	}

	return &oracleEngineConfigPayloadFormed
}

func formPostgresqlEngineConfigPayload(postgresqlEngineConfigPayloadRaw interface{}) *model.PostgresqlEngineConfigPayload {
	if postgresqlEngineConfigPayloadRaw == nil || len(postgresqlEngineConfigPayloadRaw.([]interface{})) == 0 {
		return nil
	}

	postgresqlEngineConfigPayloadData := postgresqlEngineConfigPayloadRaw.([]interface{})[0].(map[string]interface{})

	postgresqlEngineConfigPayloadFormed := model.PostgresqlEngineConfigPayload{
		ParameterProfile: helper.GetStringPointer(postgresqlEngineConfigPayloadData["parameter_profile"]),
	}

	return &postgresqlEngineConfigPayloadFormed
}

func formMySqlEngineConfigPayload(mySqlEngineConfigPayloadRaw interface{}) *model.MySqlEngineConfigPayload {
	if mySqlEngineConfigPayloadRaw == nil || len(mySqlEngineConfigPayloadRaw.([]interface{})) == 0 {
		return nil
	}

	mySqlEngineConfigPayloadData := mySqlEngineConfigPayloadRaw.([]interface{})[0].(map[string]interface{})

	mySqlEngineConfigPayloadFormed := model.MySqlEngineConfigPayload{
		ParameterProfile: helper.GetStringPointer(mySqlEngineConfigPayloadData["parameter_profile"]),
	}

	return &mySqlEngineConfigPayloadFormed
}

func formSqlServerEngineConfigPayload(sqlServerEngineConfigPayloadRaw interface{}) *model.SqlServerEngineConfigPayload {
	if sqlServerEngineConfigPayloadRaw == nil || len(sqlServerEngineConfigPayloadRaw.([]interface{})) == 0 {
		return nil
	}

	sqlServerEngineConfigPayloadData := sqlServerEngineConfigPayloadRaw.([]interface{})[0].(map[string]interface{})

	sqlServerEngineConfigPayloadFormed := model.SqlServerEngineConfigPayload{
		ParameterProfile: helper.GetStringPointer(sqlServerEngineConfigPayloadData["parameter_profile"]),
	}

	return &sqlServerEngineConfigPayloadFormed
}

func formApacheKafkaEngineConfigPayload(apacheKafkaEngineConfigPayloadRaw interface{}) *model.ApacheKafkaEngineConfigPayload {
	if apacheKafkaEngineConfigPayloadRaw == nil || len(apacheKafkaEngineConfigPayloadRaw.([]interface{})) == 0 {
		return nil
	}

	apacheKafkaEngineConfigPayloadData := apacheKafkaEngineConfigPayloadRaw.([]interface{})[0].(map[string]interface{})

	apacheKafkaEngineConfigPayloadFormed := model.ApacheKafkaEngineConfigPayload{
		ParameterProfile: helper.GetStringPointer(apacheKafkaEngineConfigPayloadData["parameter_profile"]),
	}

	return &apacheKafkaEngineConfigPayloadFormed
}

func formCreateDatabasePayload1(databasesRaw interface{}) *model.CreateDatabasePayload1 {
	if databasesRaw == nil {
		return nil
	}

	databasesData := databasesRaw.(map[string]interface{})

	createDatabasePayload1Formed := model.CreateDatabasePayload1{
		DatabaseName:          helper.GetStringPointer(databasesData["database_name"]),
		SourceDatabaseId:      helper.GetStringPointer(databasesData["source_database_id"]),
		DatabaseConfiguration: formCreateDatabasePayloadDatabaseConfiguration(databasesData["database_configuration"]),
	}

	return &createDatabasePayload1Formed
}
func formCreateDatabasePayload1List(databasesListRaw interface{}) *[]model.CreateDatabasePayload1 {
	if databasesListRaw == nil || len(databasesListRaw.([]interface{})) == 0 {
		return nil
	}

	DatabasesListFormed := make([]model.CreateDatabasePayload1, len(databasesListRaw.([]interface{})))

	for i, databases := range databasesListRaw.([]interface{}) {
		DatabasesListFormed[i] = *formCreateDatabasePayload1(databases)
	}

	return &DatabasesListFormed
}
func formCreateDatabasePayloadDatabaseConfiguration(createDatabasePayloadDatabaseConfigurationRaw interface{}) *model.CreateDatabasePayloadDatabaseConfiguration {
	if createDatabasePayloadDatabaseConfigurationRaw == nil || len(createDatabasePayloadDatabaseConfigurationRaw.([]interface{})) == 0 {
		return nil
	}

	createDatabasePayloadDatabaseConfigurationData := createDatabasePayloadDatabaseConfigurationRaw.([]interface{})[0].(map[string]interface{})

	createDatabasePayloadDatabaseConfigurationFormed := model.CreateDatabasePayloadDatabaseConfiguration{
		OracleConfig:     formOracleDatabaseConfig(createDatabasePayloadDatabaseConfigurationData["oracle_config"]),
		PostgresqlConfig: formPostgresqlDatabaseConfig(createDatabasePayloadDatabaseConfigurationData["postgresql_config"]),
		MysqlConfig:      formMySqlDatabaseConfig(createDatabasePayloadDatabaseConfigurationData["mysql_config"]),
		SqlServerConfig:  formSqlServerDatabaseConfig(createDatabasePayloadDatabaseConfigurationData["sql_server_config"]),
	}

	return &createDatabasePayloadDatabaseConfigurationFormed
}

func formOracleDatabaseConfig(oracleDatabaseConfigRaw interface{}) *model.OracleDatabaseConfig {
	if oracleDatabaseConfigRaw == nil || len(oracleDatabaseConfigRaw.([]interface{})) == 0 {
		return nil
	}

	oracleDatabaseConfigData := oracleDatabaseConfigRaw.([]interface{})[0].(map[string]interface{})

	oracleDatabaseConfigFormed := model.OracleDatabaseConfig{
		ParameterProfile: helper.GetStringPointer(oracleDatabaseConfigData["parameter_profile"]),
		OptionsProfile:   helper.GetStringPointer(oracleDatabaseConfigData["options_profile"]),
	}

	return &oracleDatabaseConfigFormed
}

func formPostgresqlDatabaseConfig(postgresqlDatabaseConfigRaw interface{}) *model.PostgresqlDatabaseConfig {
	if postgresqlDatabaseConfigRaw == nil || len(postgresqlDatabaseConfigRaw.([]interface{})) == 0 {
		return nil
	}

	postgresqlDatabaseConfigData := postgresqlDatabaseConfigRaw.([]interface{})[0].(map[string]interface{})

	postgresqlDatabaseConfigFormed := model.PostgresqlDatabaseConfig{
		ParameterProfile: helper.GetStringPointer(postgresqlDatabaseConfigData["parameter_profile"]),
	}

	return &postgresqlDatabaseConfigFormed
}

func formMySqlDatabaseConfig(mySqlDatabaseConfigRaw interface{}) *model.MySqlDatabaseConfig {
	if mySqlDatabaseConfigRaw == nil || len(mySqlDatabaseConfigRaw.([]interface{})) == 0 {
		return nil
	}

	mySqlDatabaseConfigData := mySqlDatabaseConfigRaw.([]interface{})[0].(map[string]interface{})

	mySqlDatabaseConfigFormed := model.MySqlDatabaseConfig{
		ParameterProfile: helper.GetStringPointer(mySqlDatabaseConfigData["parameter_profile"]),
	}

	return &mySqlDatabaseConfigFormed
}

func formSqlServerDatabaseConfig(sqlServerDatabaseConfigRaw interface{}) *model.SqlServerDatabaseConfig {
	if sqlServerDatabaseConfigRaw == nil || len(sqlServerDatabaseConfigRaw.([]interface{})) == 0 {
		return nil
	}

	sqlServerDatabaseConfigData := sqlServerDatabaseConfigRaw.([]interface{})[0].(map[string]interface{})

	sqlServerDatabaseConfigFormed := model.SqlServerDatabaseConfig{
		ParameterProfile: helper.GetStringPointer(sqlServerDatabaseConfigData["parameter_profile"]),
	}

	return &sqlServerDatabaseConfigFormed
}

func formTessellServiceIntegrationsPayload(tessellServiceIntegrationsPayloadRaw interface{}) *model.TessellServiceIntegrationsPayload {
	if tessellServiceIntegrationsPayloadRaw == nil || len(tessellServiceIntegrationsPayloadRaw.([]interface{})) == 0 {
		return nil
	}

	tessellServiceIntegrationsPayloadData := tessellServiceIntegrationsPayloadRaw.([]interface{})[0].(map[string]interface{})

	tessellServiceIntegrationsPayloadFormed := model.TessellServiceIntegrationsPayload{
		Integrations: helper.InterfaceToStringSlice(tessellServiceIntegrationsPayloadData["integrations"]),
	}

	return &tessellServiceIntegrationsPayloadFormed
}

func formTessellTag(tessellTagRaw interface{}) *model.TessellTag {
	if tessellTagRaw == nil {
		return nil
	}

	tessellTagData := tessellTagRaw.(map[string]interface{})

	tessellTagFormed := model.TessellTag{
		Name:  helper.GetStringPointer(tessellTagData["name"]),
		Value: helper.GetStringPointer(tessellTagData["value"]),
	}

	return &tessellTagFormed
}
func formTessellTagList(tessellTagListRaw interface{}) *[]model.TessellTag {
	if tessellTagListRaw == nil || len(tessellTagListRaw.([]interface{})) == 0 {
		return nil
	}

	TessellTagListFormed := make([]model.TessellTag, len(tessellTagListRaw.([]interface{})))

	for i, tessellTag := range tessellTagListRaw.([]interface{}) {
		TessellTagListFormed[i] = *formTessellTag(tessellTag)
	}

	return &TessellTagListFormed
}
func formTessellServiceDeletionConfig(deletionConfigRaw interface{}) *model.TessellServiceDeletionConfig {
	if deletionConfigRaw == nil || len(deletionConfigRaw.([]interface{})) == 0 {
		return nil
	}

	deletionConfigData := deletionConfigRaw.([]interface{})[0].(map[string]interface{})

	tessellServiceDeletionConfigFormed := model.TessellServiceDeletionConfig{
		RetainAvailabilityMachine: helper.GetBoolPointer(deletionConfigData["retain_availability_machine"]),
	}

	return &tessellServiceDeletionConfigFormed
}

func formTessellServiceEngineConfigurationPayload(engineConfigurationRaw interface{}) *model.TessellServiceEngineConfigurationPayload {
	if engineConfigurationRaw == nil || len(engineConfigurationRaw.([]interface{})) == 0 {
		return nil
	}

	engineConfigurationData := engineConfigurationRaw.([]interface{})[0].(map[string]interface{})

	tessellServiceEngineConfigurationPayloadFormed := model.TessellServiceEngineConfigurationPayload{
		PreScriptInfo:     formScriptInfo(engineConfigurationData["pre_script_info"]),
		PostScriptInfo:    formScriptInfo(engineConfigurationData["post_script_info"]),
		OracleConfig:      formOracleEngineConfigPayload(engineConfigurationData["oracle_config"]),
		PostgresqlConfig:  formPostgresqlEngineConfigPayload(engineConfigurationData["postgresql_config"]),
		MysqlConfig:       formMySqlEngineConfigPayload(engineConfigurationData["mysql_config"]),
		SqlServerConfig:   formSqlServerEngineConfigPayload(engineConfigurationData["sql_server_config"]),
		ApacheKafkaConfig: formApacheKafkaEngineConfigPayload(engineConfigurationData["apache_kafka_config"]),
	}

	return &tessellServiceEngineConfigurationPayloadFormed
}

func formCreateDatabasePayload(databasesRaw interface{}) *model.CreateDatabasePayload {
	if databasesRaw == nil {
		return nil
	}

	databasesData := databasesRaw.(map[string]interface{})

	createDatabasePayloadFormed := model.CreateDatabasePayload{
		DatabaseName:          helper.GetStringPointer(databasesData["database_name"]),
		SourceDatabaseId:      helper.GetStringPointer(databasesData["source_database_id"]),
		DatabaseConfiguration: formCreateDatabasePayloadDatabaseConfiguration(databasesData["database_configuration"]),
	}

	return &createDatabasePayloadFormed
}
func formCreateDatabasePayloadList(databasesListRaw interface{}) *[]model.CreateDatabasePayload {
	if databasesListRaw == nil || len(databasesListRaw.([]interface{})) == 0 {
		return nil
	}

	DatabasesListFormed := make([]model.CreateDatabasePayload, len(databasesListRaw.([]interface{})))

	for i, databases := range databasesListRaw.([]interface{}) {
		DatabasesListFormed[i] = *formCreateDatabasePayload(databases)
	}

	return &DatabasesListFormed
}
