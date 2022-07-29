package dataflix

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	apiClient "terraform-provider-tessell/internal/client"
	"terraform-provider-tessell/internal/helper"
	"terraform-provider-tessell/internal/model"
)

func DataSourceDataflixs() *schema.Resource {
	return &schema.Resource{

		ReadContext: dataSourceDataflixsRead,

		Schema: map[string]*schema.Schema{
			"dataflixs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
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
						"owner": {
							Type:        schema.TypeString,
							Description: "Availability Machine's owner",
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
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the Availability Machine (Dataflix)",
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

func dataSourceDataflixsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	name := d.Get("name").(string)
	owners := *helper.InterfaceToStringSlice(d.Get("owners"))
	loadAcls := d.Get("load_acls").(bool)

	response, _, err := client.GetDataflixes(name, loadAcls, owners)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := setDataSourceValues(d, response.Response); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("DataflixList")

	return diags
}

func setDataSourceValues(d *schema.ResourceData, DataflixList *[]model.TessellAmDataflixDTO) error {
	parsedDataflixList := make([]interface{}, 0)

	if DataflixList != nil {
		parsedDataflixList = make([]interface{}, len(*DataflixList))
		for i, Dataflix := range *DataflixList {
			parsedDataflixList[i] = map[string]interface{}{
				"availability_machine_id": Dataflix.AvailabilityMachineId,
				"tessell_service_id":      Dataflix.TessellServiceId,
				"service_name":            Dataflix.ServiceName,
				"engine_type":             Dataflix.EngineType,
				"cloud_availability":      parseCloudRegionInfo1List(Dataflix.CloudAvailability),
				"owner":                   Dataflix.Owner,
				"shared_with":             []interface{}{parseEntityAclSharingSummaryInfo(Dataflix.SharedWith)},
			}
		}
	}

	if err := d.Set("dataflixs", parsedDataflixList); err != nil {
		return err
	}
	return nil
}
