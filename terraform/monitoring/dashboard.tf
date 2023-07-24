resource "null_resource" "dashboard" {
  depends_on = [null_resource.service_account]
  provisioner "local-exec" {
    command = "kubectl apply -f ${path.module}/dashboard.yaml"
  }
}
