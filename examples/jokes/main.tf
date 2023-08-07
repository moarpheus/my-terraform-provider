# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

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

# Returns all coffees
output "all_jokes" {
  value = data.random_jokes.all.jokes
}