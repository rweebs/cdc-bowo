locals {
  // primary 
  primary_host     = data.aws_db_instance.primary.address
  primary_username = "postgres"
  primary_password = "CuTGUoIA"
  primary_database = "postgres"

  // secondary
  secondary_host     = data.aws_db_instance.secondary.address
  secondary_username = "postgres"
  secondary_password = "CuTGUoIA"
  secondary_database = "postgres"

  // service
  service_label = "blue"

  // app
  blue_label  = "blue"
  green_label = "green"

  // configmap
  blue_config  = "app-config-blue"
  green_config = "app-config-green"
  k6_config    = "k6-config"

  // image
  blue_image  = "rweebs/helper_backend:1.0.0"
  green_image = "rweebs/helper_backend:2.0.0"
}
