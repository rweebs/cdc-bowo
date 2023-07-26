data "aws_db_instance" "primary" {
  db_instance_identifier = "native-primary"
}


data "aws_db_instance" "secondary" {
  db_instance_identifier = "native-secondary"
}


data "aws_elasticache_replication_group" "redis" {
  replication_group_id = "tf-rep-group-1"
}
