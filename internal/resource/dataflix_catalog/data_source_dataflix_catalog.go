package dataflix_catalog

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	apiClient "terraform-provider-tessell/internal/client"
)

func DataSourceDataflixCatalog() *schema.Resource {
	return &schema.Resource{

		ReadContext: dataSourceDataflixCatalogRead,

		Schema: map[string]*schema.Schema{
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
			"time_zone": {
				Type:        schema.TypeString,
				Description: "Output timezone",
				Computed:    true,
			},
			"owner": {
				Type:        schema.TypeString,
				Description: "Owner of the Availability Machine",
				Computed:    true,
			},
			"pitr_catalog": {
				Type:        schema.TypeList,
				Description: "PITR availability catalog",
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
							Description: "",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"region": {
										Type:        schema.TypeString,
										Description: "Region name",
										Computed:    true,
									},
									"time_ranges": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"from_time": {
													Type:        schema.TypeString,
													Description: "PITR recovery from-time",
													Computed:    true,
												},
												"to_time": {
													Type:        schema.TypeString,
													Description: "PITR recovery to-time",
													Computed:    true,
												},
												"shared_with": {
													Type:        schema.TypeList,
													Description: "Tessell Entity ACL Sharing Summary Info",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"users": {
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
								},
							},
						},
					},
				},
			},
			"snapshot_catalog": {
				Type:        schema.TypeList,
				Description: "",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Description: "DB Service snapshot Id",
							Required:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "DB Service snapshot name",
							Computed:    true,
						},
						"description": {
							Type:        schema.TypeString,
							Description: "Description for the snapshot",
							Computed:    true,
						},
						"snapshot_time": {
							Type:        schema.TypeString,
							Description: "DB Service snapshot capture time",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
						"size": {
							Type:        schema.TypeInt,
							Description: "Database Backup size in bytes",
							Computed:    true,
						},
						"manual": {
							Type:        schema.TypeBool,
							Description: "Specifies whether the backup is captured manually",
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
						"databases": {
							Type:        schema.TypeList,
							Description: "The databases that are captured as part of the snapshot",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Description: "Databases Id",
										Required:    true,
									},
									"name": {
										Type:        schema.TypeString,
										Description: "Databases name",
										Computed:    true,
									},
									"status": {
										Type:        schema.TypeString,
										Description: "Databases status",
										Computed:    true,
									},
								},
							},
						},
						"shared_with": {
							Type:        schema.TypeList,
							Description: "Tessell Entity ACL Sharing Summary Info",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"users": {
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
		},
	}
}

func dataSourceDataflixCatalogRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	availabilityMachineId := d.Get("availability_machine_id").(string)

	response, _, err := client.GetDataflixCatalog(availabilityMachineId)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := setResourceData(d, response); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*response.AvailabilityMachineId)

	return diags
}
