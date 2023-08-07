terraform {
  required_providers {
    juris = {
      version = "0.0.1"
      source = "home.com/edu/juris"
    }
  }
}

data "random_jokes" "all" {
  provider = "juris"
}

output "all_jokes" {
  value = data.random_jokes.all.jokes
}