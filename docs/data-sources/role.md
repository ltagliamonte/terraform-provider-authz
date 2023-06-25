---
page_title: "Role Data Source - terraform-provider-authz"
subcategory: ""
description: |-
  The role data source allows you to retrieve information about an Authz roles.
---

# Data Source `authz_role`

-> Visit the [Official Authz Documentation](https://docs.authz.fr/#/) to learn more about role resourse.

The role data source allows you to retrieve information about an Authz roles.

## Example Usage

```terraform
data "authz_role" "role" {
  name = "authz-admin"
}
```

## Argument Reference

- `name` - (Required) Role Name.

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

- `policies` - Role policies.
