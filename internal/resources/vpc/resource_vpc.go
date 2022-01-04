package vpc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	apiClient "terraform-provider-tessell/internal/client"
)

func ResourceVPC() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceVPCCreate,
		ReadContext:   resourceVPCRead,
		UpdateContext: resourceVPCUpdate,
		DeleteContext: resourceVPCDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the VPC",
				Required:    true,
				ForceNew:    true,
			},
			"cidr_block": {
				Type:        schema.TypeString,
				Description: "Cidr block of the VPC",
				Required:    true,
				ForceNew:    true,
			},
			"cloud_type": {
				Type:        schema.TypeString,
				Description: "Tessell supported cloud types",
				Required:    true,
				ForceNew:    true,
			},
			"region": {
				Type:        schema.TypeString,
				Description: "Region of the VPC",
				Required:    true,
				ForceNew:    true,
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
				ForceNew:    true,
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
			"block_until_complete": {
				Type:        schema.TypeBool,
				Description: "Block the flow until the VPC has been successfully completed",
				Optional:    true,
				Default:     false,
			},
			"timeout": {
				Type:        schema.TypeInt,
				Description: "If block_until_complete is true, how long should it block for. (In seconds)",
				Optional:    true,
				Default:     600,
			},
		},
	}
}

func resourceVPCCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	name := d.Get("name").(string)

	if err := validate(d); err != nil {
		return diag.FromErr(err)
	}

	payload, err := formPayload(d)
	if err != nil {
		return diag.FromErr(err)
	}

	_, err = client.CreateVPC(payload)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(name)

	if d.Get("block_until_complete").(bool) {
		if err := waitAndCheckIfCreationSucceeded(d.Get("timeout").(int), d, meta); err != nil {
			return diag.FromErr(err)
		}
	}

	resourceVPCRead(ctx, d, meta)

	return diags
}

func resourceVPCRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	vpc, err := client.GetVPC(d.Get("name").(string), d.Get("cloud_type").(string), d.Get("region").(string), d.Get("subscription_name").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	if err := setResourceData(d, vpc); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceVPCUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return *new(diag.Diagnostics)
}

func resourceVPCDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	_, err := client.DeleteVPC(d.Get("name").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
