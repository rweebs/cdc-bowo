data "aws_db_instance" "primary" {
  db_instance_identifier = "cdc-primary"
}

data "aws_db_instance" "secondary" {
  db_instance_identifier = "cdc-secondary"
}
