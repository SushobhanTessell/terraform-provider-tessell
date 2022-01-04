package vpc_peering

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	apiClient "terraform-provider-tessell/internal/client"
)

func ResourceVPCPeering() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceVPCPeeringCreate,
		ReadContext:   resourceVPCPeeringRead,
		DeleteContext: resourceVPCPeeringDelete,

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
			"vpc_name": {
				Type:        schema.TypeString,
				Description: "Name of the source VPC to peer",
				Required:    true,
				ForceNew:    true,
			},
			"subscription_name": {
				Type:        schema.TypeString,
				Description: "Name of the subscription",
				Required:    true,
				ForceNew:    true,
			},
			"cloud_type": {
				Type:        schema.TypeString,
				Description: "Tessell supported cloud types",
				Required:    true,
				ForceNew:    true,
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringInSlice([]string{"AWS", "AZURE"}, false),
				),
			},
			"region": {
				Type:        schema.TypeString,
				Description: "Region of the Tessell VPC",
				Required:    true,
				ForceNew:    true,
			},
			"aws_client_vpc_info": {
				Type:          schema.TypeList,
				MaxItems:      1,
				Description:   "AWS VPC Peering Client Info",
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"azure_client_vpc_info"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_account_id": {
							Type:        schema.TypeString,
							Description: "AWS Account Id of the client VPC",
							Required:    true,
							ForceNew:    true,
						},
						"client_vpc_id": {
							Type:        schema.TypeString,
							Description: "Id of the client VPC",
							Required:    true,
							ForceNew:    true,
						},
						"client_vpc_region": {
							Type:        schema.TypeString,
							Description: "Region of the client VPC",
							Required:    true,
							ForceNew:    true,
						},
					},
				},
			},
			"azure_client_vpc_info": {
				Type:          schema.TypeList,
				MaxItems:      1,
				Description:   "Azure VPC Peering Client Info",
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"aws_client_vpc_info"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_subscription_id": {
							Type:        schema.TypeString,
							Description: "Azure Subscription Id of the client VPC",
							Required:    true,
							ForceNew:    true,
						},
						"client_resource_group": {
							Type:        schema.TypeString,
							Description: "Azure Resource Group of the client VPC",
							Required:    true,
							ForceNew:    true,
						},
						"client_vpc_name": {
							Type:        schema.TypeString,
							Description: "Name of the client VPC",
							Required:    true,
							ForceNew:    true,
						},
						"client_active_directory_tenant_id": {
							Type:        schema.TypeString,
							Description: "Tenant Id of the client Active Directory",
							Required:    true,
							ForceNew:    true,
						},
						"client_application_object_id": {
							Type:        schema.TypeString,
							Description: "Id of the client AD App Object",
							Required:    true,
							ForceNew:    true,
						},
					},
				},
			},
		},

		CustomizeDiff: func(c context.Context, d *schema.ResourceDiff, i interface{}) error {
			cloudType := d.Get("cloud_type").(string)
			awsClientVpcInfoEmpty := len(d.Get("aws_client_vpc_info").([]interface{})) == 0
			azureClientVpcInfoEmpty := len(d.Get("azure_client_vpc_info").([]interface{})) == 0

			if cloudType == "AWS" {
				if awsClientVpcInfoEmpty || !azureClientVpcInfoEmpty {
					return fmt.Errorf(
						"when cloud_type == 'AWS', 'azure_client_vpc_info' is not expected, but 'aws_client_vpc_info' is required",
					)
				}
			} else {
				if azureClientVpcInfoEmpty || !awsClientVpcInfoEmpty {
					return fmt.Errorf(
						"when cloud_type == 'AZURE', 'aws_client_vpc_info' is not expected, but 'azure_client_vpc_info' is required",
					)
				}
			}
			return nil
		},
	}
}

func resourceVPCPeeringCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	vpcName := d.Get("vpc_name").(string)
	cloudType := d.Get("cloud_type").(string)
	region := d.Get("region").(string)
	subscriptionName := d.Get("subscription_name").(string)

	if err := validate(d); err != nil {
		return diag.FromErr(err)
	}

	payload, err := formPayload(d)
	if err != nil {
		return diag.FromErr(err)
	}

	vpcPeeringCreationResponse, err := client.CreateVPCPeering(vpcName, payload)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("vpcPeering-%s-%s-%s-%s-%s", vpcPeeringCreationResponse.Name, vpcName, cloudType, region, subscriptionName))
	d.Set("name", vpcPeeringCreationResponse.Name)

	resourceVPCPeeringRead(ctx, d, meta)

	return diags
}

func resourceVPCPeeringRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	name := d.Get("name").(string)
	vpcName := d.Get("vpc_name").(string)
	cloudType := d.Get("cloud_type").(string)
	region := d.Get("region").(string)
	subscriptionName := d.Get("subscription_name").(string)

	vpcPeering, err := client.GetVPCPeering(name, vpcName, cloudType, region, subscriptionName)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := setResourceData(d, vpcPeering); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceVPCPeeringDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	name := d.Get("name").(string)
	vpcName := d.Get("vpc_name").(string)
	cloudType := d.Get("cloud_type").(string)
	region := d.Get("region").(string)
	subscriptionName := d.Get("subscription_name").(string)

	_, err := client.DeleteVPCPeering(name, vpcName, cloudType, region, subscriptionName)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
