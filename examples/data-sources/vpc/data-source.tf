data "tessell_vpcs" "existing_vpcs" {}

data "tessell_vpc" "existing_vpc" {
  name              = "existing-vpc-name"
  cloud_type        = "AWS"
  region            = "us-east-1"
  subscription_name = "subscription-containing-this-vpc"
}
