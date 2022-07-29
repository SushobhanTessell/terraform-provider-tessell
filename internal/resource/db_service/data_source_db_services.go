package db_service

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	apiClient "terraform-provider-tessell/internal/client"
	"terraform-provider-tessell/internal/helper"
	"terraform-provider-tessell/internal/model"
)

func DataSourceDBServices() *schema.Resource {
	return &schema.Resource{

		ReadContext: dataSourceDBServicesRead,

		Schema: map[string]*schema.Schema{
			"db_services": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Description: "Tessell generated UUID for the Tessell Service",
							Computed:    true,
						},
						"availability_machine_id": {
							Type:        schema.TypeString,
							Description: "Associated Availability Machine Id",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Name of the Tessell Service",
							Computed:    true,
						},
						"description": {
							Type:        schema.TypeString,
							Description: "Tessell Service description",
							Computed:    true,
						},
						"engine_type": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
						"topology": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
						"num_of_instances": {
							Type:        schema.TypeInt,
							Description: "",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
						"license_type": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
						"auto_minor_version_update": {
							Type:        schema.TypeBool,
							Description: "Specifies whether to automatically update minor version for Tessell Service",
							Computed:    true,
						},
						"enable_deletion_protection": {
							Type:        schema.TypeBool,
							Description: "Specifies whether to enable deletion protection for the Tessell Service",
							Computed:    true,
						},
						"software_image": {
							Type:        schema.TypeString,
							Description: "The Software Image that is used to create the Tessell Service",
							Computed:    true,
						},
						"software_image_version": {
							Type:        schema.TypeString,
							Description: "The Software Image version that is used to create the Tessell Service",
							Computed:    true,
						},
						"tenant_id": {
							Type:        schema.TypeString,
							Description: "The tenant-id for the Tessell Service",
							Computed:    true,
						},
						"subscription": {
							Type:        schema.TypeString,
							Description: "The subscription-name in which this Tessell Service is created",
							Computed:    true,
						},
						"owner": {
							Type:        schema.TypeString,
							Description: "Tessell Service owner email address",
							Computed:    true,
						},
						"logged_in_user_role": {
							Type:        schema.TypeString,
							Description: "Access role for the currently logged in user",
							Computed:    true,
						},
						"date_created": {
							Type:        schema.TypeString,
							Description: "Timestamp when the Tessell Service was created at",
							Computed:    true,
						},
						"started_at": {
							Type:        schema.TypeString,
							Description: "Timestamp when the Tessell Service was last started at",
							Computed:    true,
						},
						"stopped_at": {
							Type:        schema.TypeString,
							Description: "Timestamp when the Tessell Service was last stopped at",
							Computed:    true,
						},
						"cloned_from_info": {
							Type:        schema.TypeList,
							Description: "If the Tessell Service is created as a clone from some other Tessell Service, this section describes the parent Tessell Service and cloning details",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tessell_service_id": {
										Type:        schema.TypeString,
										Description: "The Tessell Service Id using which this Tessell Service clone is created",
										Computed:    true,
									},
									"availability_machine_id": {
										Type:        schema.TypeString,
										Description: "The Availability Machine Id using which this Tessell Service clone is created",
										Computed:    true,
									},
									"tessell_service": {
										Type:        schema.TypeString,
										Description: "The Tessell Service name using which this Tessell Service clone is created",
										Computed:    true,
									},
									"availability_machine": {
										Type:        schema.TypeString,
										Description: "The Availaility Machine name using which this Tessell Service clone is created",
										Computed:    true,
									},
									"snapshot_name": {
										Type:        schema.TypeString,
										Description: "The snapshot using which this Tessell Service clone is created",
										Computed:    true,
									},
									"snapshot_id": {
										Type:        schema.TypeString,
										Description: "The snapshot Id using which this Tessell Service clone is created",
										Computed:    true,
									},
									"pitr_time": {
										Type:        schema.TypeString,
										Description: "If the database was created using a Point-In-Time mechanism, it specifies the timestamp in UTC",
										Computed:    true,
									},
								},
							},
						},
						"service_connectivity": {
							Type:        schema.TypeList,
							Description: "Tessell Service's connectivity information",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dns_prefix": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
									"service_port": {
										Type:        schema.TypeInt,
										Description: "The connection port for the Tessell Service",
										Computed:    true,
									},
									"enable_public_access": {
										Type:        schema.TypeBool,
										Description: "Specify whether to enable public access to the Tessell Service, default false",
										Computed:    true,
									},
									"allowed_ip_addresses": {
										Type:        schema.TypeList,
										Description: "The list of allowed ipv4 addresses that can connect to the Tessell Service",
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"connect_strings": {
										Type:        schema.TypeList,
										Description: "The list of connect strings for the Tessell Service",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"type": {
													Type:        schema.TypeString,
													Description: "",
													Computed:    true,
												},
												"connect_descriptor": {
													Type:        schema.TypeString,
													Description: "",
													Computed:    true,
												},
												"endpoint": {
													Type:        schema.TypeString,
													Description: "",
													Computed:    true,
												},
												"master_user": {
													Type:        schema.TypeString,
													Description: "",
													Computed:    true,
												},
												"service_port": {
													Type:        schema.TypeInt,
													Description: "The connection port for the Tessell Service",
													Computed:    true,
												},
											},
										},
									},
									"update_in_progress_info": {
										Type:        schema.TypeList,
										Description: "Tessell Service connectivity update-in-progress details",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dns_prefix": {
													Type:        schema.TypeString,
													Description: "",
													Computed:    true,
												},
												"enable_public_access": {
													Type:        schema.TypeBool,
													Description: "Specify whether to enable public access to the Tessell Service, default false",
													Computed:    true,
												},
												"allowed_ip_addresses": {
													Type:        schema.TypeList,
													Description: "The list of allowed ipv4 addresses that can connect to the Tessell Service",
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
								},
							},
						},
						"infrastructure": {
							Type:        schema.TypeList,
							Description: "The infra details where the Tessell Service is present",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cloud": {
										Type:        schema.TypeString,
										Description: "The cloud-type in which the Tessell Service is provisioned (ex. aws, azure)",
										Computed:    true,
									},
									"region": {
										Type:        schema.TypeString,
										Description: "The region in which the Tessell Service provisioned",
										Computed:    true,
									},
									"availability_zone": {
										Type:        schema.TypeString,
										Description: "The availability-zone in which the Tessell Service is provisioned",
										Computed:    true,
									},
									"cloud_availability": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"cloud": {
													Type:        schema.TypeString,
													Description: "",
													Computed:    true,
												},
												"regions": {
													Type:        schema.TypeList,
													Description: "The regions details",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"region": {
																Type:        schema.TypeString,
																Description: "The cloud region name",
																Computed:    true,
															},
															"availability_zones": {
																Type:        schema.TypeList,
																Description: "",
																Computed:    true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},
											},
										},
									},
									"vpc": {
										Type:        schema.TypeString,
										Description: "The VPC to be used for provisioning the Tessell Service",
										Computed:    true,
									},
									"compute_type": {
										Type:        schema.TypeString,
										Description: "The compute-type to be used for provisioning the Tessell Service",
										Computed:    true,
									},
									"additional_storage": {
										Type:        schema.TypeInt,
										Description: "The additional storage (in GBs) to be provisioned for the Tessell Service. This is in addition to what is specified in the compute type.",
										Computed:    true,
									},
								},
							},
						},
						"maintenance_window": {
							Type:        schema.TypeList,
							Description: "Tessell Service's maintenance window details",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"day": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
									"time": {
										Type:        schema.TypeString,
										Description: "Time value in (hh:mm) format. ex. \"02:00\"",
										Computed:    true,
									},
									"duration": {
										Type:        schema.TypeInt,
										Description: "",
										Computed:    true,
									},
								},
							},
						},
						"engine_configuration": {
							Type:        schema.TypeList,
							Description: "",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"oracle_config": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"multi_tenant": {
													Type:        schema.TypeBool,
													Description: "Specify whether the Tessell Service is multi-tenant.",
													Computed:    true,
												},
												"parameter_profile": {
													Type:        schema.TypeString,
													Description: "The parameter profile for the database",
													Computed:    true,
												},
												"options_profile": {
													Type:        schema.TypeString,
													Description: "The options profile for the database",
													Computed:    true,
												},
												"character_set": {
													Type:        schema.TypeString,
													Description: "The character-set for the database",
													Computed:    true,
												},
												"national_character_set": {
													Type:        schema.TypeString,
													Description: "The national-character-set for the database",
													Computed:    true,
												},
											},
										},
									},
									"postgresql_config": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"parameter_profile": {
													Type:        schema.TypeString,
													Description: "The parameter profile for the database",
													Computed:    true,
												},
											},
										},
									},
									"mysql_config": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"parameter_profile": {
													Type:        schema.TypeString,
													Description: "The parameter profile for the database",
													Computed:    true,
												},
											},
										},
									},
									"sql_server_config": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"parameter_profile": {
													Type:        schema.TypeString,
													Description: "The parameter profile for the database",
													Computed:    true,
												},
											},
										},
									},
									"apache_kafka_config": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"parameter_profile": {
													Type:        schema.TypeString,
													Description: "The parameter profile for the database",
													Computed:    true,
												},
											},
										},
									},
									"pre_script_info": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"script_id": {
													Type:        schema.TypeString,
													Description: "The Tessell Script Id",
													Computed:    true,
												},
												"script_version": {
													Type:        schema.TypeString,
													Description: "The Tessell Script version",
													Computed:    true,
												},
											},
										},
									},
									"post_script_info": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"script_id": {
													Type:        schema.TypeString,
													Description: "The Tessell Script Id",
													Computed:    true,
												},
												"script_version": {
													Type:        schema.TypeString,
													Description: "The Tessell Script version",
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},
						"integrations_config": {
							Type:        schema.TypeList,
							Description: "",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"integrations": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"deletion_config": {
							Type:        schema.TypeList,
							Description: "",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"retain_availability_machine": {
										Type:        schema.TypeBool,
										Description: "If 'retainAvailabilityMachine' is true then set value of field takeFinalBackup and dapsToRetain. By default retainAvailabilityMachine is false, that means delete all details like Availability Machine, Backups, DAPs etc.",
										Computed:    true,
									},
								},
							},
						},
						"tags": {
							Type:        schema.TypeList,
							Description: "The tags associated with the Tessell Service",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Case sensitive, tag name",
										Computed:    true,
									},
									"value": {
										Type:        schema.TypeString,
										Description: "Case sensitive, tag value",
										Computed:    true,
									},
								},
							},
						},
						"instances": {
							Type:        schema.TypeList,
							Description: "Instances associated with this Tessell Service",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Description: "Tessell generated UUID for the Tessell Service Instance",
										Computed:    true,
									},
									"name": {
										Type:        schema.TypeString,
										Description: "Name of the Tessell Service Instance",
										Computed:    true,
									},
									"role": {
										Type:        schema.TypeString,
										Description: "Tessell Service Topology",
										Computed:    true,
									},
									"status": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
									"tessell_service_id": {
										Type:        schema.TypeString,
										Description: "Tessell Service Instance's associated Tessell Service id",
										Computed:    true,
									},
									"compute_type": {
										Type:        schema.TypeString,
										Description: "The compute used for creation of the Tessell Service Instance",
										Computed:    true,
									},
									"cloud": {
										Type:        schema.TypeString,
										Description: "Tessell Service Instance's cloud type",
										Computed:    true,
									},
									"region": {
										Type:        schema.TypeString,
										Description: "Tessell Service Instance's cloud region",
										Computed:    true,
									},
									"availability_zone": {
										Type:        schema.TypeString,
										Description: "Tessell Service Instance's cloud availability zone",
										Computed:    true,
									},
									"date_created": {
										Type:        schema.TypeString,
										Description: "Timestamp when the entity was created",
										Computed:    true,
									},
									"connect_string": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"connect_descriptor": {
													Type:        schema.TypeString,
													Description: "",
													Computed:    true,
												},
												"master_user": {
													Type:        schema.TypeString,
													Description: "",
													Computed:    true,
												},
												"endpoint": {
													Type:        schema.TypeString,
													Description: "",
													Computed:    true,
												},
												"service_port": {
													Type:        schema.TypeString,
													Description: "",
													Computed:    true,
												},
											},
										},
									},
									"updates_in_progress": {
										Type:        schema.TypeList,
										Description: "The updates that are in progress for this resource",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"update_type": {
													Type:        schema.TypeString,
													Description: "Type of the update",
													Computed:    true,
												},
												"reference_id": {
													Type:        schema.TypeString,
													Description: "The reference-id of the update request",
													Computed:    true,
												},
												"submitted_at": {
													Type:        schema.TypeString,
													Description: "Timestamp when the resource update was requested",
													Computed:    true,
												},
												"update_info": {
													Type:        schema.TypeMap,
													Description: "The specific details for a Tessell resource that are being updated",
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},
						"databases": {
							Type:        schema.TypeList,
							Description: "Databases that are part of this Tessell Service",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
									"database_name": {
										Type:        schema.TypeString,
										Description: "Database name",
										Computed:    true,
									},
									"description": {
										Type:        schema.TypeString,
										Description: "Database description",
										Computed:    true,
									},
									"tessell_service_id": {
										Type:        schema.TypeString,
										Description: "Associated Tessell Service Id",
										Computed:    true,
									},
									"engine_type": {
										Type:        schema.TypeString,
										Description: "Database engine type",
										Computed:    true,
									},
									"status": {
										Type:        schema.TypeString,
										Description: "Database status",
										Computed:    true,
									},
									"date_created": {
										Type:        schema.TypeString,
										Description: "Timestamp when the entity was created",
										Computed:    true,
									},
									"cloned_from_info": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"database_id": {
													Type:        schema.TypeString,
													Description: "The original database Id using which this database clone is created",
													Computed:    true,
												},
											},
										},
									},
									"database_configuration": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"oracle_config": {
													Type:        schema.TypeList,
													Description: "",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"parameter_profile": {
																Type:        schema.TypeString,
																Description: "The parameter profile for the database",
																Computed:    true,
															},
															"options_profile": {
																Type:        schema.TypeString,
																Description: "The options profile for the database",
																Computed:    true,
															},
														},
													},
												},
												"postgresql_config": {
													Type:        schema.TypeList,
													Description: "",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"parameter_profile": {
																Type:        schema.TypeString,
																Description: "The parameter profile for the database",
																Computed:    true,
															},
														},
													},
												},
												"my_sql_config": {
													Type:        schema.TypeList,
													Description: "",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"parameter_profile": {
																Type:        schema.TypeString,
																Description: "The parameter profile for the database",
																Computed:    true,
															},
														},
													},
												},
												"sql_server_config": {
													Type:        schema.TypeList,
													Description: "",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"parameter_profile": {
																Type:        schema.TypeString,
																Description: "The parameter profile for the database",
																Computed:    true,
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
						"shared_with": {
							Type:        schema.TypeList,
							Description: "Tessell Entity ACL Sharing Info",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"users": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"email_id": {
													Type:        schema.TypeString,
													Description: "",
													Computed:    true,
												},
												"role": {
													Type:        schema.TypeString,
													Description: "",
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the Tessell Service",
				Optional:    true,
			},
			"statuses": {
				Type:        schema.TypeList,
				Description: "statuses",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"engine_types": {
				Type:        schema.TypeList,
				Description: "Tessell Service's engine-types",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"cloned_from_service_id": {
				Type:        schema.TypeString,
				Description: "The id of the Tessell Service from which the services are cloned",
				Optional:    true,
			},
			"cloned_from_availability_machine_id": {
				Type:        schema.TypeString,
				Description: "The id of the Availability Machine from which the services are cloned",
				Optional:    true,
			},
			"load_instances": {
				Type:        schema.TypeBool,
				Description: "Load the instances that are part of the Tessell Service",
				Optional:    true,
				Default:     true,
			},
			"load_databases": {
				Type:        schema.TypeBool,
				Description: "Load the databases that are part of the Tessell Service",
				Optional:    true,
				Default:     true,
			},
			"owners": {
				Type:        schema.TypeList,
				Description: "List of Email Addresses for entity or resource owners",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"load_acls": {
				Type:        schema.TypeBool,
				Description: "Load ACL information",
				Optional:    true,
				Default:     false,
			},
		},
	}
}

func dataSourceDBServicesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	loadDatabases := d.Get("load_databases").(bool)
	clonedFromAvailabilityMachineId := d.Get("cloned_from_availability_machine_id").(string)
	loadAcls := d.Get("load_acls").(bool)
	name := d.Get("name").(string)
	statuses := *helper.InterfaceToStringSlice(d.Get("statuses"))
	owners := *helper.InterfaceToStringSlice(d.Get("owners"))
	engineTypes := *helper.InterfaceToStringSlice(d.Get("engine_types"))
	clonedFromServiceId := d.Get("cloned_from_service_id").(string)
	loadInstances := d.Get("load_instances").(bool)

	response, _, err := client.GetTessellServices(name, statuses, engineTypes, clonedFromServiceId, clonedFromAvailabilityMachineId, loadInstances, loadDatabases, owners, loadAcls)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := setDataSourceValues(d, response.Response); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("DBServiceList")

	return diags
}

func setDataSourceValues(d *schema.ResourceData, DBServiceList *[]model.TessellServiceDTO) error {
	parsedDBServiceList := make([]interface{}, 0)

	if DBServiceList != nil {
		parsedDBServiceList = make([]interface{}, len(*DBServiceList))
		for i, DBService := range *DBServiceList {
			parsedDBServiceList[i] = map[string]interface{}{
				"id":                         DBService.Id,
				"availability_machine_id":    DBService.AvailabilityMachineId,
				"name":                       DBService.Name,
				"description":                DBService.Description,
				"engine_type":                DBService.EngineType,
				"topology":                   DBService.Topology,
				"num_of_instances":           DBService.NumOfInstances,
				"status":                     DBService.Status,
				"license_type":               DBService.LicenseType,
				"auto_minor_version_update":  DBService.AutoMinorVersionUpdate,
				"enable_deletion_protection": DBService.EnableDeletionProtection,
				"software_image":             DBService.SoftwareImage,
				"software_image_version":     DBService.SoftwareImageVersion,
				"tenant_id":                  DBService.TenantId,
				"subscription":               DBService.Subscription,
				"owner":                      DBService.Owner,
				"logged_in_user_role":        DBService.LoggedInUserRole,
				"date_created":               DBService.DateCreated,
				"started_at":                 DBService.StartedAt,
				"stopped_at":                 DBService.StoppedAt,
				"cloned_from_info":           []interface{}{parseTessellServiceClonedFromInfo(DBService.ClonedFromInfo)},
				"service_connectivity":       []interface{}{parseTessellServiceConnectivityInfo(DBService.ServiceConnectivity)},
				"infrastructure":             []interface{}{parseTessellServiceInfrastructureInfo(DBService.Infrastructure)},
				"maintenance_window":         []interface{}{parseTessellServiceMaintenanceWindow(DBService.MaintenanceWindow)},
				"engine_configuration":       []interface{}{parseTessellServiceEngineInfo(DBService.EngineConfiguration)},
				"integrations_config":        []interface{}{parseTessellServiceIntegrationsInfo(DBService.IntegrationsConfig)},
				"deletion_config":            []interface{}{parseTessellServiceDeletionConfig(DBService.DeletionConfig)},
				"tags":                       parseTessellTagList(DBService.Tags),
				"instances":                  parseTessellServiceInstanceDTOList(DBService.Instances),
				"databases":                  parseTessellDatabaseDTOList(DBService.Databases),
				"shared_with":                []interface{}{parseEntityAclSharingInfo(DBService.SharedWith)},
			}
		}
	}

	if err := d.Set("db_services", parsedDBServiceList); err != nil {
		return err
	}
	return nil
}
