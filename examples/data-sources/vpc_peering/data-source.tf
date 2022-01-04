data "tessell_vpc_peerings" "existing_peers_of_vpc" {
  vpc_name          = "vpc-peered-with"
  cloud_type        = "AWS"
  region            = "us-east-1"
  subscription_name = "subscription-containing-this-vpc"
}

data "tessell_vpc_peering" "existing_peer_of_vpc" {
  name              = "peering-name"
  vpc_name          = "vpc-peered-with"
  cloud_type        = "AWS"
  region            = "us-east-1"
  subscription_name = "subscription-containing-this-vpc"
}
