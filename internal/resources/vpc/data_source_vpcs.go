package vpc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	apiClient "terraform-provider-tessell/internal/client"
	"terraform-provider-tessell/internal/models"
)

func DataSourceVPCs() *schema.Resource {
	return &schema.Resource{
		Description: "Sample data source in the Terraform provider Database.",

		ReadContext: dataSourceVPCsRead,

		Schema: map[string]*schema.Schema{
			"vpcs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Name of the VPC",
							Computed:    true,
						},
						"cidr_block": {
							Type:        schema.TypeString,
							Description: "Cidr block of the VPC",
							Computed:    true,
						},
						"cloud_type": {
							Type:        schema.TypeString,
							Description: "Tessell supported cloud types",
							Computed:    true,
						},
						"region": {
							Type:        schema.TypeString,
							Description: "Region of the VPC",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Tessell Vpc Status",
							Computed:    true,
						},
						"subscription_name": {
							Type:        schema.TypeString,
							Description: "Subscription of the VPC",
							Computed:    true,
						},
						"metadata": {
							Type:        schema.TypeList,
							Description: "Tessell Metadata for Vpc",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"validation_failure_reason": {
										Type:     schema.TypeString,
										Computed: true,
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

func dataSourceVPCsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	vpcs, err := client.GetVPCs()
	if err != nil {
		return diag.FromErr(err)
	}

	if err := setDataSourceValues(d, &vpcs.VPCs); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("vpcs")

	return diags
}

func setDataSourceValues(d *schema.ResourceData, vpcs *[]models.VPC) error {
	vpcList := make([]interface{}, 0)

	if vpcs != nil {
		vpcList = make([]interface{}, len(*vpcs))
		for i, vpc := range *vpcs {
			vpcList[i] = map[string]interface{}{
				"name":              vpc.Name,
				"cidr_block":        vpc.CidrBlock,
				"cloud_type":        vpc.CloudType,
				"region":            vpc.Region,
				"status":            vpc.Status,
				"subscription_name": vpc.SubscriptionName,
				"metadata":          parseMetadata(&vpc),
			}
		}
	}

	if err := d.Set("vpcs", vpcList); err != nil {
		return err
	}
	return nil
}
