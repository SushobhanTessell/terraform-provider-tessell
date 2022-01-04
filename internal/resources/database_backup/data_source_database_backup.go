package database_backup

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	apiClient "terraform-provider-tessell/internal/client"
)

func DataSourceDatabaseBackup() *schema.Resource {
	return &schema.Resource{
		Description: "Sample data source in the Terraform provider Database.",

		ReadContext: dataSourceDatabaseBackupRead,

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
			"task_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"backup_name": {
				Type:        schema.TypeString,
				Description: "Database backup name",
				Required:    true,
			},
			"dmm_name": {
				Type:        schema.TypeString,
				Description: "Database backup DMM name",
				Required:    true,
			},
		},
	}
}

func dataSourceDatabaseBackupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	backupName := d.Get("backup_name").(string)
	dmmName := d.Get("dmm_name").(string)

	databaseBackup, err := client.GetDatabaseBackupByName(backupName, dmmName)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := setResourceData(d, databaseBackup); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("databaseBackup-%s-%s", backupName, dmmName))

	return diags
}
