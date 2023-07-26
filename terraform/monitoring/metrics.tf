resource "null_resource" "metrics" {
  depends_on = [null_resource.service_account]
  provisioner "local-exec" {
    command = "kubectl apply -f ${path.module}/metrics.yaml"
  }
}
