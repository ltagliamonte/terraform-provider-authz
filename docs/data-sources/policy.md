---
page_title: "Policy Data Source - terraform-provider-authz"
subcategory: ""
description: |-
  The policy data source allows you to retrieve information about an Authz policies.
---

# Data Source `authz_policy`

-> Visit the [Official Authz Documentation](https://docs.authz.fr/#/) to learn more about policy resourse.

The policy data source allows you to retrieve information about an Authz policy.

## Example Usage

```terraform
data "authz_policy" "policy" {
  name = "authz-policies-admin"
}
```

## Argument Reference

- `name` - (Required) Policy Name.

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

- `resources` - Resources this policy applies to.
- `actions` - Actions this policy applies to.
- `attribute_rules` - Attributes rules of the policy.
