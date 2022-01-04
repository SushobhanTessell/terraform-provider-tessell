terraform {
  required_providers {
    tessell = {
      version = "0.0.1"
      source  = "tessell.com/dev/tessell"
    }
  }
}

provider "tessell" {
  email_id    = "TESSELL_EMAIL_ID"    // Can be skipped if TESSELL_EMAIL_ID is set in env
  password    = "TESSELL_PASSWORD"    // Can be skipped if TESSELL_PASSWORD is set in env
  api_address = "TESSELL_API_ADDRESS" // Can be skipped if TESSELL_API_ADDRESS is set in env
}
