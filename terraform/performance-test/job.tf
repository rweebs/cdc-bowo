resource "kubectl_manifest" "postgres" {
  yaml_body = templatefile("${path.module}/job.yaml.tftpl", {
    source_host     = local.source_host
    source_password = local.source_password
    script_config   = local.script_config
    dest_host       = local.dest_host

  })
  depends_on = [kubernetes_config_map.update, kubernetes_config_map.insert, kubernetes_config_map.delete, kubernetes_config_map.test_script, sql_migrate.blue, sql_migrate.green, kubectl_manifest.deployment, kubernetes_namespace.example]
  #   provisioner "local-exec" {
  #     command = "bash generate-report.sh"
  #   }
}

# resource "null_resource" "name" {
#   provisioner "local-exec" {
#     command = "POSTGRES_HOST=${local.dest_host} python3 -m generate-report.py"
#   }
#   depends_on = [  ]
# }
