resource "kubernetes_config_map" "config_blue" {
  metadata {
    name = local.blue_config
  }

  data = {
    "config.yaml" = templatefile("${path.module}/config.yaml.tftpl", {
      db_host     = local.primary_host
      db_name     = local.primary_database
      db_username = local.primary_username
      db_password = local.primary_password
    })
  }
}

resource "kubernetes_config_map" "green_config" {
  metadata {
    name = local.green_config
  }

  data = {
    "config.yaml" = templatefile("${path.module}/config.yaml.tftpl", {
      db_host     = local.secondary_host
      db_name     = local.secondary_database
      db_username = local.secondary_username
      db_password = local.secondary_password
    })
  }
}

resource "kubernetes_config_map" "k6_config" {
  metadata {
    name = local.k6_config
  }

  data = {
    "script.js" = file("script.js")
  }
}
