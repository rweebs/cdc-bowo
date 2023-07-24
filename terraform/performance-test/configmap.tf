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
    "test.sql" = file("./test.sql")
    "demo.sql" = file("./demo.sql")
  }
}

resource "kubernetes_config_map" "insert" {
  metadata {
    name = "insert"
  }
  data = {
    "10.sql"  = file("./script/insert/10.sql")
    "20.sql"  = file("./script/insert/20.sql")
    "30.sql"  = file("./script/insert/30.sql")
    "40.sql"  = file("./script/insert/40.sql")
    "50.sql"  = file("./script/insert/50.sql")
    "100.sql" = file("./script/insert/100.sql")
  }
}

resource "kubernetes_config_map" "update" {
  metadata {
    name = "update"
  }
  data = {
    "10.sql"  = file("./script/update/10.sql")
    "20.sql"  = file("./script/update/20.sql")
    "30.sql"  = file("./script/update/30.sql")
    "40.sql"  = file("./script/update/40.sql")
    "50.sql"  = file("./script/update/50.sql")
    "100.sql" = file("./script/update/100.sql")
  }
}

resource "kubernetes_config_map" "delete" {
  metadata {
    name = "delete"
  }
  data = {
    "10.sql"  = file("./script/delete/10.sql")
    "20.sql"  = file("./script/delete/20.sql")
    "30.sql"  = file("./script/delete/30.sql")
    "40.sql"  = file("./script/delete/40.sql")
    "50.sql"  = file("./script/delete/50.sql")
    "100.sql" = file("./script/delete/100.sql")
  }
}

resource "kubernetes_config_map" "test_script" {
  metadata {
    name = "test-script"
  }
  data = {
    "test.sh" = file("./script/test.sh")
  }
}
