package dap

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-tessell/internal/models"
	"terraform-provider-tessell/internal/utils"
)

func validate(d *schema.ResourceData) error {
	return nil
}

func formPayload(d *schema.ResourceData) (models.DAPCreationUpdationPayload, error) {
	payload := models.DAPCreationUpdationPayload{
		Name:                 d.Get("name").(string),
		UserGroups:           utils.InteraceSliceToStringSlice(d.Get("user_groups").([]interface{})),
		Users:                utils.InteraceSliceToStringSlice(d.Get("users").([]interface{})),
		TargetCloudLocations: models.DAPTargetCloudLocations{},
		RetentionConfig:      models.DAPRetentionConfig{},
	}

	targetCloudLocations := ((d.Get("target_cloud_locations").([]interface{}))[0]).(map[string]interface{})
	if v, ok := targetCloudLocations["aws"]; ok {
		payload.TargetCloudLocations.AWS = utils.InteraceSliceToStringSlice(v.([]interface{}))
	}
	if v, ok := targetCloudLocations["azure"]; ok {
		payload.TargetCloudLocations.Azure = utils.InteraceSliceToStringSlice(v.([]interface{}))
	}

	retentionConfig := ((d.Get("retention_config").([]interface{}))[0]).(map[string]interface{})
	if v, ok := retentionConfig["pitr_retention"]; ok {
		v2 := v.([]interface{})
		if len(v2) > 0 {
			payload.RetentionConfig.PitrRetention = models.DAPPitrRetention{
				Days: (v2[0]).(map[string]interface{})["days"].(int),
			}
		}
	}
	if v, ok := retentionConfig["daily_retention"]; ok {
		v2 := v.([]interface{})
		if len(v2) > 0 {
			payload.RetentionConfig.DailyRetention = models.DAPDailyRetention{
				Days: (v2[0]).(map[string]interface{})["days"].(int),
			}
		}
	}
	if v, ok := retentionConfig["weekly_retention"]; ok {
		v2 := v.([]interface{})
		if len(v2) > 0 {
			payload.RetentionConfig.WeeklyRetention = models.DAPWeeklyRetention{
				Weeks: (v2[0]).(map[string]interface{})["weeks"].(int),
				Days:  utils.InteraceSliceToStringSlice((v2[0]).(map[string]interface{})["days"].([]interface{})),
			}
		}
	}
	if v, ok := retentionConfig["monthly_retention"]; ok {
		v2 := v.([]interface{})
		if len(v2) > 0 {
			commonSchedule := ((v2[0]).(map[string]interface{})["common_schedule"]).([]interface{})
			payload.RetentionConfig.MonthlyRetention = models.DAPMonthlyRetention{
				Months:                (v2[0]).(map[string]interface{})["months"].(int),
				MonthSpecificSchedule: formPayloadMonthSpecificSchedule((v2[0]).(map[string]interface{})["month_specific_schedule"].([]interface{})),
			}
			if len(commonSchedule) > 0 {
				commSchedule := models.DAPMonthlyRetentionCommonSchedule{
					Dates:          utils.InteraceSliceToIntSlice((commonSchedule[0]).(map[string]interface{})["dates"].([]interface{})),
					LastDayOfMonth: (commonSchedule[0]).(map[string]interface{})["last_day_of_month"].(bool),
				}
				payload.RetentionConfig.MonthlyRetention.CommonSchedule = commSchedule
			}
		}
	}
	if v, ok := retentionConfig["yearly_retention"]; ok {
		v2 := v.([]interface{})
		if len(v2) > 0 {
			payload.RetentionConfig.YearlyRetention = models.DAPYearlyRetention{
				Years:                 (v2[0]).(map[string]interface{})["years"].(int),
				MonthSpecificSchedule: formPayloadMonthSpecificSchedule((v2[0]).(map[string]interface{})["month_specific_schedule"].([]interface{})),
			}
		}
	}

	return payload, nil
}

func formPayloadMonthSpecificSchedule(v []interface{}) []models.DAPMonthSpecificSchedule {
	schedules := []models.DAPMonthSpecificSchedule{}

	for _, s := range v {
		schedule := models.DAPMonthSpecificSchedule{
			Month: (s.(map[string]interface{}))["month"].(string),
			Dates: utils.InteraceSliceToIntSlice((s.(map[string]interface{}))["dates"].([]interface{})),
		}
		schedules = append(schedules, schedule)
	}

	return schedules
}

func setResourceData(d *schema.ResourceData, dap *models.DAP) error {
	if err := d.Set("availability_machine", dap.AvailabilityMachine); err != nil {
		return err
	}

	if err := d.Set("date_created", dap.DateCreated); err != nil {
		return err
	}

	if err := d.Set("date_modified", dap.DateModified); err != nil {
		return err
	}

	if err := d.Set("engine_type", dap.EngineType); err != nil {
		return err
	}

	if err := d.Set("name", dap.Name); err != nil {
		return err
	}

	if err := d.Set("owner", dap.Owner); err != nil {
		return err
	}

	if err := d.Set("shared_with_users", dap.SharedWithUsers); err != nil {
		return err
	}

	if err := d.Set("shared_with_user_groups", dap.SharedWithUserGroups); err != nil {
		return err
	}

	if err := d.Set("status", dap.Status); err != nil {
		return err
	}

	if err := d.Set("target_cloud_locations", parseTargetCloudLocations(dap)); err != nil {
		return err
	}

	if err := d.Set("retention_config", parseRetentionConfig(dap)); err != nil {
		return err
	}

	return nil
}

func parseTargetCloudLocations(dap *models.DAP) []interface{} {
	targetCloud := make(map[string]interface{})

	if awsLocations := dap.TargetCloudLocations.AWS; awsLocations != nil {
		targetCloud["aws"] = awsLocations[:]
	}
	if azureLocations := dap.TargetCloudLocations.Azure; azureLocations != nil {
		targetCloud["azure"] = azureLocations[:]
	}

	return []interface{}{targetCloud}
}

func parseRetentionConfig(dapRcvd *models.DAP) []interface{} {

	retentionCfg := dapRcvd.RetentionConfig
	return []interface{}{map[string]interface{}{
		"pitr_retention": []interface{}{map[string]interface{}{
			"days": retentionCfg.PitrRetention.Days,
		}},
		"daily_retention": []interface{}{map[string]interface{}{
			"days": retentionCfg.DailyRetention.Days,
		}},
		"weekly_retention": []interface{}{map[string]interface{}{
			"weeks": retentionCfg.WeeklyRetention.Weeks,
			"days":  utils.ToInterfaceSlice(retentionCfg.WeeklyRetention.Days),
		}},
		"monthly_retention": []interface{}{map[string]interface{}{
			"months": int(retentionCfg.MonthlyRetention.Months),
			"common_schedule": []interface{}{map[string]interface{}{
				"dates":             utils.ToInterfaceSlice(retentionCfg.MonthlyRetention.CommonSchedule.Dates),
				"last_day_of_month": retentionCfg.MonthlyRetention.CommonSchedule.LastDayOfMonth,
			}},
			"month_specific_schedule": utils.ToInterfaceSlice(retentionCfg.MonthlyRetention.MonthSpecificSchedule),
		}},
		"yearly_retention": []interface{}{map[string]interface{}{
			"years":                   retentionCfg.YearlyRetention.Years,
			"month_specific_schedule": utils.ToInterfaceSlice(retentionCfg.MonthlyRetention.MonthSpecificSchedule),
		}},
	}}
}
