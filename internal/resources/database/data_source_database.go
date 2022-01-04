package database

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	apiClient "terraform-provider-tessell/internal/client"
)

func DataSourceDatabase() *schema.Resource {
	return &schema.Resource{

		ReadContext: dataSourceDatabaseRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subscription": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"availability_machine": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"engine_type": {
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
			"config_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"topology": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"database_configuration": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_minor_version_update": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"character_set": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"create_replica": {
							Type:     schema.TypeBool,
							Computed: true,
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
							Computed: true,
						},
						"license_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"maintenance_window": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"day": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"duration": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"time": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"multitenant": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"national_character_set": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"options_profile": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"parameter_profile": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"software_image_version": {
							Type:     schema.TypeString,
							Computed: true,
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
			"infrastructure": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
													Optional: true,
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
		},
	}
}

func dataSourceDatabaseRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	database, err := client.GetDatabaseByName(d.Get("name").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	if err := setResourceData(d, database); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("data-db-%s-%s", database.Name, database.EngineType))

	return diags
}
