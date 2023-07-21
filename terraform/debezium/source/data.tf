data "aws_db_instance" "primary" {
  db_instance_identifier = "cdc-primary"
}


data "aws_elasticache_replication_group" "redis" {
  replication_group_id = "cdc-rep-group-1"
}
