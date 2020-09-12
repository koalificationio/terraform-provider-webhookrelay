---
layout: "webhookrelay"
page_title: "Webhookrelay: webhookrelay_function"
description: |-
  Manage Webhookrelay function.
---

# Resource: webhookrelay_function

Use this resource to manage Webhookrelay function.

## Example Usage

```hcl
resource "webhookrelay_function" "foo" {
  name    = "foo"
  payload = filebase64("my_function.lua")
  driver  = "lua"

  config = {
    "SOME_NAME" = "secret"
  }
}
```

## Argument Reference

* `name` - (Required) Name of a bucket to create.
* `payload` - (Required) base64 encoded function payload. Check [webhookrelay documentation][1] on creating functions.
* `driver` - (Optional) Driver name. Currently available are `lua` or `wasi`. Defaults to `lua`.*
* `config` - (Optional) Mapping of config values for function.

## Attributes Reference

`id` is set to ID of the function.

### Config

The `config` field allows following attributes:

* `key` - (Required) Key that will be used to access it from the function
* `value` - (required) Your configuration value

## Import

Function can be imported using its `id`, e.g.

```
$ terraform import webhookrelay_function.foo e3cb4587-6e3d-4c64-9b50-e9c4c7ce27aa
```

[1]: https://webhookrelay.com/v1/guide/functions#Lua-functions-reference
