resource "kubernetes_config_map" "source_config" {
  metadata {
    name      = local.debezium_config
    namespace = "performance-bench"
  }

  data = {
    "application.properties" = templatefile("${path.module}/application.properties.tftmpl", {
      db_host      = local.db_host
      db_name      = local.db_database
      db_username  = local.db_username
      db_password  = local.db_password
      redis_host   = local.redis_host
      topic_prefix = local.topic_prefix
    })
  }
  depends_on = [kubernetes_namespace.example]
}

