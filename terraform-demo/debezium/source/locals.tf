locals {
  # primary
  db_host     = data.aws_db_instance.primary.address
  db_username = "postgres"
  db_password = "CuTGUoIA"
  db_database = "postgres"

  # redis

  redis_host = data.aws_elasticache_replication_group.redis.primary_endpoint_address

  #config
  app_config   = "source-config"
  app_name     = "debezium-source"
  topic_prefix = "cdc-source"
  app_label    = "debezium-source"

}
