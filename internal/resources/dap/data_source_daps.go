package dap

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	apiClient "terraform-provider-tessell/internal/client"
	"terraform-provider-tessell/internal/models"
)

func DataSourceDAPs() *schema.Resource {
	return &schema.Resource{

		ReadContext: dataSourceDAPsRead,

		Schema: map[string]*schema.Schema{
			"dmm_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"daps": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
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
							Computed: true,
						},
						"target_cloud_locations": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"aws": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"azure": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"retention_config": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"pitr_retention": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"days": {
													Type:        schema.TypeInt,
													Description: "Number of days for which the pitr backups to retain",
													Computed:    true,
												},
											},
										},
									},
									"daily_retention": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"days": {
													Type:        schema.TypeInt,
													Description: "Number of days for which the daily backup to retain",
													Computed:    true,
												},
											},
										},
									},
									"weekly_retention": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"weeks": {
													Type:        schema.TypeInt,
													Description: "Number of weeks for which the weekly backup to retain",
													Computed:    true,
												},
												"days": {
													Type:        schema.TypeList,
													Description: "Which days of a week, the backup should be made available for the Data Access Policy",
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"monthly_retention": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"months": {
													Type:        schema.TypeInt,
													Description: "Number of months for which the monthly backup to retain",
													Computed:    true,
												},
												"common_schedule": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"dates": {
																Type:        schema.TypeList,
																Description: "Dates in a month to retain monthly backups for",
																Computed:    true,
																Elem: &schema.Schema{
																	Type: schema.TypeInt,
																},
															},
															"last_day_of_month": {
																Type:        schema.TypeBool,
																Description: "Number of months for which the monthly backup to retain",
																Computed:    true,
															},
														},
													},
												},
												"month_specific_schedule": {
													Type:        schema.TypeList,
													Description: "Which days of a week, the backup should be made available for the Data Access Policy",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"dates": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeInt,
																},
															},
															"month": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"yearly_retention": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"years": {
													Type:        schema.TypeInt,
													Description: "Number of years for which the yearly backup to retain",
													Computed:    true,
												},
												"month_specific_schedule": {
													Type:        schema.TypeList,
													Description: "Which days of a week, the backup should be made available for the Data Access Policy",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"dates": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeInt,
																},
															},
															"month": {
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
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceDAPsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	daps, err := client.GetDAPs(d.Get("dmm_name").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	if err := setDataSourceValues(d, &daps.DAPs); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("daps")

	return diags
}

func setDataSourceValues(d *schema.ResourceData, daps *[]models.DAP) error {
	dapList := make([]interface{}, 0)

	if daps != nil {
		dapList = make([]interface{}, len(*daps))
		for i, dap := range *daps {

			dp := map[string]interface{}{
				"availability_machine":    dap.AvailabilityMachine,
				"date_created":            dap.DateCreated,
				"date_modified":           dap.DateModified,
				"engine_type":             dap.EngineType,
				"owner":                   dap.Owner,
				"status":                  dap.Status,
				"shared_with_users":       dap.SharedWithUsers,
				"shared_with_user_groups": dap.SharedWithUserGroups,
				"name":                    dap.Name,
				"target_cloud_locations":  parseTargetCloudLocations(&dap),
				"retention_config":        parseRetentionConfig(&dap),
			}
			dapList[i] = dp
		}
	}

	if err := d.Set("daps", dapList); err != nil {
		return err
	}
	return nil
}
