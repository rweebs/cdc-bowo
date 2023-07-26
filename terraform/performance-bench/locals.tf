locals {
  # primary
  source_host                 = data.aws_db_instance.primary.address
  source_username             = "postgres"
  source_password             = "CuTGUoIA"
  source_database             = "postgres"
  source_redis_topic_prefix   = "cdc-source"
  source_debezium_publication = "dbz_publication"
  source_port                 = 5432
  # secondary
  dest_host                 = data.aws_db_instance.secondary.address
  dest_username             = "postgres"
  dest_password             = "CuTGUoIA"
  dest_database             = "postgres"
  dest_redis_topic_prefix   = "cdc-dest"
  dest_debezium_publication = "dbz_publication"
  dest_port                 = 5432
  dest_subscription_name    = "my_subscription"
  sql_file_path             = "/config/test.sql"

  # redis

  redis_host     = data.aws_elasticache_replication_group.redis.primary_endpoint_address
  redis_port     = 6379
  redis_password = ""

  #config
  app_image     = "rweebs/cdc:1.0.4"
  app_config    = "app-config"
  app_label     = "cdc-app"
  test_config   = "script_config"
  test_label    = "script"
  script_config = "script-config"

  # postgres
  primary_endpoint_address   = data.aws_db_instance.primary.address
  secondary_endpoint_address = data.aws_db_instance.secondary.address
  primary_username           = "postgres"
  primary_password           = "CuTGUoIA"
  secondary_username         = "postgres"
  secondary_password         = "CuTGUoIA"
  primary_database           = "postgres"
  secondary_database         = "postgres"

  #debezium
  # primary
  db_host     = data.aws_db_instance.primary.address
  db_username = "postgres"
  db_password = "CuTGUoIA"
  db_database = "postgres"

  #config

  debezium_config = "source-config"
  debezium_name   = "debezium-source"
  topic_prefix    = "cdc-source"
  debezium_label  = "debezium-source"
  job             = "change-ddl"

}
