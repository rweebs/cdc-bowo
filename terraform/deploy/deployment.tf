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
resource "kubectl_manifest" "k6_pod" {
  yaml_body = templatefile("k6.yaml.tftpl", {
    app_config = local.k6_config
  })
}
