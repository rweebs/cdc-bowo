resource "kubectl_manifest" "deployment" {
  yaml_body = templatefile("${path.module}/deployment.yaml.tftpl", {
    app_config = local.app_config
    app_image  = local.app_image
    app_label  = local.app_label
  })

  depends_on = [kubernetes_config_map.app_config]
}
resource "kubectl_manifest" "postgres" {
  yaml_body = templatefile("${path.module}/postgres.yaml.tftpl", {
    source_host     = local.source_host
    source_password = local.source_password
    script_config   = local.script_config

  })
  depends_on = [null_resource.update]
}
