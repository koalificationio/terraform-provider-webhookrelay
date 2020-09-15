---
description: |-
  Terraform webhookrelay provider.
---

# Webhookrelay Provider

The provider needs to be configured
with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
# Configure the Webhookrelay Provider
provider "webhookrelay" {
  version    = "~> 0.2"
}

# Create bucket
resource "webhookrelay_bucket" "example" {
  name                 = "foo"
  description          = "bar"
}
```

## Authentication

The following authentication methods are supported:
  * Static credentials
  * Environment variables

### Static credentials

!> Hard-coding credentials into any Terraform configuration is not recommended,
and risks secret leakage should this file ever be committed to a public version
control system.

Static credentials can be provided by adding an `api_key` and `api_secret` in-line
in the Webhookrelay provider block:

```hcl
# Configure the Webhookrelay Provider
provider "webhookrelay" {
  api_key    = "1234qwerty"
  api_secret = "secret"
}
```

### Environment variables

You can provide your credentials via the `RELAY_KEY` and `RELAY_SECRET` environment variables.

Provider configuration:
```hcl
provider "webhookrelay" {}
```

Usage:
```shell
$ export RELAY_KEY="1234qwerty"
$ export RELAY_SECRET="secret"
$ terraform plan
```

## Argument Reference

In addition to [generic `provider` arguments](https://www.terraform.io/docs/configuration/providers.html)
(e.g. `alias` and `version`), the following arguments are supported in the Webhookrelay
 `provider` block:

* `api_key` - (Optional) This is the Webhookrelay API key. It must be provided, but
  it can also be sourced from the `RELAY_KEY` environment variable.
* `api_secret` - (Optional) This is the Webhookrelay API secret. It must be provided, but
  it can also be sourced from the `RELAY_SECRET` environment variable.
