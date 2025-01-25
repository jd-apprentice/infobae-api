variable "infobae_namespace" {
  type      = string
  default   = "infobae-api"
  sensitive = false
}

variable "infobae-kubeconfig" {
  type      = string
  default   = "~/.kube/config.yaml"
  sensitive = false
}
