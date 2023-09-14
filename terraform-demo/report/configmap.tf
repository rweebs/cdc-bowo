resource "kubernetes_config_map" "app_config" {
  metadata {
    name = local.app_config
  }

  data = {
    "config.json" = templatefile("${path.module}/config.json.tftpl", {
      source_host                 = local.source_host
      source_username             = local.source_username
      source_password             = local.source_password
      source_database             = local.source_database
      source_redis_topic_prefix   = local.source_redis_topic_prefix
      source_debezium_publication = local.source_debezium_publication
      source_port                 = local.source_port

      dest_host                 = local.dest_host
      dest_username             = local.dest_username
      dest_password             = local.dest_password
      dest_database             = local.dest_database
      dest_redis_topic_prefix   = local.dest_redis_topic_prefix
      dest_debezium_publication = local.dest_debezium_publication
      dest_port                 = local.dest_port
      dest_subscription_name    = local.dest_subscription_name

      redis_host     = local.redis_host
      redis_port     = local.redis_port
      redis_password = local.redis_password

      sql_file_path = local.sql_file_path
    })

    "demo.sql" = file("./demo.sql")
    "run.sh"   = file("./run.sh")
  }
}

