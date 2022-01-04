package database_backup

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	apiClient "terraform-provider-tessell/internal/client"
	"terraform-provider-tessell/internal/models"
)

func DataSourceDatabaseBackups() *schema.Resource {
	return &schema.Resource{
		Description: "Sample data source in the Terraform provider Database.",

		ReadContext: dataSourceDatabaseBackupsRead,

		Schema: map[string]*schema.Schema{
			"dmm_name": {
				Type:        schema.TypeString,
				Description: "Database backup DMM name",
				Required:    true,
				ForceNew:    true,
			},
			"database_backups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backup_time": {
							Type:        schema.TypeString,
							Description: "Database Backup capture time",
							Computed:    true,
						},
						"backup_type": {
							Type:        schema.TypeString,
							Description: "Database Backup's type",
							Computed:    true,
						},
						"cloud_location": {
							Type:        schema.TypeString,
							Description: "Database Backup's location in the cloud",
							Computed:    true,
						},
						"manual": {
							Type:        schema.TypeBool,
							Description: "Specifies if Database Backup's is manually created on a user request",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Database Snapshot name",
							Computed:    true,
						},
						"retention_type": {
							Type:        schema.TypeString,
							Description: "Database Backup's retention type",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Database Backup status",
							Computed:    true,
						},
						"size": {
							Type:        schema.TypeInt,
							Description: "Database Backup size in bytes",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceDatabaseBackupsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	dmmName := d.Get("dmm_name").(string)

	databaseBackups, err := client.GetDatabaseBackups(dmmName)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := setDataSourceValues(d, &databaseBackups.DatabaseBackups); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("databaseBackups-" + dmmName)

	return diags
}

func setDataSourceValues(d *schema.ResourceData, backups *[]models.DatabaseBackup) error {
	backupList := make([]interface{}, 0)

	if backups != nil {
		backupList = make([]interface{}, len(*backups))
		for i, backup := range *backups {
			backupList[i] = map[string]interface{}{
				"backup_time":    backup.BackupTime,
				"backup_type":    backup.BackupType,
				"cloud_location": backup.CloudLocation,
				"manual":         backup.Manual,
				"name":           backup.Name,
				"retention_type": backup.RetentionType,
				"size":           backup.Size,
				"status":         backup.Status,
			}
		}
	}

	if err := d.Set("database_backups", backupList); err != nil {
		return err
	}
	return nil
}
