---
page_title: "Resource Data Source - terraform-provider-authz"
subcategory: ""
description: |-
  The resource data source allows you to retrieve information about an Authz resources.
---

# Data Source `authz_resource`

-> Visit the [Official Authz Documentation](https://docs.authz.fr/#/) to learn more about resource resourse.

The resource data source allows you to retrieve information about an Authz resource.

## Example Usage

```terraform
data "authz_resource" "resource" {
  name = "authz.audits.*"
}
```

## Argument Reference

- `name` - (Required) Resource Name.

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

- `kind` - Resource kind.
- `value` - Resource value.
- `attributes` - Resource attributes.
