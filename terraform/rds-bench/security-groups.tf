data "aws_vpc" "cdc_vpc" {
  filter {
    name   = "tag:Name"
    values = ["cdc-vpc"]
  }
}


module "security_groups" {
  source  = "terraform-aws-modules/security-group/aws"
  version = "5.1.0"

  name        = "postgres-security-group"
  description = "Security group for PostgreSQL publicly open"
  vpc_id      = data.aws_vpc.cdc_vpc.id

  ingress_cidr_blocks = ["10.10.0.0/16", "0.0.0.0/0"]
  ingress_with_cidr_blocks = [
    {
      rule        = "postgresql-tcp"
      description = "eks cluster access"
      cidr_blocks = "10.0.0.0/16"
    },
    {
      rule        = "postgresql-tcp"
      cidr_blocks = "0.0.0.0/0"
      description = "public access"
    },
  ]
  egress_with_cidr_blocks = [
    {
      rule        = "all-all"
      cidr_blocks = "0.0.0.0/0"
      description = "allow all access"
    },
  ]
}
