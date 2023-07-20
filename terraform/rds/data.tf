data "aws_subnets" "db" {
  filter {

    name   = "tag:Name"
    values = ["cdc-vpc-public*"]
  }
}
