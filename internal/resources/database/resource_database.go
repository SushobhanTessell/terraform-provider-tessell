package database

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	apiClient "terraform-provider-tessell/internal/client"
	models "terraform-provider-tessell/internal/models"
	"terraform-provider-tessell/internal/utils"
)

func ResourceDatabase() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceDatabaseCreate,
		ReadContext:   resourceDatabaseRead,
		UpdateContext: resourceDatabaseUpdate,
		DeleteContext: resourceDatabaseDelete,

		Schema: map[string]*schema.Schema{
			"backup_configuration": {
				Type:          schema.TypeList,
				MaxItems:      1,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"clone_database_name"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_backup": {
							Type:     schema.TypeBool,
							Required: true,
							ForceNew: true,
						},
						"backup_sla": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"backup_window": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Required: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"duration": {
										Type:     schema.TypeInt,
										Required: true,
										ForceNew: true,
									},
									"time": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
								},
							},
						},
					},
				},
			},
			"config_type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"database_configuration": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_minor_version_update": {
							Type:     schema.TypeBool,
							Optional: true,
							ForceNew: true,
						},
						"character_set": {
							Type:          schema.TypeString,
							Optional:      true,
							ForceNew:      true,
							ConflictsWith: []string{"clone_database_name"},
						},
						"create_replica": {
							Type:          schema.TypeBool,
							Optional:      true,
							ForceNew:      true,
							ConflictsWith: []string{"clone_database_name"},
						},
						"database_compute_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"database_network_profile_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"database_options_profile_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"database_parameter_profile_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"database_software_image_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"enable_deletion_protection": {
							Type:     schema.TypeBool,
							Optional: true,
							ForceNew: true,
						},
						"license_type": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"maintenance_window": {
							Type:          schema.TypeList,
							MaxItems:      1,
							Optional:      true,
							ForceNew:      true,
							ConflictsWith: []string{"clone_database_name"},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"day": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"duration": {
										Type:     schema.TypeInt,
										Required: true,
										ForceNew: true,
									},
									"time": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
								},
							},
						},
						"multitenant": {
							Type:          schema.TypeBool,
							Optional:      true,
							ForceNew:      true,
							ConflictsWith: []string{"clone_database_name"},
						},
						"national_character_set": {
							Type:          schema.TypeString,
							Optional:      true,
							ForceNew:      true,
							ConflictsWith: []string{"clone_database_name"},
						},
						"options_profile": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"parameter_profile": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"software_image_version": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"clone_metadata_info": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"created_from_backup_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"created_from_backup_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"created_from_pitr": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"rpo_based": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"manual_backup": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"created_from_dap_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"created_from_dap_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"database_name": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ExactlyOneOf: []string{"database_name", "clone_database_name"},
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringMatch(
						regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9]{2,7}$`),
						"Name should match regexp `^[a-zA-Z][a-zA-Z0-9]{2,7}$`",
					),
				),
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dmm_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"engine_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"infrastructure": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allowed_ip_addresses": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"availability_zone": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"cloud": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"compute_type": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"database_port": {
							Type:     schema.TypeInt,
							Required: true,
							ForceNew: true,
						},
						"enable_public_access": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"region": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"vpc_name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"cloud_availability": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"aws": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"azure": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"connection_info": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"data": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"connect_desc": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"endpoint": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"database_port": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"allowed_ip_addresses": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"public_access": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"deployment_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"license_token": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"compute": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"read_iops": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"write_iops": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"memory_gb": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"no_of_disks": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"storage_gb": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"vcpus": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"master_password": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"master_user": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"post_script": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"pre_script": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"service_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subscription": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeMap,
				},
			},
			"backup": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"database_name", "pitr"},
				RequiredWith:  []string{"clone_database_name"},
			},
			"pitr": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"database_name", "backup"},
				RequiredWith:  []string{"clone_database_name"},
			},
			"clone_database_name": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ExactlyOneOf: []string{"database_name", "clone_database_name"},
				ConflictsWith: []string{"database_configuration.0.character_set",
					"database_configuration.0.create_replica", "database_configuration.0.maintenance_window",
					"database_configuration.0.multitenant", "database_configuration.0.national_character_set",
				},
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringMatch(
						regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9]{2,7}$`),
						"Name should match regexp `^[a-zA-Z][a-zA-Z0-9]{2,7}$`",
					),
				),
			},
			"clone_dmm_name": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				RequiredWith: []string{"clone_database_name"},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"availability_machine": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"clone": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"topology": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"info": {
				Type:     schema.TypeMap,
				Computed: true,
			},
			"date_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"date_modified": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"user": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"on_delete": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"daps_to_retain": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeMap,
							},
						},
						"retain_availability_machine": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"take_final_backup": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
			"task_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"block_until_complete": {
				Type:        schema.TypeBool,
				Description: "Block the flow until the VPC has been successfully completed",
				Optional:    true,
				Default:     true,
			},
			"timeout": {
				Type:        schema.TypeInt,
				Description: "If block_until_complete is true, how long should it block for. (In seconds)",
				Optional:    true,
				Default:     1200,
			},
		},

		CustomizeDiff: func(c context.Context, d *schema.ResourceDiff, i interface{}) error {
			cloudType := d.Get("infrastructure.0.cloud_type").(string)
			az := d.Get("infrastructure.0.availability_zone").(string)
			errMsgPrefix := "when 'infrastructure.cloud_type' == 'AWS', 'infrastructure.availability_zone'"

			if cloudType == "AWS" {
				if !strings.HasPrefix(az, d.Get("infrastructure.0.region").(string)) {
					return fmt.Errorf("%s is not valid for the provided 'infrastructure.region'", errMsgPrefix)
				}
				matched, _ := regexp.MatchString(`[a-z]{2}-[a-z]*-[0-9][a-z]`, az)
				if !matched {
					return fmt.Errorf("%s must match regexp `[a-z]{2}-[a-z]*-[0-9][a-z]`", errMsgPrefix)
				}
			} else {
				if az != "" {
					return fmt.Errorf(
						"when 'infrastructure.cloud_type' == '%s', 'infrastructure.availability_zone' is not expected",
						cloudType,
					)
				}
			}
			return nil
		},
	}
}

func resourceDatabaseCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	if err := validate(d); err != nil {
		return diag.FromErr(err)
	}

	task := "Database Creation"
	var databaseCreationResponse *models.DatabaseCreationResponse
	if d.Get("database_name").(string) != "" {
		payload, err := formDatabaseCreationPayload(d)
		if err != nil {
			return diag.FromErr(err)
		}
		databaseCreationResponse, err = client.CreateDatabase(payload)
		if err != nil {
			return diag.FromErr(err)
		}
	} else if d.Get("clone_database_name").(string) != "" {
		payload, err := formDatabaseCloningPayload(d)
		if err != nil {
			return diag.FromErr(err)
		}
		databaseCreationResponse, err = client.CloneDatabase(d.Get("clone_dmm_name").(string), payload)
		if err != nil {
			return diag.FromErr(err)
		}
		task = "Database Cloning"
	}

	taskId := databaseCreationResponse.TaskId
	d.SetId(taskId)
	d.Set("task_id", taskId)

	if d.Get("block_until_complete").(bool) {
		if err := client.WaitTillWorkflowCompleted(taskId, d.Get("timeout").(int), task); err != nil {
			return diag.FromErr(err)
		}
	}

	resourceDatabaseRead(ctx, d, meta)

	return diags
}

func resourceDatabaseRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	database, err := client.GetDatabaseByName(d.Get("service_name").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	if err := setResourceData(d, database); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceDatabaseUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	if !d.HasChanges("description", "tags", "infrastructure.0.enable_public_access", "infrastructure.0.allowed_ip_addresses") {
		return diags
	}

	payload := models.DatabaseUpdationPayload{
		AllowedIpAddresses: utils.InteraceSliceToStringSlice(d.Get("infrastructure.0.allowed_ip_addresses").([]interface{})),
		EnablePublicAccess: d.Get("infrastructure.0.enable_public_access").(bool),
		Description:        d.Get("description").(string),
		Tags:               d.Get("tags").([]interface{}),
	}

	databaseUpdationResponse, err := client.UpdateDatabase(d.Get("name").(string), payload)
	if err != nil {
		return diag.FromErr(err)
	}

	if d.Get("block_until_complete").(bool) {
		err := client.WaitTillWorkflowCompleted(databaseUpdationResponse.TaskId, d.Get("timeout").(int), "Database Updation")
		if err != nil {
			return diag.FromErr(err)
		}
	}

	resourceDatabaseRead(ctx, d, meta)

	return diags
}

func resourceDatabaseDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	payload := models.DatabaseDeletionPayload{
		DapsToRetain:              utils.InteraceSliceToStringSlice(d.Get("on_delete.0.daps_to_retain").([]interface{})),
		RetainAvailabilityMachine: d.Get("on_delete.0.retain_availability_machine").(bool),
		TakeFinalBackup:           d.Get("on_delete.0.take_final_backup").(bool),
	}

	databaseDeletionResponse, err := client.DeleteDatabase(d.Get("name").(string), payload)
	if err != nil {
		return diag.FromErr(err)
	}

	err = client.WaitTillWorkflowCompleted(databaseDeletionResponse.TaskId, d.Get("timeout").(int), "Database Deletion")
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
