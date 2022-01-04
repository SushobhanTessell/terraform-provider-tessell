package database_backup

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-tessell/internal/models"
)

func setResourceData(d *schema.ResourceData, databaseBackup *models.DatabaseBackup) error {
	if err := d.Set("backup_time", databaseBackup.BackupTime); err != nil {
		return err
	}

	if err := d.Set("backup_type", databaseBackup.BackupType); err != nil {
		return err
	}

	if err := d.Set("cloud_location", databaseBackup.CloudLocation); err != nil {
		return err
	}

	if err := d.Set("manual", databaseBackup.Manual); err != nil {
		return err
	}

	if err := d.Set("name", databaseBackup.Name); err != nil {
		return err
	}

	if err := d.Set("retention_type", databaseBackup.RetentionType); err != nil {
		return err
	}

	if err := d.Set("status", databaseBackup.Status); err != nil {
		return err
	}

	if err := d.Set("size", databaseBackup.Size); err != nil {
		return err
	}

	return nil
}
