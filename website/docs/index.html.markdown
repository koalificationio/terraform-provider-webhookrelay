---
layout: "webhookrelay"
page_title: "Provider: webhookrelay"
description: |-
  Terraform webhookrelay provider.
---

# Webhookrelay Provider

The provider needs to be configured
with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
# Configure the webhookrelay Provider
provider "webhookrelay" {
  version    = "~> 0.1"
  api_key    = "123abc"
  api_secret = "1234567"
}
```

## Argument Reference

In addition to [generic `provider` arguments](https://www.terraform.io/docs/configuration/providers.html)
(e.g. `alias` and `version`), the following arguments are supported in the webhookrelay
 `provider` block:

* `api_key` - (Optional) This is the Webhookrelay API key. It must be provided, but
  it can also be sourced from the `RELAY_KEY` environment variable.
* `api_secret` - (Optional) This is the Webhookrelay API secret. It must be provided, but
  it can also be sourced from the `RELAY_SECRET` environment variable.
