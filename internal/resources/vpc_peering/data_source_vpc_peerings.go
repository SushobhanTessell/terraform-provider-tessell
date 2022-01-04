package vpc_peering

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	apiClient "terraform-provider-tessell/internal/client"
	"terraform-provider-tessell/internal/models"
)

func DataSourceVPCPeerings() *schema.Resource {
	return &schema.Resource{

		ReadContext: dataSourceVPCPeeringsRead,

		Schema: map[string]*schema.Schema{
			"vpc_name": {
				Type:        schema.TypeString,
				Description: "Name of the source VPC to peer",
				Required:    true,
			},
			"subscription_name": {
				Type:        schema.TypeString,
				Description: "Name of the subscription",
				Required:    true,
			},
			"cloud_type": {
				Type:        schema.TypeString,
				Description: "Tessell supported cloud types",
				Required:    true,
			},
			"region": {
				Type:        schema.TypeString,
				Description: "Region of the Tessell VPC",
				Required:    true,
			},
			"vpc_peerings": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Tessell VPC Peering Name",
							Computed:    true,
						},
						"cloud_id": {
							Type:        schema.TypeString,
							Description: "CloudId of Peering Connection",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Tessell VPC Peering Status",
							Computed:    true,
						},
						"metadata": {
							Type:        schema.TypeList,
							Description: "Metadata about Vpc Peering",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"azure_pending_peer_metadata": {
										Type:        schema.TypeList,
										Description: "Metadata about Azure VPC Peering in Pending Peer Status",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"tenant_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"vnet_resource_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"aws_client_info": {
							Type:        schema.TypeList,
							Description: "AWS VPC Peering Client Info for Service Consumer",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"client_account_id": {
										Type:        schema.TypeString,
										Description: "Account Id of the client VPC",
										Computed:    true,
									},
									"client_vpc_id": {
										Type:        schema.TypeString,
										Description: "Id of the client VPC",
										Computed:    true,
									},
									"client_vpc_region": {
										Type:        schema.TypeString,
										Description: "Region of the client VPC",
										Computed:    true,
									},
								},
							},
						},
						"azure_client_info": {
							Type:        schema.TypeList,
							Description: "Azure VPC Peering Client Info",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"client_subscription_id": {
										Type:        schema.TypeString,
										Description: "Azure Subscription Id of the client VPC",
										Computed:    true,
									},
									"client_resource_group": {
										Type:        schema.TypeString,
										Description: "Azure Resource Group of the client VPC",
										Computed:    true,
									},
									"client_vpc_name": {
										Type:        schema.TypeString,
										Description: "Name of the client VPC",
										Computed:    true,
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

func dataSourceVPCPeeringsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	vpcName := d.Get("vpc_name").(string)
	cloudType := d.Get("cloud_type").(string)
	region := d.Get("region").(string)
	subscriptionName := d.Get("subscription_name").(string)

	vpcPeerings, err := client.GetVPCPeerings(vpcName, cloudType, region, subscriptionName)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := setDataSourceValues(d, &vpcPeerings.VPCPeerings); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("data-vpcPeerings-%s-%s-%s-%s", vpcName, cloudType, region, subscriptionName))

	return diags
}

func setDataSourceValues(d *schema.ResourceData, vpcPeerings *[]models.VPCPeering) error {
	vpcPeeringList := make([]interface{}, 0)

	if vpcPeerings != nil {
		vpcPeeringList = make([]interface{}, len(*vpcPeerings))
		for i, vpcPeering := range *vpcPeerings {

			vpcPeeringList[i] = map[string]interface{}{
				"name":              vpcPeering.Name,
				"cloud_id":          vpcPeering.CloudId,
				"status":            vpcPeering.Status,
				"metadata":          parseMetadata(&vpcPeering),
				"aws_client_info":   parseAWSClientInfo(&vpcPeering.AWSClientInfo),
				"azure_client_info": parseAzureClientInfo(&vpcPeering.AzureClientInfo),
			}
		}
	}

	if err := d.Set("vpc_peerings", vpcPeeringList); err != nil {
		return err
	}
	return nil
}
