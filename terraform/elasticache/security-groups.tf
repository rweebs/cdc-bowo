module "security_groups" {
  source = "terraform-aws-modules/security-group/aws"

  name        = "redis-security-group"
  description = "Security group for PostgreSQL publicly open"
  vpc_id      = data.aws_vpc.cdc_vpc.id

  ingress_cidr_blocks = ["10.0.0.0/16", "0.0.0.0/0"]
  ingress_with_cidr_blocks = [
    {
      rule        = "redis-tcp"
      description = "eks cluster access"
      cidr_blocks = "10.0.0.0/16"
    },
    {
      rule        = "redis-tcp"
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
