data "tessell_daps" "daps_for_dmm" {
  dmm_name = "database-dmm"
}

data "tessell_dap" "specific_dap_for_dmm" {
  dmm_name = "database-dmm"
  name     = "existing-dap"
}
