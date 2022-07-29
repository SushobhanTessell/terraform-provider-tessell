package db_service

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	apiClient "terraform-provider-tessell/internal/client"
)

func ResourceDBService() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceDBServiceCreate,
		ReadContext:   resourceDBServiceRead,
		UpdateContext: resourceDBServiceUpdate,
		DeleteContext: resourceDBServiceDelete,

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
			"snapshot_id": {
				Type:        schema.TypeString,
				Description: "Tessell service snapshot Id, using which the clone is to be created",
				Optional:    true,
				ForceNew:    true,
			},
			"pitr": {
				Type:        schema.TypeString,
				Description: "PITR Timestamp, using which the clone is to be created",
				Optional:    true,
				ForceNew:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the Tessell Service",
				Required:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Description: "Tessell Service's description",
				Optional:    true,
			},
			"tenant_id": {
				Type:        schema.TypeString,
				Description: "The tenant-id for the Tessell Service",
				Computed:    true,
			},
			"subscription": {
				Type:        schema.TypeString,
				Description: "Tessell Subscription in which the Tessell Service is to be created",
				Required:    true,
				ForceNew:    true,
			},
			"engine_type": {
				Type:        schema.TypeString,
				Description: "",
				Required:    true,
				ForceNew:    true,
			},
			"topology": {
				Type:        schema.TypeString,
				Description: "",
				Required:    true,
				ForceNew:    true,
			},
			"num_of_instances": {
				Type:        schema.TypeInt,
				Description: "Number of instance (nodes) to be created for the Tessell Service. This is a required input for Apache Kafka. For all other engines, this input would be ignored even if specified.",
				Optional:    true,
				ForceNew:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "",
				Computed:    true,
			},
			"license_type": {
				Type:        schema.TypeString,
				Description: "",
				Optional:    true,
				ForceNew:    true,
			},
			"software_image": {
				Type:        schema.TypeString,
				Description: "Software Image to be used to create the Tessell Service",
				Required:    true,
				ForceNew:    true,
			},
			"software_image_version": {
				Type:        schema.TypeString,
				Description: "Software Image Version to be used to create the Tessell Service",
				Required:    true,
				ForceNew:    true,
			},
			"auto_minor_version_update": {
				Type:        schema.TypeBool,
				Description: "Specify whether to automatically update minor version for Tessell Service",
				Optional:    true,
				Default:     true,
			},
			"enable_deletion_protection": {
				Type:        schema.TypeBool,
				Description: "Specify whether to enable deletion protection for the Tessell Service",
				Optional:    true,
				Default:     true,
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
							Optional:    true,
							ForceNew:    true,
						},
						"availability_machine_id": {
							Type:        schema.TypeString,
							Description: "The Availability Machine Id using which this Tessell Service clone is created",
							Optional:    true,
							ForceNew:    true,
						},
						"tessell_service": {
							Type:        schema.TypeString,
							Description: "The Tessell Service name using which this Tessell Service clone is created",
							Optional:    true,
							ForceNew:    true,
						},
						"availability_machine": {
							Type:        schema.TypeString,
							Description: "The Availaility Machine name using which this Tessell Service clone is created",
							Optional:    true,
							ForceNew:    true,
						},
						"snapshot_name": {
							Type:        schema.TypeString,
							Description: "The snapshot using which this Tessell Service clone is created",
							Optional:    true,
							ForceNew:    true,
						},
						"snapshot_id": {
							Type:        schema.TypeString,
							Description: "The snapshot Id using which this Tessell Service clone is created",
							Optional:    true,
							ForceNew:    true,
						},
						"pitr_time": {
							Type:        schema.TypeString,
							Description: "If the database was created using a Point-In-Time mechanism, it specifies the timestamp in UTC",
							Optional:    true,
							ForceNew:    true,
						},
					},
				},
			},
			"infrastructure": {
				Type:        schema.TypeList,
				Description: "The infra details where the Tessell Service is present",
				Required:    true,
				ForceNew:    true,
				MaxItems:    1,
				MinItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cloud": {
							Type:        schema.TypeString,
							Description: "The cloud-type in which the Tessell Service is provisioned (ex. aws, azure)",
							Optional:    true,
							ForceNew:    true,
						},
						"region": {
							Type:        schema.TypeString,
							Description: "The region in which the Tessell Service provisioned",
							Optional:    true,
							ForceNew:    true,
						},
						"availability_zone": {
							Type:        schema.TypeString,
							Description: "The availability-zone in which the Tessell Service is provisioned",
							Optional:    true,
							ForceNew:    true,
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
										Required:    true,
										ForceNew:    true,
									},
									"regions": {
										Type:        schema.TypeList,
										Description: "The regions details",
										Optional:    true,
										ForceNew:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"region": {
													Type:        schema.TypeString,
													Description: "The cloud region name",
													Required:    true,
													ForceNew:    true,
												},
												"availability_zones": {
													Type:        schema.TypeList,
													Description: "",
													Optional:    true,
													ForceNew:    true,
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
							Optional:    true,
							ForceNew:    true,
						},
						"compute_type": {
							Type:        schema.TypeString,
							Description: "The compute-type to be used for provisioning the Tessell Service",
							Optional:    true,
							ForceNew:    true,
						},
						"additional_storage": {
							Type:        schema.TypeInt,
							Description: "The additional storage (in GBs) to be provisioned for the Tessell Service. This is in addition to what is specified in the compute type.",
							Optional:    true,
							ForceNew:    true,
							Default:     0,
						},
					},
				},
			},
			"service_connectivity": {
				Type:        schema.TypeList,
				Description: "Tessell Service's connectivity information",
				Required:    true,
				MaxItems:    1,
				MinItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dns_prefix": {
							Type:        schema.TypeString,
							Description: "",
							Optional:    true,
						},
						"service_port": {
							Type:        schema.TypeInt,
							Description: "The connection port for the Tessell Service",
							Optional:    true,
							ForceNew:    true,
						},
						"enable_public_access": {
							Type:        schema.TypeBool,
							Description: "Specify whether to enable public access to the Tessell Service, default false",
							Optional:    true,
						},
						"allowed_ip_addresses": {
							Type:        schema.TypeList,
							Description: "The list of allowed ipv4 addresses that can connect to the Tessell Service",
							Optional:    true,
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
										Optional:    true,
										ForceNew:    true,
									},
									"connect_descriptor": {
										Type:        schema.TypeString,
										Description: "",
										Optional:    true,
										ForceNew:    true,
									},
									"endpoint": {
										Type:        schema.TypeString,
										Description: "",
										Optional:    true,
										ForceNew:    true,
									},
									"master_user": {
										Type:        schema.TypeString,
										Description: "",
										Optional:    true,
										ForceNew:    true,
									},
									"service_port": {
										Type:        schema.TypeInt,
										Description: "The connection port for the Tessell Service",
										Optional:    true,
										ForceNew:    true,
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
										Optional:    true,
										ForceNew:    true,
									},
									"enable_public_access": {
										Type:        schema.TypeBool,
										Description: "Specify whether to enable public access to the Tessell Service, default false",
										Optional:    true,
										ForceNew:    true,
									},
									"allowed_ip_addresses": {
										Type:        schema.TypeList,
										Description: "The list of allowed ipv4 addresses that can connect to the Tessell Service",
										Optional:    true,
										ForceNew:    true,
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
			"creds": {
				Type:        schema.TypeList,
				Description: "Tessell Service's credential details",
				Required:    true,
				ForceNew:    true,
				MaxItems:    1,
				MinItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"master_user": {
							Type:        schema.TypeString,
							Description: "Tessell Service's master username",
							Required:    true,
						},
						"master_password": {
							Type:        schema.TypeString,
							Description: "Tessell Service's master password",
							Required:    true,
						},
					},
				},
			},
			"maintenance_window": {
				Type:        schema.TypeList,
				Description: "Tessell Service's maintenance window details",
				Optional:    true,
				ForceNew:    true,
				MaxItems:    1,
				MinItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"day": {
							Type:        schema.TypeString,
							Description: "",
							Required:    true,
							ForceNew:    true,
						},
						"time": {
							Type:        schema.TypeString,
							Description: "Time value in (hh:mm) format. ex. \"02:00\"",
							Required:    true,
							ForceNew:    true,
						},
						"duration": {
							Type:        schema.TypeInt,
							Description: "",
							Required:    true,
							ForceNew:    true,
						},
					},
				},
			},
			"snapshot_configuration": {
				Type:        schema.TypeList,
				Description: "Tessell Service's backup configurations. If not specified, the default recommended backup configurations would be applied.",
				Optional:    true,
				MaxItems:    1,
				MinItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_snapshot": {
							Type:        schema.TypeBool,
							Description: "Specify whether to capture automated snapshots for the Tessell Service, default true.",
							Optional:    true,
							Default:     true,
						},
						"sla": {
							Type:        schema.TypeString,
							Description: "The snapshot SLA for the Tessell Service. If not specified, a default SLA would be associated with the Tessell Service",
							Optional:    true,
						},
						"snapshot_window": {
							Type:        schema.TypeList,
							Description: "",
							Optional:    true,
							MaxItems:    1,
							MinItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"time": {
										Type:        schema.TypeString,
										Description: "Time value in (hh:mm) format. ex. \"02:00\"",
										Optional:    true,
									},
									"duration": {
										Type:        schema.TypeInt,
										Description: "The allowed duration for capturing the Tessell Service backup",
										Optional:    true,
									},
								},
							},
						},
					},
				},
			},
			"engine_configuration": {
				Type:        schema.TypeList,
				Description: "",
				Required:    true,
				ForceNew:    true,
				MaxItems:    1,
				MinItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oracle_config": {
							Type:        schema.TypeList,
							Description: "",
							Optional:    true,
							ForceNew:    true,
							MaxItems:    1,
							MinItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"multi_tenant": {
										Type:        schema.TypeBool,
										Description: "Specify whether the Tessell Service is multi-tenant.",
										Optional:    true,
										ForceNew:    true,
										Default:     false,
									},
									"parameter_profile": {
										Type:        schema.TypeString,
										Description: "The parameter profile for the database",
										Optional:    true,
										ForceNew:    true,
									},
									"options_profile": {
										Type:        schema.TypeString,
										Description: "The options profile for the database",
										Optional:    true,
										ForceNew:    true,
									},
									"character_set": {
										Type:        schema.TypeString,
										Description: "The character-set for the database",
										Optional:    true,
										ForceNew:    true,
									},
									"national_character_set": {
										Type:        schema.TypeString,
										Description: "The national-character-set for the database",
										Optional:    true,
										ForceNew:    true,
									},
								},
							},
						},
						"postgresql_config": {
							Type:        schema.TypeList,
							Description: "",
							Optional:    true,
							ForceNew:    true,
							MaxItems:    1,
							MinItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"parameter_profile": {
										Type:        schema.TypeString,
										Description: "The parameter profile for the database",
										Optional:    true,
										ForceNew:    true,
									},
								},
							},
						},
						"mysql_config": {
							Type:        schema.TypeList,
							Description: "",
							Optional:    true,
							ForceNew:    true,
							MaxItems:    1,
							MinItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"parameter_profile": {
										Type:        schema.TypeString,
										Description: "The parameter profile for the database",
										Optional:    true,
										ForceNew:    true,
									},
								},
							},
						},
						"sql_server_config": {
							Type:        schema.TypeList,
							Description: "",
							Optional:    true,
							ForceNew:    true,
							MaxItems:    1,
							MinItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"parameter_profile": {
										Type:        schema.TypeString,
										Description: "The parameter profile for the database",
										Optional:    true,
										ForceNew:    true,
									},
								},
							},
						},
						"apache_kafka_config": {
							Type:        schema.TypeList,
							Description: "",
							Optional:    true,
							ForceNew:    true,
							MaxItems:    1,
							MinItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"parameter_profile": {
										Type:        schema.TypeString,
										Description: "The parameter profile for the database",
										Optional:    true,
										ForceNew:    true,
									},
								},
							},
						},
						"pre_script_info": {
							Type:        schema.TypeList,
							Description: "",
							Optional:    true,
							ForceNew:    true,
							MaxItems:    1,
							MinItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"script_id": {
										Type:        schema.TypeString,
										Description: "The Tessell Script Id",
										Optional:    true,
										ForceNew:    true,
									},
									"script_version": {
										Type:        schema.TypeString,
										Description: "The Tessell Script version",
										Optional:    true,
										ForceNew:    true,
									},
								},
							},
						},
						"post_script_info": {
							Type:        schema.TypeList,
							Description: "",
							Optional:    true,
							ForceNew:    true,
							MaxItems:    1,
							MinItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"script_id": {
										Type:        schema.TypeString,
										Description: "The Tessell Script Id",
										Optional:    true,
										ForceNew:    true,
									},
									"script_version": {
										Type:        schema.TypeString,
										Description: "The Tessell Script version",
										Optional:    true,
										ForceNew:    true,
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
				Optional:    true,
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
							Optional:    true,
							ForceNew:    true,
						},
						"description": {
							Type:        schema.TypeString,
							Description: "Database description",
							Optional:    true,
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
							Optional:    true,
							ForceNew:    true,
							MaxItems:    1,
							MinItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"oracle_config": {
										Type:        schema.TypeList,
										Description: "",
										Optional:    true,
										ForceNew:    true,
										MaxItems:    1,
										MinItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"parameter_profile": {
													Type:        schema.TypeString,
													Description: "The parameter profile for the database",
													Optional:    true,
													ForceNew:    true,
												},
												"options_profile": {
													Type:        schema.TypeString,
													Description: "The options profile for the database",
													Optional:    true,
													ForceNew:    true,
												},
											},
										},
									},
									"postgresql_config": {
										Type:        schema.TypeList,
										Description: "",
										Optional:    true,
										ForceNew:    true,
										MaxItems:    1,
										MinItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"parameter_profile": {
													Type:        schema.TypeString,
													Description: "The parameter profile for the database",
													Optional:    true,
													ForceNew:    true,
												},
											},
										},
									},
									"my_sql_config": {
										Type:        schema.TypeList,
										Description: "",
										Optional:    true,
										ForceNew:    true,
										MaxItems:    1,
										MinItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"parameter_profile": {
													Type:        schema.TypeString,
													Description: "The parameter profile for the database",
													Optional:    true,
													ForceNew:    true,
												},
											},
										},
									},
									"sql_server_config": {
										Type:        schema.TypeList,
										Description: "",
										Optional:    true,
										ForceNew:    true,
										MaxItems:    1,
										MinItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"parameter_profile": {
													Type:        schema.TypeString,
													Description: "The parameter profile for the database",
													Optional:    true,
													ForceNew:    true,
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
			"integrations_config": {
				Type:        schema.TypeList,
				Description: "",
				Optional:    true,
				MaxItems:    1,
				MinItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"integrations": {
							Type:        schema.TypeList,
							Description: "",
							Optional:    true,
							ForceNew:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if old == "0" && new == "1" {
						integrations := d.Get("integrations_config.0.integrations")
						if len(integrations.([]interface{})) == 0 {
							return true
						}
					} else if old == "1" && new == "0" {
						integrationsConfig := d.GetRawState().GetAttr("integrations_config").AsValueSlice()[0]
						integrations := integrationsConfig.GetAttr("integrations").AsValueSlice()
						if len(integrations) == 0 {
							return true
						}
					}
					return false
				},
			},
			"tags": {
				Type:        schema.TypeList,
				Description: "The tags to be associated with the Tessell Service",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Case sensitive, tag name",
							Optional:    true,
							ForceNew:    true,
						},
						"value": {
							Type:        schema.TypeString,
							Description: "Case sensitive, tag value",
							Optional:    true,
							ForceNew:    true,
						},
					},
				},
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if old == "0" && new == "1" {
						tags := d.Get("tags.0")
						if len(tags.(map[string]interface{})) == 0 {
							return true
						}
					} else if old == "1" && new == "0" {
						tags := d.GetRawState().GetAttr("tags").AsValueSlice()[0].AsValueMap()
						if len(tags) == 0 {
							return true
						}
					}
					return false
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
										Optional:    true,
										ForceNew:    true,
									},
									"reference_id": {
										Type:        schema.TypeString,
										Description: "The reference-id of the update request",
										Optional:    true,
										ForceNew:    true,
									},
									"submitted_at": {
										Type:        schema.TypeString,
										Description: "Timestamp when the resource update was requested",
										Optional:    true,
										ForceNew:    true,
									},
									"update_info": {
										Type:        schema.TypeMap,
										Description: "The specific details for a Tessell resource that are being updated",
										Optional:    true,
										ForceNew:    true,
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
				Optional:    true,
				ForceNew:    true,
				MaxItems:    1,
				MinItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"users": {
							Type:        schema.TypeList,
							Description: "",
							Optional:    true,
							ForceNew:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"email_id": {
										Type:        schema.TypeString,
										Description: "",
										Optional:    true,
										ForceNew:    true,
									},
									"role": {
										Type:        schema.TypeString,
										Description: "",
										Optional:    true,
										ForceNew:    true,
									},
								},
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
			"parent_availability_machine_id": {
				Type:        schema.TypeString,
				Description: "Id of the parent AvailabilityMachine, required when creating a clone",
				Optional:    true,
				ForceNew:    true,
			},
			"block_until_complete": {
				Type:        schema.TypeBool,
				Description: "For any operation on this resource, block the flow until the action has completed successfully",
				Optional:    true,
				Default:     true,
			},
			"timeout": {
				Type:        schema.TypeInt,
				Description: "If block_until_complete is true, how long it should block for. (In seconds)",
				Optional:    true,
				Default:     1200,
			},
			"expected_status": {
				Type:        schema.TypeString,
				Description: "If provided, invoke the DB Service start/stop API",
				Optional:    true,
			},
		},
	}
}

func resourceDBServiceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics
	var resourceId string

	availabilityMachineId := d.Get("availability_machine_id").(string)
	snapshotId := d.Get("snapshot_id").(string)
	pitr := d.Get("pitr").(string)

	if snapshotId != "" || pitr != "" {
		payload := formPayloadForCloneTessellService(d)

		response, _, err := client.CloneTessellService(availabilityMachineId, payload)
		if err != nil {
			return diag.FromErr(err)
		}
		resourceId = *response.ResourceId
	} else {
		payload := formPayloadForProvisionTessellService(d)

		response, _, err := client.ProvisionTessellService(payload)
		if err != nil {
			return diag.FromErr(err)
		}
		resourceId = *response.ResourceId
	}

	d.SetId(resourceId)

	if d.Get("block_until_complete").(bool) {
		//if err := client.WaitTillReady(resourceId, d.Get("timeout").(int)); err != nil {
		if err := client.DBServicePollForStatus(resourceId, "READY", d.Get("timeout").(int), 60); err != nil {
			return diag.FromErr(err)
		}
	}

	resourceDBServiceRead(ctx, d, meta)

	return diags
}

func resourceDBServiceRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	id := d.Get("id").(string)

	response, _, err := client.GetTessellService(id)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := setResourceData(d, response); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*response.Id)

	return diags
}
func resourceDBServiceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics
	var pollBreakValue string
	var pollFunc func(string, string, int, int) error
	shouldPoll := false

	expectedStatus := d.Get("expected_status").(string)
	id := d.Get("id").(string)

	if d.HasChanges("expected_status") && expectedStatus == "READY" {

		_, _, err := client.StartTessellService(id)
		if err != nil {
			return diag.FromErr(err)
		}
		shouldPoll = true
		pollBreakValue = "READY"
		pollFunc = client.DBServicePollForStatus
	} else if d.HasChanges("expected_status") && expectedStatus == "STOPPED" {

		_, _, err := client.StopTessellService(id)
		if err != nil {
			return diag.FromErr(err)
		}
		shouldPoll = true
		pollBreakValue = "STOPPED"
		pollFunc = client.DBServicePollForStatus
	}

	if shouldPoll {
		if err := pollFunc(d.Get("id").(string), pollBreakValue, d.Get("timeout").(int), 30); err != nil {
			return diag.FromErr(err)
		}
	}

	resourceDBServiceRead(ctx, d, meta)

	return diags
}

func resourceDBServiceDelete(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	id := d.Get("id").(string)

	payload := formPayloadForDeleteTessellService(d)

	response, statusCode, err := client.DeleteTessellService(id, payload)
	if err != nil {
		return diag.FromErr(err)
	}

	if statusCode != 200 {
		return diag.FromErr(fmt.Errorf("deletion failed for tessell_db_service with resourceId %s. Received response: %+v", id, response))
	}

	//err = client.WaitTillDeleted(databaseDeletionResponse.TaskId, d.Get("timeout").(int), "Database Deletion")
	err = client.DBServicePollForStatusCode(id, 404, d.Get("timeout").(int), 30)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
