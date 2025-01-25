resource "kubernetes_namespace" "example" {
  metadata {
    name = var.infobae_namespace
  }
}

resource "kubernetes_deployment" "infobae_api_prod" {
  metadata {
    name = "infobae-api-prod"
    labels = {
      io_kompose_service = "infobae-api-prod"
    }
  }

  spec {
    replicas = 2

    selector {
      match_labels = {
        io_kompose_service = "infobae-api-prod"
      }
    }

    template {
      metadata {
        labels = {
          io_kompose_service = "infobae-api-prod"
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
    name = "infobae-api-prod"
    labels = {
      "io.kompose.service" = "infobae-api-prod"
    }
  }

  spec {
    selector = {
      "io.kompose.service" = "infobae-api-prod"
    }

    port {
      name        = "tcp"
      port        = 3000
      target_port = 3000
    }

    type = "NodePort"
  }
}

