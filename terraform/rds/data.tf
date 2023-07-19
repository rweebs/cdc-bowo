data "aws_subnets" "db" {
  filter {

    name   = "tag:Name"
    values = ["bowo-vpc-public*"]
  }
}
