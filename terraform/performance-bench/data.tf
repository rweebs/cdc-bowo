data "aws_db_instance" "primary" {
  db_instance_identifier = "bench-primary"
}


data "aws_db_instance" "secondary" {
  db_instance_identifier = "bench-secondary"
}


data "aws_elasticache_replication_group" "redis" {
  replication_group_id = "tf-rep-group-2"
}
