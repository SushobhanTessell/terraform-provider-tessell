package dap

import (
	"context"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	apiClient "terraform-provider-tessell/internal/client"
)

func ResourceDAP() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceDAPCreate,
		ReadContext:   resourceDAPRead,
		UpdateContext: resourceDAPUpdate,
		DeleteContext: resourceDAPDelete,

		CustomizeDiff: customdiff.Sequence(
			func(c context.Context, diff *schema.ResourceDiff, v interface{}) error {
				if v, ok := diff.GetOk("name"); !ok {
					return fmt.Errorf("%v", v.(string))
				}
				if v, ok := diff.GetOk("name"); !ok {
					return fmt.Errorf("%v", v.(string))
				}
				return errors.New(`az_mode "cross-az" is not supported with num_cache_nodes = 1`)
			},
		),
		Schema: map[string]*schema.Schema{
			"availability_machine": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"date_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"date_modified": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dmm_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"engine_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"owner": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"shared_with_users": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"shared_with_user_groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"user_groups": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"users": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"target_cloud_locations": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"aws": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"azure": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"retention_config": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pitr_retention": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"days": {
										Type:        schema.TypeInt,
										Description: "Number of days for which the pitr backups to retain",
										Required:    true,
									},
								},
							},
						},
						"daily_retention": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"days": {
										Type:        schema.TypeInt,
										Description: "Number of days for which the daily backup to retain",
										Required:    true,
									},
								},
							},
						},
						"weekly_retention": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"weeks": {
										Type:        schema.TypeInt,
										Description: "Number of weeks for which the weekly backup to retain",
										Required:    true,
									},
									"days": {
										Type:        schema.TypeList,
										Description: "Which days of a week, the backup should be made available for the Data Access Policy",
										Required:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"monthly_retention": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"months": {
										Type:        schema.TypeInt,
										Description: "Number of months for which the monthly backup to retain",
										Required:    true,
									},
									"common_schedule": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dates": {
													Type:        schema.TypeList,
													Description: "Dates in a month to retain monthly backups for",
													Required:    true,
													Elem: &schema.Schema{
														Type: schema.TypeInt,
													},
												},
												"last_day_of_month": {
													Type:        schema.TypeBool,
													Description: "Number of months for which the monthly backup to retain",
													Required:    true,
												},
											},
										},
									},
									"month_specific_schedule": {
										Type:        schema.TypeList,
										Description: "Which days of a week, the backup should be made available for the Data Access Policy",
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dates": {
													Type:     schema.TypeList,
													Required: true,
													Elem: &schema.Schema{
														Type: schema.TypeInt,
													},
												},
												"month": {
													Type:     schema.TypeString,
													Required: true,
												},
											},
										},
									},
								},
							},
						},
						"yearly_retention": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"years": {
										Type:        schema.TypeInt,
										Description: "Number of years for which the yearly backup to retain",
										Required:    true,
									},
									"month_specific_schedule": {
										Type:        schema.TypeList,
										Description: "Which days of a week, the backup should be made available for the Data Access Policy",
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dates": {
													Type:     schema.TypeList,
													Required: true,
													Elem: &schema.Schema{
														Type: schema.TypeInt,
													},
												},
												"month": {
													Type:     schema.TypeString,
													Required: true,
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
	}
}

func resourceDAPCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	name := d.Get("name").(string)
	dmmName := d.Get("dmm_name").(string)

	if err := validate(d); err != nil {
		return diag.FromErr(err)
	}

	payload, err := formPayload(d)
	if err != nil {
		return diag.FromErr(err)
	}

	_, err = client.CreateDAP(dmmName, payload)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("dap-" + dmmName + name)

	resourceDAPRead(ctx, d, meta)

	return diags
}

func resourceDAPRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	dap, err := client.GetDAPByName(d.Get("name").(string), d.Get("dmm_name").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	if err := setResourceData(d, dap); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceDAPUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	if !d.HasChange("") {
		return diags
	}

	name := d.Get("name").(string)
	dmmName := d.Get("dmm_name").(string)

	if err := validate(d); err != nil {
		return diag.FromErr(err)
	}

	payload, err := formPayload(d)
	if err != nil {
		return diag.FromErr(err)
	}

	_, err = client.UpdateDAP(name, dmmName, payload)
	if err != nil {
		return diag.FromErr(err)
	}

	resourceDAPRead(ctx, d, meta)

	return diags
}

func resourceDAPDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	_, err := client.DeleteDAP(d.Get("name").(string), d.Get("dmm_name").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
