resource "kubectl_manifest" "debezium_pod" {
  yaml_body = templatefile("deployment.yaml.tftpl", {
    app_config = local.app_config,
    app_name   = local.app_name,
    app_label  = local.app_label,
  })

  depends_on = [kubernetes_config_map.source_config]
}
