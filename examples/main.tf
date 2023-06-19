# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

terraform {
  required_providers {
    authz = {
      version = "0.3.1"
      source = "localhost/edu/authz"
    }
  }
}

provider "authz" {
  username = "057ac096-0eef-11ee-8475-0242ac110002"
  password = "RdP9yfWD06CQXVa1hbpAWbaXdhuJ-ehmHN86c0qfcUUGGvEP"
  host = "dns:///127.0.0.1:8081"
}

/*
resource "authz_policy" "policy" {
  name = "post-manage"
  resources = ["post.*"]
  actions = ["delete", "edit"]
  attribute_rules = ["resource.owner_email == principal.email"]
}*/
