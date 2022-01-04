package database

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-tessell/internal/models"
	"terraform-provider-tessell/internal/utils"
)

func validate(d *schema.ResourceData) error {
	return nil
}

func formDatabaseCreationPayload(d *schema.ResourceData) (models.DatabaseCreationPayload, error) {
	payload := models.DatabaseCreationPayload{
		ConfigType:     d.Get("config_type").(string),
		DatabaseName:   d.Get("database_name").(string),
		Description:    d.Get("description").(string),
		DmmName:        d.Get("dmm_name").(string),
		EngineType:     d.Get("engine_type").(string),
		MasterPassword: d.Get("master_password").(string),
		MasterUser:     d.Get("master_user").(string),
		PostScript:     d.Get("post_script").(string),
		PreScript:      d.Get("pre_script").(string),
		ServiceName:    d.Get("service_name").(string),
		Subscription:   d.Get("subscription").(string),
	}

	payload.Infrastructure = formPayloadInfrastructure(d)

	database_configuration := ((d.Get("database_configuration").([]interface{}))[0]).(map[string]interface{})
	maintenance_window := ((database_configuration["maintenance_window"].([]interface{}))[0]).(map[string]interface{})
	payload.DatabaseConfiguration = models.DatabaseCreationDatabaseConfiguration{
		SoftwareImageVersion:     database_configuration["software_image_version"].(string),
		CreateReplica:            database_configuration["create_replica"].(bool),
		LicenseType:              database_configuration["license_type"].(string),
		Multitenant:              database_configuration["multitenant"].(bool),
		CharacterSet:             database_configuration["character_set"].(string),
		NationalCharacterSet:     database_configuration["national_character_set"].(string),
		ParameterProfile:         database_configuration["parameter_profile"].(string),
		OptionsProfile:           database_configuration["options_profile"].(string),
		AutoMinorVersionUpdate:   database_configuration["auto_minor_version_update"].(bool),
		EnableDeletionProtection: database_configuration["enable_deletion_protection"].(bool),
		MaintenanceWindow: models.DatabaseCreationMaintenanceWindow{
			Day:      maintenance_window["day"].(string),
			Time:     maintenance_window["time"].(string),
			Duration: maintenance_window["duration"].(int),
		},
	}

	backup_configuration := ((d.Get("backup_configuration").([]interface{}))[0]).(map[string]interface{})
	backup_window := ((backup_configuration["backup_window"].([]interface{}))[0]).(map[string]interface{})
	payload.BackupConfiguration = models.DatabaseCreationBackupConfiguration{
		AutoBackup: backup_configuration["auto_backup"].(bool),
		BackupSLA:  backup_configuration["backup_sla"].(string),
		BackupWindow: models.DatabaseCreationBackupWindow{
			Time:     backup_window["time"].(string),
			Duration: backup_window["duration"].(int),
		},
	}

	if v, ok := d.GetOk("tags"); ok {
		payload.Tags = v.([]interface{})
	}

	return payload, nil
}

func formDatabaseCloningPayload(d *schema.ResourceData) (models.DatabaseCloningPayload, error) {
	payload := models.DatabaseCloningPayload{
		Backup:            d.Get("backup").(string),
		CloneDatabaseName: d.Get("clone_database_name").(string),
		ConfigType:        d.Get("config_type").(string),
		Description:       d.Get("description").(string),
		DmmName:           d.Get("dmm_name").(string),
		EngineType:        d.Get("engine_type").(string),
		MasterPassword:    d.Get("master_password").(string),
		MasterUser:        d.Get("master_user").(string),
		Pitr:              d.Get("pitr").(string),
		PostScript:        d.Get("post_script").(string),
		PreScript:         d.Get("pre_script").(string),
		ServiceName:       d.Get("service_name").(string),
		Subscription:      d.Get("subscription").(string),
	}

	payload.Infrastructure = formPayloadInfrastructure(d)

	database_configuration := ((d.Get("database_configuration").([]interface{}))[0]).(map[string]interface{})
	payload.DatabaseConfiguration = models.DatabaseCloningDatabaseConfiguration{
		SoftwareImageVersion:     database_configuration["software_image_version"].(string),
		LicenseType:              database_configuration["license_type"].(string),
		ParameterProfile:         database_configuration["parameter_profile"].(string),
		OptionsProfile:           database_configuration["options_profile"].(string),
		AutoMinorVersionUpdate:   database_configuration["auto_minor_version_update"].(bool),
		EnableDeletionProtection: database_configuration["enable_deletion_protection"].(bool),
	}

	if v, ok := d.GetOk("tags"); ok {
		payload.Tags = v.([]interface{})
	}

	return payload, nil
}

func formPayloadInfrastructure(d *schema.ResourceData) models.DatabaseCreationInfrastructure {
	infrastructure := ((d.Get("infrastructure").([]interface{}))[0]).(map[string]interface{})
	payloadInfrastructure := models.DatabaseCreationInfrastructure{
		AvailabilityZone:   infrastructure["availability_zone"].(string),
		Cloud:              infrastructure["cloud"].(string),
		ComputeType:        infrastructure["compute_type"].(string),
		DatabasePort:       infrastructure["database_port"].(int),
		EnablePublicAccess: infrastructure["enable_public_access"].(bool),
		Region:             infrastructure["region"].(string),
		VPCName:            infrastructure["vpc_name"].(string),
	}

	if allowedIpAddresses := infrastructure["allowed_ip_addresses"].([]interface{}); len(allowedIpAddresses) > 0 {
		payloadInfrastructure.AllowedIpAddresses = utils.InteraceSliceToStringSlice(allowedIpAddresses)
	}

	return payloadInfrastructure
}

func setResourceData(d *schema.ResourceData, database *models.Database) error {
	if err := d.Set("name", database.Name); err != nil {
		return err
	}

	if err := d.Set("subscription", database.Subscription); err != nil {
		return err
	}

	if err := d.Set("description", database.Description); err != nil {
		return err
	}

	if err := d.Set("availability_machine", database.AvailabilityMachine); err != nil {
		return err
	}

	if err := d.Set("engine_type", database.EngineType); err != nil {
		return err
	}

	if err := d.Set("status", database.Status); err != nil {
		return err
	}

	if err := d.Set("clone", database.Clone); err != nil {
		return err
	}

	if err := d.Set("config_type", database.ConfigType); err != nil {
		return err
	}

	if err := d.Set("topology", database.Topology); err != nil {
		return err
	}

	if err := d.Set("date_created", database.DateCreated); err != nil {
		return err
	}

	if err := d.Set("date_modified", database.DateModified); err != nil {
		return err
	}

	if err := d.Set("user", database.User); err != nil {
		return err
	}

	if err := d.Set("info", database.Info); err != nil {
		return err
	}

	if err := d.Set("database_configuration", parseDatabaseConfiguration(database, d)); err != nil {
		return err
	}

	if err := d.Set("infrastructure", parseDatabaseInfrastructure(&database.Infrastructure, d)); err != nil {
		return err
	}

	return nil
}

func parseDatabaseConfiguration(database *models.Database, d *schema.ResourceData) []interface{} {
	config := database.DatabaseConfiguration

	databaseConfiguration := make(map[string]interface{})
	if d.Get("database_configuration") != nil {
		databaseConfigurationResourceData := d.Get("database_configuration").([]interface{})
		if len(databaseConfigurationResourceData) > 0 {
			databaseConfiguration = (databaseConfigurationResourceData[0]).(map[string]interface{})
		}
	}

	databaseConfiguration["database_compute_id"] = config.DatabaseComputeId
	databaseConfiguration["database_network_profile_id"] = config.DatabaseNetworkProfileId
	databaseConfiguration["database_options_profile_id"] = config.DatabaseOptionsProfileId
	databaseConfiguration["database_parameter_profile_id"] = config.DatabaseParameterProfileId
	databaseConfiguration["database_software_image_id"] = config.DatabaseSoftwareImageId

	if database.Clone {
		cloneMetadataInfo := config.CloneMetadataInfo
		databaseConfiguration["clone_metadata_info"] = []interface{}{map[string]interface{}{
			"created_from_backup_id":   cloneMetadataInfo.CreatedFromBackupId,
			"created_from_backup_name": cloneMetadataInfo.CreatedFromBackupName,
			"created_from_dap_id":      cloneMetadataInfo.CreatedFromDapId,
			"created_from_dap_name":    cloneMetadataInfo.CreatedFromDapName,
			"created_from_pitr":        cloneMetadataInfo.CreatedFromPitr,
			"manual_backup":            cloneMetadataInfo.ManualBackup,
			"rpo_based":                cloneMetadataInfo.RpoBased,
		}}
	}

	return []interface{}{databaseConfiguration}
}

func parseDatabaseInfrastructure(infrastructure *models.DatabaseInfrastructure, d *schema.ResourceData) []interface{} {
	databaseInfrastructure := make(map[string]interface{})
	if d.Get("infrastructure") != nil {
		databaseInfrastructureResourceData := d.Get("infrastructure").([]interface{})
		if len(databaseInfrastructureResourceData) > 0 {
			databaseInfrastructure = (databaseInfrastructureResourceData[0]).(map[string]interface{})
		}
	}

	data := infrastructure.ConnectionInfo.Data
	databaseInfrastructure["connection_info"] = []interface{}{map[string]interface{}{
		"data": []interface{}{map[string]interface{}{
			"allowed_ip_addresses": data.AllowedIpAddresses,
			"connect_desc":         data.ConnectDesc,
			"database_port":        data.DatabasePort,
			"deployment_id":        data.DeploymentId,
			"endpoint":             data.Endpoint,
			"license_token":        data.LicenseToken,
			"public_access":        data.PublicAccess,
		}},
	}}

	cloudAvailability := infrastructure.CloudAvailability
	databaseInfrastructure["cloud_availability"] = []interface{}{map[string]interface{}{
		"aws":   cloudAvailability.AWS,
		"azure": cloudAvailability.Azure,
	}}

	compute := infrastructure.Compute
	databaseInfrastructure["compute"] = []interface{}{map[string]interface{}{
		"memory_gb":   compute.MemoryGB,
		"name":        compute.Name,
		"no_of_disks": compute.NoOfDisks,
		"read_iops":   compute.ReadIops,
		"storage_gb":  compute.StorageGB,
		"vcpus":       compute.Vcpus,
		"write_iops":  compute.WriteIops,
	}}

	return []interface{}{databaseInfrastructure}
}
