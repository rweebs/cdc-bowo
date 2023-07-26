output "rds_password" {
  description = "The password for the master DB user"
  value       = "CuTGUoIA"
}

output "blue_db_instance_address" {
  description = "The address of the RDS instance"
  value       = module.blue.db_instance_address
}


output "blue_db_instance_endpoint" {
  description = "The connection endpoint"
  value       = module.blue.db_instance_endpoint
}

output "blue_db_instance_identifier" {
  description = "The RDS instance identifier"
  value       = module.blue.db_instance_identifier
}

output "green_db_instance_address" {
  description = "The address of the RDS instance"
  value       = module.green.db_instance_address
}


output "green_db_instance_endpoint" {
  description = "The connection endpoint"
  value       = module.green.db_instance_endpoint
}

output "green_db_instance_identifier" {
  description = "The RDS instance identifier"
  value       = module.green.db_instance_identifier
}




