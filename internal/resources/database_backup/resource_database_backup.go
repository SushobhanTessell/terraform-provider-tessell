package database_backup

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	apiClient "terraform-provider-tessell/internal/client"
	models "terraform-provider-tessell/internal/models"
)

func ResourceDatabaseBackup() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceDatabaseBackupCreate,
		ReadContext:   resourceDatabaseBackupRead,
		DeleteContext: resourceDatabaseBackupDelete,

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
				ForceNew:    true,
			},
			"dmm_name": {
				Type:        schema.TypeString,
				Description: "Database backup DMM name",
				Required:    true,
				ForceNew:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Description: "Additional description for the backup",
				Optional:    true,
				Default:     "",
				ForceNew:    true,
			},
		},
	}
}

func resourceDatabaseBackupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	backupName := d.Get("backup_name").(string)
	dmmName := d.Get("dmm_name").(string)

	payload := models.DatabaseBackupCreationPayload{
		BackupName:  backupName,
		Description: d.Get("description").(string),
	}

	databaseBackupCreationResponse, err := client.CreateDatabaseBackup(dmmName, payload)
	if err != nil {
		return diag.FromErr(err)
	}

	taskId := databaseBackupCreationResponse.TaskId
	d.SetId(taskId)
	d.Set("task_id", taskId)

	if err := client.WaitTillWorkflowCompleted(taskId, 1200, "Database Backup Creation"); err != nil {
		return diag.FromErr(err)
	}

	resourceDatabaseBackupRead(ctx, d, meta)

	return diags
}

func resourceDatabaseBackupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	databaseBackup, err := client.GetDatabaseBackupByName(d.Get("backup_name").(string), d.Get("dmm_name").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	if err := setResourceData(d, databaseBackup); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceDatabaseBackupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	backupName := d.Get("backup_name").(string)
	dmmName := d.Get("dmm_name").(string)

	databaseBackupDeletionResponse, err := client.DeleteDatabaseBackup(backupName, dmmName)
	if err != nil {
		return diag.FromErr(err)
	}

	err = client.WaitTillWorkflowCompleted(databaseBackupDeletionResponse.TaskId, 600, "Database Backup Deletion")
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
