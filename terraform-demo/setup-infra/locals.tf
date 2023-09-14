locals {
  cluster_name = "cdc-eks-${random_string.suffix.result}"
  region       = "ap-southeast-1"
}