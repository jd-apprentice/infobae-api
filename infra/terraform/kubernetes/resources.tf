resource "kubernetes_namespace" "infobae_api_prod" {
  metadata {
    name = var.infobae_namespace
  }
}

resource "kubernetes_deployment" "infobae_api_prod" {
  metadata {
    namespace = var.infobae_namespace
    name      = "infobae-api-prod"
    labels = {
      "io.kompose.service" = "infobae-api-prod"
      restarted_at         = replace(replace(replace(timestamp(), ":", ""), "T", "_"), "Z", "")
    }

    annotations = {
      restarted_at = timestamp()
    }
  }

  spec {
    replicas = 3

    selector {
      match_labels = {
        "io.kompose.service" = "infobae-api-prod"
      }
    }

    template {
      metadata {
        labels = {
          "io.kompose.service" = "infobae-api-prod"
        }
      }

      spec {
        container {
          name              = "infobae-api-prod"
          image             = "dyallo/infobae_api:X64_latest"
          image_pull_policy = "Always"

          port {
            container_port = 3000
            protocol       = "TCP"
          }
        }

        restart_policy = "Always"

      }
    }
  }
}

resource "kubernetes_service" "infobae_api_prod" {
  metadata {
    namespace = var.infobae_namespace
    name      = "infobae-api-prod"
    labels = {
      "io.kompose.service" = "infobae-api-prod"
      restarted_at         = replace(replace(replace(timestamp(), ":", ""), "T", "_"), "Z", "")
    }

    annotations = {
      restarted_at = timestamp()
    }
  }

  spec {
    selector = {
      "io.kompose.service" = "infobae-api-prod"
    }

    port {
      name        = "3000"
      port        = 3000
      target_port = 3000
      node_port   = 31430
    }

    type = "NodePort"
  }
}
