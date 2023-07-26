locals {
  region = "us-west-2"
}
resource "aws_elasticache_subnet_group" "subnet_group" {
  name       = "cdc-cache-bench-subnet"
  subnet_ids = data.aws_subnets.db.ids
}

resource "aws_elasticache_replication_group" "redis" {
  automatic_failover_enabled    = false
  availability_zones            = ["us-west-2a"]
  replication_group_id          = "tf-rep-group-2"
  replication_group_description = "test description"
  node_type                     = "cache.t3.micro"
  number_cache_clusters         = 1
  engine_version                = "6.x"
  port                          = 6379
  security_group_ids            = [module.security_groups.security_group_id]
  subnet_group_name             = aws_elasticache_subnet_group.subnet_group.name
}
