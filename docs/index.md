---
page_title: "Provider: Authz"
subcategory: ""
description: |-
  Terraform provider for interacting with [Authz](https://github.com/eko/authz) API.
---

# Authz Provider

The Authz provider is used to interact with [Authz](https://github.com/eko/authz) API.

1. use providers to [create, read, update and delete Authz (CRUD) resources](https://github.com/eko/authz/blob/master/backend/api/proto/api.proto) using Terraform.

Use the navigation to the left to read about the available resources.

## Example Usage

Do not keep your authentication password in HCL for production environments, use Terraform environment variables.

```terraform
provider "authz" {
  client_id = "[authz_sa_client_id]"
  client_secret = "[authz_sa_secret_id]"
  host = "[Dns/Ip:Port]"
}
```

## Schema

### Required

- **client_id** (String, Required) Authz service account client_id to authenticate Authz API
- **client_secret** (String, Required) Authz service account client_secret to authenticate Authz API
- **host** (String, Required) Authz API address