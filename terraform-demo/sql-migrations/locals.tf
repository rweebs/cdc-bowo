locals {
  primary_endpoint_address   = data.aws_db_instance.primary.address
  secondary_endpoint_address = data.aws_db_instance.secondary.address
  primary_username           = "postgres"
  primary_password           = "CuTGUoIA"
  secondary_username         = "postgres"
  secondary_password         = "CuTGUoIA"
  primary_database           = "postgres"
  secondary_database         = "postgres"
}
