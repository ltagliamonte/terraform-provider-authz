---
page_title: "Role Resource - terraform-provider-authz"
subcategory: ""
description: |-
  The role resource allows you to create/update/delete Authz roles.
---

# Resource `authz_role`

-> Visit the [Official Authz Documentation](https://docs.authz.fr/#/) to learn more about Role resourse.

## Example Usage

```terraform
resource "authz_role" "role" {
  name = "test-role"
  policies = ["authz-audits-admin", "authz-policies-admin"]
}
```

## Argument Reference

- `name` - (Required) Unique Role Name.
- `policies` - (Optional) Role policies.
