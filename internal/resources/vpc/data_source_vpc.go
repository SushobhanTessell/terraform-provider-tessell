package vpc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	apiClient "terraform-provider-tessell/internal/client"
)

func DataSourceVPC() *schema.Resource {
	return &schema.Resource{

		ReadContext: dataSourceVPCRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the VPC",
				Required:    true,
			},
			"cidr_block": {
				Type:        schema.TypeString,
				Description: "Cidr block of the VPC",
				Computed:    true,
			},
			"cloud_type": {
				Type:        schema.TypeString,
				Description: "Tessell supported cloud types",
				Required:    true,
			},
			"region": {
				Type:        schema.TypeString,
				Description: "Region of the VPC",
				Required:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Tessell Vpc Status",
				Computed:    true,
			},
			"subscription_name": {
				Type:        schema.TypeString,
				Description: "Subscription of the VPC",
				Required:    true,
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
	}
}

func dataSourceVPCRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	name := d.Get("name").(string)
	cloudType := d.Get("cloud_type").(string)
	region := d.Get("region").(string)
	subscriptionName := d.Get("subscription_name").(string)

	vpc, err := client.GetVPC(name, cloudType, region, subscriptionName)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := setResourceData(d, vpc); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(name)

	return diags
}
