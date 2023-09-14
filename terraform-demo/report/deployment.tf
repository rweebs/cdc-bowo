resource "kubectl_manifest" "deployment" {
  yaml_body = templatefile("${path.module}/deployment.yaml.tftpl", {
    app_config = local.app_config
    app_image  = local.app_image
    app_label  = local.app_label
  })

  depends_on = [kubernetes_config_map.app_config]
}
