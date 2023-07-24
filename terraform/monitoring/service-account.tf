resource "kubectl_manifest" "service_account" {
  yaml_body = file("${path.module}/service-account.yaml")
}
