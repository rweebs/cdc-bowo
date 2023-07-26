
resource "kubernetes_config_map" "insert" {
  metadata {
    name      = "insert"
    namespace = "performance-native"
  }
  data = {
    "10.sql"  = file("./script/insert/10.sql")
    "20.sql"  = file("./script/insert/20.sql")
    "30.sql"  = file("./script/insert/30.sql")
    "40.sql"  = file("./script/insert/40.sql")
    "50.sql"  = file("./script/insert/50.sql")
    "100.sql" = file("./script/insert/100.sql")
  }
  depends_on = [kubernetes_namespace.example]
}


resource "kubernetes_config_map" "delete" {
  metadata {
    name      = "delete"
    namespace = "performance-native"
  }
  data = {
    "10.sql"  = file("./script/delete/10.sql")
    "20.sql"  = file("./script/delete/20.sql")
    "30.sql"  = file("./script/delete/30.sql")
    "40.sql"  = file("./script/delete/40.sql")
    "50.sql"  = file("./script/delete/50.sql")
    "100.sql" = file("./script/delete/100.sql")
  }
  depends_on = [kubernetes_namespace.example]
}
