resource "kubectl_manifest" "dashboard" {
  yaml_body = file("${path.module}/dashboard.yaml")
}
