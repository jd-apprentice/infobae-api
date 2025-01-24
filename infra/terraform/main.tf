terraform {
  cloud {
    organization = "dyallab"

    workspaces {
      name = "infobae-api"
    }
  }
}

module "infobae" {
  source = "./kubernetes"
}
