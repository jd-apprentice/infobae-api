resource "kubernetes_namespace" "example" {
  metadata {
    name = var.infobae_namespace
  }
}
