terraform {
  required_providers {
    juris = {
      version = "0.0.1"
      source = "home.com/edu/juris"
    }
  }
}

provider "juris" {}

module "test" {
  source = "./jokes"
}

output "psl" {
  value = module.test.all_jokes
}