resource "kubernetes_token_request_v1" "test" {
  metadata {
    name      = "eks-admin"
    namespace = "kube-system"
  }
  depends_on = [kubectl_manifest.service_account]
}

output "tokenValue" {
  value     = kubernetes_token_request_v1.test.token
  sensitive = true
}
