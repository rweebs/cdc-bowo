resource "aws_elasticache_subnet_group" "subnet_group" {
  name       = "cdc-cache-subnet"
  subnet_ids = module.vpc.private_subnets
}

resource "aws_elasticache_replication_group" "redis" {
  automatic_failover_enabled  = false
  replication_group_id        = "tf-rep-group-1"
  node_type                   = "cache.t3.micro"
  preferred_cache_cluster_azs = ["ap-southeast-1a"]
  num_cache_clusters          = 1
  engine_version              = "6.x"
  port                        = 6379
  security_group_ids          = [module.security_groups_redis.security_group_id]
  subnet_group_name           = aws_elasticache_subnet_group.subnet_group.name
  description                 = "example description"
}
