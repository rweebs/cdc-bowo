resource "kubectl_manifest" "service" {
  yaml_body = templatefile("${path.module}/service.yaml.tftpl", {
    service_label = local.service_label
  })
}
