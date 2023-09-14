data "aws_db_instance" "primary" {
  db_instance_identifier = "cdc-primary"
}

data "aws_db_instance" "secondary" {
  db_instance_identifier = "cdc-secondary"
}

data "aws_subnets" "db" {
  filter {

    name   = "tag:Name"
    values = ["cdc-vpc-public*"]
  }
}


data "aws_vpc" "cdc_vpc" {
  filter {

    name   = "tag:Name"
    values = ["cdc-vpc"]
  }
}