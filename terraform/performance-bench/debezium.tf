resource "kubectl_manifest" "debezium_pod" {
  yaml_body = templatefile("debezium.yaml.tftpl", {
    app_config = local.debezium_config,
    app_name   = local.debezium_name,
    app_label  = local.debezium_label,
  })

  depends_on = [kubernetes_config_map.source_config, sql_migrate.blue, sql_migrate.green, kubernetes_namespace.example]
}
