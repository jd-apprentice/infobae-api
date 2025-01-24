variable "infobae_namespace" {
  type      = string
  default   = "infobae-api"
  sensitive = false
}

variable "infobae-kubeconfig" {
  type      = string
  default   = "/etc/rancher/k3s/k3s.yaml"
  sensitive = false
}
