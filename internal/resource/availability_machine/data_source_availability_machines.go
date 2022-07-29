package availability_machine

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	apiClient "terraform-provider-tessell/internal/client"
	"terraform-provider-tessell/internal/helper"
	"terraform-provider-tessell/internal/model"
)

func DataSourceAvailabilityMachines() *schema.Resource {
	return &schema.Resource{

		ReadContext: dataSourceAvailabilityMachinesRead,

		Schema: map[string]*schema.Schema{
			"availability_machines": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
						"tessell_service_id": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
						"service_name": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
						"tenant": {
							Type:        schema.TypeString,
							Description: "Dmm's tenancy details",
							Computed:    true,
						},
						"subscription": {
							Type:        schema.TypeString,
							Description: "Dmm's subscription name",
							Computed:    true,
						},
						"engine_type": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
						"user_id": {
							Type:        schema.TypeString,
							Description: "Data Management Machine's user",
							Computed:    true,
						},
						"owner": {
							Type:        schema.TypeString,
							Description: "Availability Machine's owner",
							Computed:    true,
						},
						"logged_in_user_role": {
							Type:        schema.TypeString,
							Description: "The role of the logged in user for accessing the Availability Machine",
							Computed:    true,
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
						"rpo_sla": {
							Type:        schema.TypeList,
							Description: "This is a definition for Tessell Availability Machine's cloud and region availability details",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"availability_machine_id": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
									"availability_machine": {
										Type:        schema.TypeString,
										Description: "Associated Availability Machine Name",
										Computed:    true,
									},
									"cloud_availability": {
										Type:        schema.TypeList,
										Description: "The availability location details: cloudAccount to region",
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
									"rpo_sla_status": {
										Type:        schema.TypeString,
										Description: "The availability status",
										Computed:    true,
									},
									"sla": {
										Type:        schema.TypeString,
										Description: "Associated SLA",
										Computed:    true,
									},
									"schedule": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"backup_start_time": {
													Type:        schema.TypeList,
													Description: "",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"hour": {
																Type:        schema.TypeInt,
																Description: "",
																Computed:    true,
															},
															"minute": {
																Type:        schema.TypeInt,
																Description: "",
																Computed:    true,
															},
														},
													},
												},
												"daily_schedule": {
													Type:        schema.TypeList,
													Description: "",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"backups_per_day": {
																Type:        schema.TypeInt,
																Description: "The number of backups to be captured per day, this is exclusive with 'backupStartTimes'",
																Computed:    true,
															},
															"backup_start_times": {
																Type:        schema.TypeList,
																Description: "List of times when backup(s) has to be captured at. If this is specified, the 'backupStartTime' (at top level) value would be overridern by the first entry in this list",
																Computed:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"hour": {
																			Type:        schema.TypeInt,
																			Description: "",
																			Computed:    true,
																		},
																		"minute": {
																			Type:        schema.TypeInt,
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
								},
							},
						},
						"daps": {
							Type:        schema.TypeList,
							Description: "",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
									"name": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
									"availability_machine_id": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
									"tessell_service_id": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
									"service_name": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
									"engine_type": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
									"content_type": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
									"status": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
									"content_info": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"sanitized_content": {
													Type:        schema.TypeList,
													Description: "",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"automated": {
																Type:        schema.TypeList,
																Description: "",
																Computed:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"sanitization_schedule_id": {
																			Type:        schema.TypeString,
																			Description: "Id of the sanitization schedule to process automated backups, required only if contentType = Sanitized.",
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
									"data_access_config": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"daily_backups": {
													Type:        schema.TypeInt,
													Description: "Number of daily backups to replicate",
													Computed:    true,
												},
											},
										},
									},
									"owner": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
									"logged_in_user_role": {
										Type:        schema.TypeString,
										Description: "The role of the logged in user for accessing the Availability Machine",
										Computed:    true,
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
									"date_created": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
									"date_modified": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
								},
							},
						},
						"clones": {
							Type:        schema.TypeList,
							Description: "Clone databases that are created from this Availability Machine",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
									"name": {
										Type:        schema.TypeString,
										Description: "Name of the clone database",
										Computed:    true,
									},
									"subscription": {
										Type:        schema.TypeString,
										Description: "Clone's subsription name",
										Computed:    true,
									},
									"status": {
										Type:        schema.TypeString,
										Description: "Status of the clone database",
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
									"clone_info": {
										Type:        schema.TypeMap,
										Description: "Miscellaneous information",
										Computed:    true,
									},
									"owner": {
										Type:        schema.TypeString,
										Description: "The user who created database clone",
										Computed:    true,
									},
									"date_created": {
										Type:        schema.TypeString,
										Description: "Timestamp when the entity was created",
										Computed:    true,
									},
								},
							},
						},
						"date_created": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
						"date_modified": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the Availability Machine",
				Optional:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "status",
				Optional:    true,
			},
			"engine_type": {
				Type:        schema.TypeString,
				Description: "Availaility Machine's engine-types",
				Optional:    true,
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

func dataSourceAvailabilityMachinesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	name := d.Get("name").(string)
	engineType := d.Get("engine_type").(string)
	owners := *helper.InterfaceToStringSlice(d.Get("owners"))
	loadAcls := d.Get("load_acls").(bool)
	status := d.Get("status").(string)

	response, _, err := client.GetAvailabilityMachines(name, status, engineType, loadAcls, owners)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := setDataSourceValues(d, response.Response); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("AvailabilityMachineList")

	return diags
}

func setDataSourceValues(d *schema.ResourceData, AvailabilityMachineList *[]model.TessellDmmServiceConsumerDTO) error {
	parsedAvailabilityMachineList := make([]interface{}, 0)

	if AvailabilityMachineList != nil {
		parsedAvailabilityMachineList = make([]interface{}, len(*AvailabilityMachineList))
		for i, AvailabilityMachine := range *AvailabilityMachineList {
			parsedAvailabilityMachineList[i] = map[string]interface{}{
				"id":                  AvailabilityMachine.Id,
				"tessell_service_id":  AvailabilityMachine.TessellServiceId,
				"service_name":        AvailabilityMachine.ServiceName,
				"tenant":              AvailabilityMachine.Tenant,
				"subscription":        AvailabilityMachine.Subscription,
				"engine_type":         AvailabilityMachine.EngineType,
				"status":              AvailabilityMachine.Status,
				"user_id":             AvailabilityMachine.UserId,
				"owner":               AvailabilityMachine.Owner,
				"logged_in_user_role": AvailabilityMachine.LoggedInUserRole,
				"shared_with":         []interface{}{parseEntityAclSharingInfo(AvailabilityMachine.SharedWith)},
				"cloud_availability":  parseCloudRegionInfo1List(AvailabilityMachine.CloudAvailability),
				"rpo_sla":             []interface{}{parseTessellDmmAvailabilityServiceView(AvailabilityMachine.RpoSla)},
				"daps":                parseTessellDapServiceDTOList(AvailabilityMachine.Daps),
				"clones":              parseTessellCloneSummaryInfoList(AvailabilityMachine.Clones),
				"date_created":        AvailabilityMachine.DateCreated,
				"date_modified":       AvailabilityMachine.DateModified,
			}
		}
	}

	if err := d.Set("availability_machines", parsedAvailabilityMachineList); err != nil {
		return err
	}
	return nil
}
