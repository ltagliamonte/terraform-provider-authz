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
  client_id = "dacc46e0-12d0-11ee-a138-0242ac160002"
  client_secret = "6s89GprH2auEKoXzYWPlkYrvBjpGLcgCKjohWdOGi45891jl"
  host = "127.0.0.1:8081"
}


resource "authz_policy" "policy" {
  name = "post-manage"
  resources = ["post.*"]
  actions = ["edit","delete"]
  attribute_rules = ["resource.owner_email == principal.email"]
}

/*
resource "authz_resource" "resource" {
  name = "deploys.*"
  kind = "deploy"
  value = "*"

  attributes = {
    Name        = "test"
    Environment = "test0"
  }
}*/

/*
resource "authz_principal" "principal" {
  name = "test-principal"
  roles = ["authz-admin"]

  attributes = {
    Name        = "test"
    Environment = "test0"
  }
}*/

/*
resource "authz_role" "role" {
  name = "test-role"
  policies = ["authz-audits-admin","authz-policies-admin"]
}*/

/*
data "authz_policy" "policy" {
  name = "authz-policies-admin"
}

resource "authz_policy" "policy" {
  name = data.authz_policy.policy.name
  resources = data.authz_policy.policy.resources
  actions = ["edit","delete"]
  attribute_rules = ["resource.owner_email == principal.email"]
}
*/

/*
data "authz_principal" "principal" {
  name = "authz-user-admin"
}

resource "authz_principal" "principal" {
  name = data.authz_principal.principal.name
  roles = data.authz_principal.principal.roles

  attributes = {
    Name        = "test"
    Environment = "test0"
  }
}
*/

/*
data "authz_resource" "resource" {
  name = "authz.audits.*"
}

resource "authz_resource" "resource" {
  name = data.authz_resource.resource.name
  kind = data.authz_resource.resource.kind
  value = "*"
  attributes = data.authz_resource.resource.attributes
}
*/

/*
data "authz_role" "role" {
  name = "authz-admin"
}

resource "authz_role" "role" {
  name = data.authz_role.role.name
  policies = data.authz_role.role.policies
}*/