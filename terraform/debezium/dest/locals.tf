locals {
  # primary
  db_host     = data.aws_db_instance.secondary.address
  db_username = "postgres"
  db_password = "CuTGUoIA"
  db_database = "postgres"

  # redis

  redis_host = data.aws_elasticache_replication_group.redis.primary_endpoint_address

  # config
  app_config   = "dest-config"
  topic_prefix = "cdc-dest"
  app_name     = "debezium-dest"
  app_label    = "debezium-dest"
}
