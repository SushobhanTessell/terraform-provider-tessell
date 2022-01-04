resource "tessell_dap" "example_dap" {
  dmm_name    = "database-dmm"
  name        = "example-dap"
  user_groups = [""]
  users       = [""]

  target_cloud_locations {
    aws = ["us-east-1"]
  }

  retention_config {
    pitr_retention {
      days = 7
    }
    daily_retention {
      days = 7
    }
    monthly_retention {
      common_schedule {
        dates             = [1, 7, 15, 21]
        last_day_of_month = false
      }
      months = 10
    }
    weekly_retention {
      days  = ["Sunday"]
      weeks = 5
    }
    yearly_retention {
      month_specific_schedule {
        dates = [31]
        month = "January"
      }
      years = 5
    }
  }
}
