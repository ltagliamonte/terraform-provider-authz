---
page_title: "Principal Resource - terraform-provider-authz"
subcategory: ""
description: |-
  The principal resource allows you to create/update/delete Authz principals.
---

# Resource `authz_principal`

-> Visit the [Official Authz Documentation](https://docs.authz.fr/#/) to learn more about principal resourse.

## Example Usage

```terraform
resource "authz_principal" "principal" {
  name = "test-principal"
  roles = ["authz-admin"]

  attributes = {
    Name        = "test"
    Environment = "test0"
  }
}
```

## Argument Reference

- `name` - (Required) Unique Principal Name.
- `roles` - (Optional) Principal roles.
- `attributes` - (Optional) Principal attributes.
