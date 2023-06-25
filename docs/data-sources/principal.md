---
page_title: "Principal Data Source - terraform-provider-authz"
subcategory: ""
description: |-
  The principal data source allows you to retrieve information about an Authz principals.
---

# Data Source `authz_principal`

-> Visit the [Official Authz Documentation](https://docs.authz.fr/#/) to learn more about principal resourse.

The principal data source allows you to retrieve information about an Authz principal.

## Example Usage

```terraform
data "authz_principal" "principal" {
  name = "authz-user-admin"
}
```

## Argument Reference

- `name` - (Required) Principal Name.

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

- `roles` - Principal roles.
- `attributes` - Principal attributes.
