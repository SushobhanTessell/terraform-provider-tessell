resource "tessell_vpc" "new_vpc" {
  name              = "new-vpc"
  cloud_type        = "AWS"
  region            = "us-west-1"
  subscription_name = "existing-subscription"
  cidr_block        = "10.0.0.0/16"

  block_until_complete = true // Default: false // Optional
}
