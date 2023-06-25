---
page_title: "policy Resource - terraform-provider-authz"
subcategory: ""
description: |-
  The policy resource allows you to create/update/delete Authz policies.
---

# Resource `authz_policy`

-> Visit the [Official Authz Documentation](https://docs.authz.fr/#/) to learn more about policy resourse.

## Example Usage

```terraform
resource "authz_policy" "policy" {
  name = "unique-policy-name"
  resources = ["post.*"]
  actions = ["edit","delete"]
  attribute_rules = ["resource.owner_email == principal.email"]
}
```

## Argument Reference

- `name` - (Required) Unique Policy Name.
- `resources` - (Required) Resources this policy applies to.
- `actions` - (Required) Actions this policy applies to.
- `attribute_rules` - (Optional) Attributes rules for the policy.
