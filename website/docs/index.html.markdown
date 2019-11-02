---
layout: "webhookrelay"
page_title: "Provider: webhookrelay"
sidebar_current: "docs-webhookrelay-index"
description: |-
  Terraform webhookrelay provider.
---

# webhookrelay Provider

The provider needs to be configured
with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
# Configure the webhookrelay Provider
provider "webhookrelay" {
  version   = "~> 0.1"
  api_token = "123abc"
}
```

## Argument Reference

In addition to [generic `provider` arguments](https://www.terraform.io/docs/configuration/providers.html)
(e.g. `alias` and `version`), the following arguments are supported in the webhookrelay
 `provider` block:

* `api_token` - (Optional) This is the webhookrelay api token. It must be provided, but
  it can also be sourced from the `webhookrelay_TOKEN` environment variable.
