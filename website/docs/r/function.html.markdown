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
}
```

## Argument Reference

* `name` - (Required) Name of a bucket to create.
* `payload` - (Required) base64 encoded function payload. Check [webhookrelay documentation][1] on creating functions.
* `driver` - (Optional) Driver name. Currently available are `lua` or `wasi`. Defaults to `lua`.

## Attributes Reference

`id` is set to ID of the function.

## Import

Function can be imported using its `id`, e.g.

```
$ terraform import webhookrelay_function.foo e3cb4587-6e3d-4c64-9b50-e9c4c7ce27aa
```

[1]: https://webhookrelay.com/v1/guide/functions#Lua-functions-reference
