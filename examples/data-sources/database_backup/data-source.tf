data "tessell_database_backups" "all_backups_within_dmm" {
  dmm_name = "existing-dmm"
}

data "tessell_database_backup" "backup_within_dmm" {
  dmm_name    = "existing-dmm"
  backup_name = "existing-backup"
}
