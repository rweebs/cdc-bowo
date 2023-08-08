# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

data "aws_availability_zones" "available" {}

locals {
  cluster_name = "cdc-eks-${random_string.suffix.result}"
  region       = "us-west-2"
}

resource "random_string" "suffix" {
  length  = 8
  special = false
}

module "security_groups" {
  source  = "terraform-aws-modules/security-group/aws"
  version = "5.1.0"

  name        = "postgres-security-group"
  description = "Security group for PostgreSQL publicly open"
  vpc_id      = module.vpc.default_vpc_id

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
  depends_on = [module.vpc]
}

module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "5.0.0"

  name = "cdc-vpc"

  cidr = "10.0.0.0/16"
  azs  = slice(data.aws_availability_zones.available.names, 0, 3)

  private_subnets = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]
  public_subnets  = ["10.0.4.0/24", "10.0.5.0/24", "10.0.6.0/24"]

  enable_nat_gateway   = true
  single_nat_gateway   = true
  enable_dns_hostnames = true

  public_subnet_tags = {
    "kubernetes.io/cluster/${local.cluster_name}" = "shared"
    "kubernetes.io/role/elb"                      = 1
  }

  private_subnet_tags = {
    "kubernetes.io/cluster/${local.cluster_name}" = "shared"
    "kubernetes.io/role/internal-elb"             = 1
  }
}

module "blue" {
  source  = "terraform-aws-modules/rds/aws"
  version = "6.1.0"

  identifier = "cdc-primary"

  engine            = "postgres"
  engine_version    = "14.7"
  instance_class    = "db.t3.micro"
  allocated_storage = 5

  username = "postgres"
  password = "CuTGUoIA"

  vpc_security_group_ids = [module.security_groups.security_group_id]

  maintenance_window = "Mon:00:00-Mon:03:00"
  backup_window      = "03:00-06:00"

  # Enhanced Monitoring - see example for details on how to create the role
  # by yourself, in case you don't want to create it automatically
  monitoring_interval    = "30"
  monitoring_role_name   = "RDSMonitoringRole"
  create_monitoring_role = true
  publicly_accessible    = true
  apply_immediately      = true

  tags = {
    Owner       = "user"
    Environment = "dev"
  }

  # DB subnet group
  create_db_subnet_group = true
  subnet_ids             = module.vpc.public_subnets_cidr_blocks

  # DB parameter group
  family = "postgres14"

  # DB option group
  major_engine_version = "14.7"

  iam_database_authentication_enabled = false
  manage_master_user_password         = false

  parameters = [
    {
      name         = "rds.logical_replication"
      value        = 1
      apply_method = "pending-reboot"
    }
  ]

  depends_on = [module.security_groups]
}

