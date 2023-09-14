resource "kubernetes_token_request_v1" "test" {
  metadata {
    name      = "eks-admin"
    namespace = "kube-system"
  }
  spec {
    expiration_seconds = 86400
  }
  depends_on = [null_resource.service_account]
}

output "tokenValue" {
  value     = kubernetes_token_request_v1.test.token
  sensitive = true
}
