
resource "kubernetes_namespace" "linkerd" {
  metadata {
    name = "linkerd"
  }
}

resource "helm_release" "linkerd-crds" {
  name       = "linkerd"
  repository = "https://helm.linkerd.io/stable"
  chart      = "linkerd-crds"
  namespace  = "linkerd"
  depends_on = [kubernetes_namespace.linkerd]
}

resource "helm_release" "linkerd-control-plane" {
  name       = "linkerd-control-plane"
  repository = "https://helm.linkerd.io/stable"
  chart      = "linkerd-control-plane"
  namespace  = "linkerd"
  depends_on = [helm_release.linkerd-crds]
  set {
    name  = "installNamespace"
    value = "false"
  }

  set {
    name  = "identityTrustAnchorsPEM"
    value = file("ca.crt")
  }

  set {
    name  = "identity.issuer.tls.crtPEM"
    value = file("issuer.crt")
  }

  set {
    name  = "identity.issuer.tls.keyPEM"
    value = file("issuer.key")
  }

}

resource "helm_release" "linkerd_viz" {
  name             = "linkerd-viz"
  namespace        = "linkerd"
  repository       = "https://helm.linkerd.io/stable"
  chart            = "linkerd-viz"
  create_namespace = false
}
