data "aws_db_instance" "primary" {
  db_instance_identifier = "cdc-bowo-primary"
}


data "aws_elasticache_replication_group" "redis" {
  replication_group_id = "tf-rep-group-1"
}
