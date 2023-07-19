resource "kubectl_manifest" "change-ddl" {
  yaml_body = templatefile("${path.module}/job.yaml.tftpl", {
    image      = local.app_image
    job        = local.job
    app_config = local.app_config
  })
  depends_on = [kubernetes_config_map.app_config]
}
