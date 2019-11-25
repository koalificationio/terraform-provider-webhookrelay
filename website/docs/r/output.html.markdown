---
layout: "webhookrelay"
page_title: "Webhookrelay: webhookrelay_output"
description: |-
  Manage Webhookrelay output.
---

# Resource: webhookrelay_output

Use this resource to manage Webhookrelay output.

## Example Usage

```hcl
resource "webhookrelay_bucket" "foo" {
  name        = "foo"
  description = "foo"
}

resource "webhookrelay_output" "foo" {
  name        = "foo"
  description = "bar"
  destination = "https://example.com:8080"
  internal    = false
  bucket_id   = webhookrelay_bucket.foo.id
}
```

## Argument Reference

* `name` - (Required) Name of an output to create.
* `bucket_id` - (Required) bucket id to attach output to.

-> **NOTE:** Changing `bucket_id` will recreate output

* `description` - (Optional) description of an output.
* `destination` - (Required) output destination.
* `internal` - (Optional) set to `true` if output is on local network. Defaults to `false`.

## Attributes Reference

`id` is set to the ID of the output.

## Import

Output can be imported using the `id`, e.g.

```
$ terraform import webhookrelay_output.foo e3cb4587-6e3d-4c64-9b50-e9c4c7ce27aa
```
