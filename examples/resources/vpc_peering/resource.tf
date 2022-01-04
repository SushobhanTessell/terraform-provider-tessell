resource "tessell_vpc_peering" "new_aws_vpc_peering" {
  vpc_name          = "new-aws-vpc-peering"
  subscription_name = "subscription-containing-this-vpc"
  cloud_type        = "AWS"
  region            = "us-east-1"

  aws_client_vpc_info {
    client_account_id = "123456789012"
    client_vpc_id     = "vpc-7df1im2pc2efqqvpo"
    client_vpc_region = "us-west-1"
  }
}

resource "tessell_vpc_peering" "new_azure_vpc_peering" {
  vpc_name          = "new-azure-vpc-peering"
  subscription_name = "subscription-containing-this-vpc"
  cloud_type        = "AZURE"
  region            = "eastUS"

  azure_client_vpc_info {
    client_active_directory_tenant_id = "2bcacba8-8f72-469a-ad92-7bc75cb9ca9d"
    client_subscription_id            = "9cdf208b-a0be-4db6-a618-d0c0e64c5a27"
    client_resource_group             = "example_rg"
    client_vpc_name                   = "existing_vpc_in_example_rg"
    client_application_object_id      = "54b93613-94ae-429a-ac97-48fef253b960"
  }
}
