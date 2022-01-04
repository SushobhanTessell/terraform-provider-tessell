resource "tessell_database" "new_database" {
  service_name    = "new-database"
  dmm_name        = "new-database-dmm"
  database_name   = "db1"
  description     = "Example Database Description"
  master_user     = "anything"
  master_password = "NotEasyPassword"
  engine_type     = "POSTGRESQL"
  subscription    = "existing-subscription"

  infrastructure {
    allowed_ip_addresses = []
    cloud                = "azure"
    compute_type         = "Standard_L8s_v2"
    database_port        = "5432"
    enable_public_access = true
    region               = "eastUS"
    vpc_name             = "existing-vpc"
  }

  database_configuration {
    auto_minor_version_update  = true
    character_set              = "AL32UTF16"
    create_replica             = false
    enable_deletion_protection = true
    license_type               = "BYOL"
    multitenant                = false
    national_character_set     = "AL16UTF16"
    options_profile            = null
    parameter_profile          = "PostgreSQL 13 Profile"
    software_image_version     = "PostgreSQL 13.3 on RHEL8.3"
    maintenance_window {
      day      = "Sunday"
      time     = "10:57"
      duration = 120
    }
  }
  backup_configuration {
    auto_backup = false
    backup_sla  = "Brass"
    backup_window {
      time     = "10:57"
      duration = 120
    }
  }
}



resource "tessell_database" "cloned_database" {
  service_name        = "clone-database"
  clone_database_name = "clonedb1"
  dmm_name            = "clone-database-dmm"
  backup              = "existing-backup-name"
  clone_dmm_name      = "backup-source-dmm"
  master_user         = "master"
  master_password     = "veryStrongPassword"
  subscription        = "existing-subscription"
  engine_type         = "ORACLE"
  infrastructure {
    allowed_ip_addresses = []
    availability_zone    = "us-east-1a"
    cloud                = "aws"
    compute_type         = "i3.small"
    database_port        = "1521"
    enable_public_access = true
    region               = "us-east-1"
    vpc_name             = "existing-vpc"
  }
  database_configuration {
    software_image_version    = "Oracle 12.1 RTM on RHEL8.3"
    license_type              = "BYOL"
    parameter_profile         = "Oracle Parameter Profile"
    options_profile           = "Oracle 12.1.0.2.0 Options Profile"
    auto_minor_version_update = true
  }
}
