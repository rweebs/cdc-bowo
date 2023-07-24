resource "null_resource" "service_account" {
  provisioner "local-exec" {
    command = "kubectl apply -f ${path.module}/service-account.yaml"
  }
}
