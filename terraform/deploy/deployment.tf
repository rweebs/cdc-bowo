resource "kubectl_manifest" "blue_deployment" {
  yaml_body = templatefile("deployment.yaml.tftpl", {
    app_label  = local.blue_label
    app_config = local.blue_config
    app_image  = local.blue_image
  })
}

resource "kubectl_manifest" "green_deployment" {
  yaml_body = templatefile("deployment.yaml.tftpl", {
    app_label  = local.green_label
    app_config = local.green_config
    app_image  = local.green_image
  })
}
