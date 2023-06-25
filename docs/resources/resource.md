---
page_title: "Resource Resource - terraform-provider-authz"
subcategory: ""
description: |-
  The resource resource allows you to create/update/delete Authz resources.
---

# Resource `authz_resource`

-> Visit the [Official Authz Documentation](https://docs.authz.fr/#/) to learn more about Resource resourse.

## Example Usage

```terraform
resource "authz_resource" "resource" {
  name = "deploys.*"
  kind = "deploy"
  value = "*"

  attributes = {
    key_0 = "value"
    key_1 = "value"
  }
}
```

## Argument Reference

- `name` - (Required) Unique Resource Name.
- `kind` - (Required) Resource kind.
- `value` - (Required) Resource value.
- `attributes` - (Optional) Resource attributes.
