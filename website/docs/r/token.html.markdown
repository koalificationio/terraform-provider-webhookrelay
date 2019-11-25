---
layout: "webhookrelay"
page_title: "Webhookrelay: webhookrelay_token"
description: |-
  Manage Webhookrelay token.
---

# Resource: webhookrelay_token

Use this resource to manage Webhookrelay token.

## Example Usage

```hcl
resource "webhookrelay_token" "foo" {
  description = "foo"
  api_access  = "disabled"
  scopes {
    buckets = ["test1", "test2"]
  }
}
```

## Argument Reference

* `description` - (Optional) description of a token.
* `api_access` - (Optional) allowed values: `disabled`, `enabled`. Defaults to `enabled`
* `scopes` - (Optional) set access scopes for a token. See [Scopes](#scopes) below for details.

### Scopes

The `scopes` mapping provides following attributes:

* `buckets` - List of allowed bucket names.
* `tunnels` - List of allowed tunnels.

## Attributes Reference

`id` is set to the Key of the token. In addition, the following attributes are exported:

* `key` - Token key.
* `secret` - Token secret.

~> **Note:** Api Secret attribute will be stored in the raw state as plain-text.
[Read more about sensitive data in state](/docs/state/sensitive-data.html).
