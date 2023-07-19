locals {
  region = "us-west-2"
}

# module "redis" {
#   source = "cloudposse/elasticache-redis/aws"
#   # Cloud Posse recommends pinning every module to a specific version
#   availability_zones            = ["us-west-2a", "us-west-2b", "us-west-2c"]
#   vpc_id                        = data.aws_vpc.cdc_vpc.id
#   allowed_cidr_blocks           = ["10.0.0.0/16", "0.0.0.0/0"]
#   subnets                       = data.aws_subnets.db.ids
#   cluster_size                  = 1
#   instance_type                 = "cache.t3.micro"
#   apply_immediately             = true
#   automatic_failover_enabled    = false
#   engine_version                = "6.x"
#   family                        = "redis6.x"
#   at_rest_encryption_enabled    = false
#   transit_encryption_enabled    = false
#   elasticache_subnet_group_name = "cdc-redis"
#   depends_on                    = [module.security_groups]
#   cluster_mode_enabled          = false
#   enabled                       = true
#   parameter_gro
# }

resource "aws_elasticache_subnet_group" "subnet_group" {
  name       = "cdc-cache-subnet"
  subnet_ids = data.aws_subnets.db.ids
}

resource "aws_elasticache_replication_group" "redis" {
  automatic_failover_enabled    = false
  availability_zones            = ["us-west-2a"]
  replication_group_id          = "tf-rep-group-1"
  replication_group_description = "test description"
  node_type                     = "cache.t3.micro"
  number_cache_clusters         = 1
  engine_version                = "6.x"
  port                          = 6379
  security_group_ids            = [module.security_groups.security_group_id]
  subnet_group_name             = aws_elasticache_subnet_group.subnet_group.name
}
