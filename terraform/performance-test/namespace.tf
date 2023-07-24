resource "kubernetes_namespace" "example" {
  metadata {
    name = "performance-testing"
  }
}
